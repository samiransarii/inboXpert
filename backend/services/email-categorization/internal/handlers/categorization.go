package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/samiransarii/inboXpert/services/email-categorization/internal/models"
	"github.com/samiransarii/inboXpert/services/email-categorization/internal/models/db"
	"github.com/samiransarii/inboXpert/services/email-categorization/internal/utils/converter"

	mlclient "github.com/samiransarii/inboXpert/services/common/ml_client"
	mlpb "github.com/samiransarii/inboXpert/services/common/ml_server_protogen"
	pb "github.com/samiransarii/inboXpert/services/email-categorization/proto"
)

// CategorizationHandler manages the lifecycle of email categorization requests.
// It handles the process from saving emails to the database, sending them to
// an ML service for categorization, handling batch requests, and storing the results.
type CategorizationHandler struct {
	config     *models.Config
	workerPool chan struct{}
	mlClient   mlclient.Service
	emailRepo  *EmailRepository
	pb.UnimplementedEmailCategorizationServiceServer
}

// NewCategorizationHandler creates a new CategorizationHandler given a machine learning client service,
// configuration parameters, and an EmailRepository for database persistence.
func NewCategorizationHandler(mlClient mlclient.Service, config *models.Config, emailRepo *EmailRepository) *CategorizationHandler {
	return &CategorizationHandler{
		config:     config,
		workerPool: make(chan struct{}, config.NumWorkers),
		mlClient:   mlClient,
		emailRepo:  emailRepo,
	}
}

// CategorizeEmail handles a single email categorization request.
// It:
// 1. Assigns a new UUID to the email and saves it to the database.
// 2. Sends the email to the ML service for categorization.
// 3. Stores the categorization results in the database.
// 4. Returns the categorization response as a protobuf message.
func (h *CategorizationHandler) CategorizeEmail(ctx context.Context, req *pb.CategorizeRequest) (*pb.CategorizeResponse, error) {
	// Generate a new UUID for tracking the email
	emailID := uuid.New().String()

	// Convert the incoming protobuf email into the internal service model
	internalEmail := converter.FromProtoEmail(req.Email)
	internalEmail.ID = emailID

	// Save the email details to the database
	err := h.emailRepo.SaveEmail(ctx, *internalEmail)
	if err != nil {
		return nil, fmt.Errorf("failed to save email to the database: %w", err)
	}

	// Process the email categorization via the ML service
	result, err := h.processSingleEmail(ctx, internalEmail)
	if err != nil {
		return nil, fmt.Errorf("failed to process email: %w", err)
	}

	// Convert the categories to JSON for storage
	categoriesJSON, err := json.Marshal(result.Categories)
	if err != nil {
		log.Printf("Failed to serialize categories: %v", err)
	}

	// Construct the categorization record and store it in the database
	categorizationRecord := db.CatgegoryRecord{
		ID:              uuid.New().String(),
		EmailID:         emailID,
		Categories:      string(categoriesJSON),
		ConfidenceScore: result.ConfidenceScore,
	}
	err = h.emailRepo.SaveCategory(ctx, categorizationRecord)
	if err != nil {
		log.Printf("Failed to save categorization record: %v", err)
	}

	// Convert the internal categorization result into a protobuf response
	return &pb.CategorizeResponse{
		Result: &pb.CategoryResult{
			Id:              result.EmailID,
			Categories:      result.Categories,
			ConfidenceScore: result.ConfidenceScore,
		},
	}, nil
}

// BatchCategorizeEmails handles batch categorization requests. It accepts a list of emails
// and processes them concurrently, respecting the configured worker count and maximum batch size.
// It returns a BatchCategorizeResponse containing categorized results for each processed email.
func (h *CategorizationHandler) BatchCategorizeEmails(ctx context.Context, req *pb.BatchCategorizeRequest) (*pb.BatchCategorizeResponse, error) {
	if len(req.Emails) > h.config.MaxBatchSize {
		return nil, fmt.Errorf("batch size %d exceeds maximum allowed size %d", len(req.Emails), h.config.MaxBatchSize)
	}

	results := make([]*pb.CategoryResult, 0, len(req.Emails))
	errChan := make(chan error, len(req.Emails))
	resultChan := make(chan *pb.CategoryResult, len(req.Emails))

	var wg sync.WaitGroup

	// Process each email in a separate goroutine, limited by h.workerPool.
	for _, pbEmail := range req.Emails {
		wg.Add(1)
		go func(pbEmail *pb.Email) {
			defer wg.Done()

			// Acquire a worker slot
			h.workerPool <- struct{}{}
			defer func() { <-h.workerPool }()

			internalEmail := converter.FromProtoEmail(pbEmail)
			result, err := h.processSingleEmail(ctx, internalEmail)
			if err != nil {
				errChan <- fmt.Errorf("email %s: %w", pbEmail.Id, err)
				return
			}

			resultChan <- converter.ToProtoCategoryResult(result)
		}(pbEmail)
	}

	// Once all goroutines complete, close the channels
	go func() {
		wg.Wait()
		close(resultChan)
		close(errChan)
	}()

	// Collect all results
	for result := range resultChan {
		results = append(results, result)
	}

	// If needed, errors can be read from errChan. Currently, the handler does not aggregate them.

	return &pb.BatchCategorizeResponse{
		Results: results,
	}, nil
}

// processSingleEmail sends a single email to the ML service and returns the categorization result.
// It includes a retry mechanism, attempting categorization multiple times if errors occur.
// On success, it returns a CategoryResult with the email ID, categories, and confidence score.
func (h *CategorizationHandler) processSingleEmail(ctx context.Context, email *models.Email) (*models.CategoryResult, error) {
	mlReq := &models.MLRequest{
		ID:        email.ID,
		Subject:   email.Subject,
		Body:      email.Body,
		Sender:    email.Sender,
		Recipents: email.Recipients,
		Headers:   email.Headers,
	}

	var serverResponse *mlpb.CategoryResponse
	var err error

	// Attempt to categorize the email multiple times if retries are configured
	for attempt := 0; attempt < h.config.RetryAttempts; attempt++ {
		serverResponse, err = h.mlClient.CategorizeEmail(ctx, converter.ToMLRequest(mlReq))
		if err == nil {
			break
		}
		log.Printf("Attempt %d failed: %v", attempt+1, err)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to categorize email after %d attempts: %w", h.config.RetryAttempts, err)
	}

	mlResponse := converter.FromMLResponse(serverResponse)

	return &models.CategoryResult{
		EmailID:         mlResponse.ID,
		Categories:      []string{mlResponse.Category},
		ConfidenceScore: mlResponse.ConfidenceScore,
	}, nil
}
