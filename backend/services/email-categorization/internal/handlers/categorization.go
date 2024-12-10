package handlers

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/samiransarii/inboXpert/services/email-categorization/internal/models"
	"github.com/samiransarii/inboXpert/services/email-categorization/internal/utils/converter"

	mlclient "github.com/samiransarii/inboXpert/services/common/ml_client"
	mlpb "github.com/samiransarii/inboXpert/services/common/ml_server_protogen"
	pb "github.com/samiransarii/inboXpert/services/email-categorization/proto"
)

type CategorizationHandler struct {
	mlClient   mlclient.Service
	config     *models.Config
	workerPool chan struct{}
	pb.UnimplementedEmailCategorizationServiceServer
}

func NewCategorizationHandler(mlClient mlclient.Service, config *models.Config) *CategorizationHandler {
	return &CategorizationHandler{
		mlClient:   mlClient,
		config:     config,
		workerPool: make(chan struct{}, config.NumWorkers),
	}
}

func (h *CategorizationHandler) CategorizeEmail(ctx context.Context, req *pb.CategorizeRequest) (*pb.CategorizeResponse, error) {
	internalEmail := converter.FromProtoEmail(req.Email)

	result, err := h.processSingleEmail(ctx, internalEmail)
	if err != nil {
		return nil, fmt.Errorf("failed to process email: %w", err)
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

	// var errs []error
	// var errStrings []string
	// for err := range errChan {
	// 	errs = append(errs, err)
	// 	errStrings = append(errStrings, err.Error())
	// }

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

	// for _, alt := range mlResponse.Alternatives {
	// 	result.Categories = append(result.Categories, alt.Category)
	// }

	return result, nil
}
