package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	utils "github.com/samiransarii/inboXpert/backend/common/utils"
	pb "github.com/samiransarii/inboXpert/backend/proto/gen/services/email_categorization_service"
)

type CategorizationHandler struct {
	grpcManager *utils.GRPCClientManager
	serviceAddr string
	grpcTimeout time.Duration
}

func NewCategorizationHandler() *CategorizationHandler {
	return &CategorizationHandler{
		grpcManager: utils.GetGRPCClientManager(),
		serviceAddr: "localhost:50051",
		grpcTimeout: 10 * time.Second,
	}
}

func (h *CategorizationHandler) Handle(c *gin.Context) {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), h.grpcTimeout)
	defer cancel()

	// Parse and validate request
	var requestData CategorizeServiceRequest
	if err := h.parseRequest(c, &requestData); err != nil {
		h.handleError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Create connection
	conn, err := h.grpcManager.GetConnection(ctx, h.serviceAddr)
	if err != nil {
		h.handleError(c, http.StatusServiceUnavailable, "Failed to connect to service", err)
		return
	}

	// Create service client and request
	client := pb.NewEmailCategorizationServiceClient(conn)
	grpcRequest := h.createGRPCRequest(requestData)

	// Make gRPC call
	response, err := client.CategorizeEmails(ctx, grpcRequest)
	if err != nil {
		h.handleError(c, http.StatusInternalServerError, "Failed to categorize emails", err)
		return
	}

	// Return successful response
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   response,
	})
}

// Helper methods
func (h *CategorizationHandler) createGRPCRequest(req CategorizeServiceRequest) *pb.CategorizeRequest {
	grpcEmails := make([]*pb.Email, len(req.Emails))
	for i, email := range req.Emails {
		grpcEmails[i] = &pb.Email{
			Id:        email.ID,
			Subject:   email.Subject,
			Body:      email.Body,
			Sender:    email.Sender,
			Recipents: email.Recipents,
			Headers:   email.Headers,
		}
	}
	return &pb.CategorizeRequest{
		Emails: grpcEmails,
	}
}

func (h *CategorizationHandler) parseRequest(c *gin.Context, req *CategorizeServiceRequest) error {
	if err := c.ShouldBindJSON(req); err != nil {
		return err
	}
	if len(req.Emails) == 0 {
		return fmt.Errorf("at least one email is required")
	}
	return nil
}

func (h *CategorizationHandler) handleError(c *gin.Context, status int, message string, err error) {
	log.Printf("Error in categorization handler: %v", err)
	c.JSON(status, gin.H{
		"status":  "error",
		"message": message,
		"error":   err.Error(),
	})
}
