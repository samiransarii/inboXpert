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

type CategorizationHandler struct {
	config     *models.Config
	workerPool chan struct{}
	mlClient   mlclient.Service
	emailRepo  *EmailRepository
	pb.UnimplementedEmailCategorizationServiceServer
}

func NewCategorizationHandler(mlClient mlclient.Service, config *models.Config, emailRepo *EmailRepository) *CategorizationHandler {
	return &CategorizationHandler{
		config:     config,
		workerPool: make(chan struct{}, config.NumWorkers),
		mlClient:   mlClient,
		emailRepo:  emailRepo,
	}
}

func (h *CategorizationHandler) CategorizeEmail(ctx context.Context, req *pb.CategorizeRequest) (*pb.CategorizeResponse, error) {
	// Generate a new UUID for the email
	emailID := uuid.New().String()

	// Convert protobuf email to internal email model
	internalEmail := converter.FromProtoEmail(req.Email)
	internalEmail.ID = emailID

	// save email to the database
	err := h.emailRepo.SaveEmail(ctx, *internalEmail)
	if err != nil {
		return nil, fmt.Errorf("failed to save email to the database: %w", err)
	}

	// Process the email categorization
	result, err := h.processSingleEmail(ctx, internalEmail)
	if err != nil {
		return nil, fmt.Errorf("failed to process email: %w", err)
	}

	// Save categorization results to the database
	categoriesJSON, err := json.Marshal(result.Categories)
	if err != nil {
		log.Printf("Failed to categorize the email")
	}
	categorizationRecord := db.CatgegoryRecord{
		ID:              uuid.New().String(),
		EmailID:         emailID,
		Categories:      string(categoriesJSON),
		ConfidenceScore: result.ConfidenceScore,
	}
	err = h.emailRepo.SaveCategory(ctx, categorizationRecord)
	if err != nil {
		log.Printf("Failed to save categorizatrion record: %v", err)
	}

	// Convert internal CategoryResult to protobuf CategoryResult
	return &pb.CategorizeResponse{
		Result: &pb.CategoryResult{
			Id:              result.EmailID,
			Categories:      result.Categories,
			ConfidenceScore: result.ConfidenceScore,
		},
	}, nil
}

func (h *CategorizationHandler) BatchCategorizeEmails(ctx context.Context, req *pb.BatchCategorizeRequest) (*pb.BatchCategorizeResponse, error) {
	if len(req.Emails) > h.config.MaxBatchSize {
		return nil, fmt.Errorf("batch size %d exceeds maximum allowed size %d", len(req.Emails), h.config.MaxBatchSize)
	}

	results := make([]*pb.CategoryResult, 0, len(req.Emails))
	errChan := make(chan error, len(req.Emails))
	resultChan := make(chan *pb.CategoryResult, len(req.Emails))

	var wg sync.WaitGroup

	for _, pbEmail := range req.Emails {
		wg.Add(1)
		go func(pbEmail *pb.Email) {
			defer wg.Done()

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

	go func() {
		wg.Wait()
		close(resultChan)
		close(errChan)
	}()

	for result := range resultChan {
		results = append(results, result)
	}

	return &pb.BatchCategorizeResponse{
		Results: results,
	}, nil
}

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

	for attempt := 0; attempt < h.config.RetryAttempts; attempt++ {
		serverResponse, err = h.mlClient.CategorizeEmail(ctx, converter.ToMLRequest(mlReq))
		if err == nil {
			break
		}
		log.Printf("Attempt %d failed: %v", attempt+1, err)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to categorize email after %d attempts %w", h.config.RetryAttempts, err)
	}

	mlResponse := converter.FromMLResponse(serverResponse)

	result := &models.CategoryResult{
		EmailID:         mlResponse.ID,
		Categories:      []string{mlResponse.Category},
		ConfidenceScore: mlResponse.ConfidenceScore,
	}

	return result, nil
}
