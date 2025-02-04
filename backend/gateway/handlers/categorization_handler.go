package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	utils "github.com/samiransarii/inboXpert/common/utils"
	pb "github.com/samiransarii/inboXpert/services/email-categorization/proto"
)

// CategorizationHandler coordinates the process of receiving incoming email data,
// forwarding it to the email categorization gRPC service, and returning categorized results.
// It:
// - Parses incoming JSON requests containing one or more emails.
// - Connects to the email categorization service via gRPC.
// - Sends emails for categorization and collects responses.
// - Returns a consolidated response that includes both successful results and any failures.
type CategorizationHandler struct {
	grpcManager *utils.GRPCClientManager
	serviceAddr string
	grpcTimeout time.Duration
}

// NewCategorizationHandler creates and returns a new instance of CategorizationHandler with a default
// gRPC connection manager, the service address, and a timeout configured.
func NewCategorizationHandler() *CategorizationHandler {
	return &CategorizationHandler{
		grpcManager: utils.GetGRPCClientManager(),
		serviceAddr: "localhost:50051",
		grpcTimeout: 15 * time.Second,
	}
}

// Handle is the main entry point for categorizing emails. It expects a JSON payload containing
// an array of emails. It then:
//   - Parses and validates the incoming request.
//   - Establishes a connection to the gRPC categorization service.
//   - Iterates over each email, sends it for categorization, and records the result.
//   - Responds with a JSON object reporting how many emails were processed, how many succeeded,
//     how many failed, and the details of the results.
func (h *CategorizationHandler) Handle(c *gin.Context) {
	// Create a context with a timeout for the gRPC call
	ctx, cancel := context.WithTimeout(context.Background(), h.grpcTimeout)
	defer cancel()

	var requestData CategorizeServiceRequest
	if err := h.parseRequest(c, &requestData); err != nil {
		h.handleError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Attempt to establish a gRPC connection to the categorization service
	conn, err := h.grpcManager.GetConnection(ctx, h.serviceAddr)
	if err != nil {
		h.handleError(c, http.StatusServiceUnavailable, "Failed to connect to service", err)
		return
	}

	client := pb.NewEmailCategorizationServiceClient(conn)

	var responses []*pb.CategorizeResponse
	var failedEmails []FailedEmail

	// Process each email individually
	for _, email := range requestData.Emails {
		grpcRequest := h.createGRPCRequest(email)

		// Send the email for categorization via gRPC
		response, err := client.CategorizeEmail(ctx, grpcRequest)
		if err != nil {
			log.Printf("Error processing email %s: %v", email.ID, err)
			failedEmails = append(failedEmails, FailedEmail{
				ID:    email.ID,
				Error: err.Error(),
			})
			continue
		}

		responses = append(responses, response)
	}

	// Construct a unified JSON response indicating the success/failure of each processed email
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"total_processed":      len(requestData.Emails),
			"successful_responses": len(responses),
			"failed_responses":     len(failedEmails),
			"results":              responses,
			"failed":               failedEmails,
		},
	})
}

// FailedEmail represents information about an email that could not be categorized successfully.
// It stores the email's unique identifier and the error message returned by the service.
type FailedEmail struct {
	// ID is the unique identifier of the email that failed categorization.
	ID string `json:"id"`
	// Error is the error message describing why the email failed to be categorized.
	Error string `json:"error"`
}

// createGRPCRequest transforms an EmailRequest into a gRPC CategorizeRequest message
// to be sent to the categorization service.
func (h *CategorizationHandler) createGRPCRequest(email EmailRequest) *pb.CategorizeRequest {
	return &pb.CategorizeRequest{
		Email: &pb.Email{
			Id:         email.ID,
			Subject:    email.Subject,
			Body:       email.Body,
			Sender:     email.Sender,
			Recipients: email.Recipients,
			Headers:    email.Headers,
		},
	}
}

// parseRequest attempts to parse the incoming JSON request body into a CategorizeServiceRequest.
// It also validates that at least one email is provided.
func (h *CategorizationHandler) parseRequest(c *gin.Context, req *CategorizeServiceRequest) error {
	if err := c.ShouldBindJSON(req); err != nil {
		return err
	}
	if len(req.Emails) == 0 {
		return fmt.Errorf("at least one email is required")
	}
	return nil
}

// handleError logs the specified error and returns a JSON response with the provided status code
// and a descriptive message, along with the error details. This method ensures consistent error
// handling and uniform response formatting.
func (h *CategorizationHandler) handleError(c *gin.Context, status int, message string, err error) {
	log.Printf("Error in categorization handler: %v", err)
	c.JSON(status, gin.H{
		"status":  "error",
		"message": message,
		"error":   err.Error(),
	})
}
