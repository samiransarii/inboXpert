package handlers

import (
	"context"

	"github.com/samiransarii/inboXpert/services/email-categorization/internal/models"
)

// EmailCategorizer defines the interface for handling email categorization requests.
// Implementations of this interface should provide methods to process both single and
// multiple email categorization tasks, leveraging any necessary services or data sources.
type EmailCategorizer interface {
	// CategorizeEmails takes a context and a single EmailRequest, processes it (e.g., by
	// calling an ML service for categorization), and returns an EmailResponse containing
	// the categorized result.
	CategorizeEmails(ctx context.Context, req *models.EmailRequest) (*models.EmailResponse, error)

	// BatchCategorizeEmails accepts a context and a BatchEmailRequest containing multiple
	// emails, processes them concurrently or in bulk, and returns a BatchEmailResponse
	// with all the categorization results. This method allows for efficient handling of
	// multiple emails in a single request.
	BatchCategorizeEmails(ctx context.Context, batch *models.BatchEmailRequest) (*models.BatchEmailResponse, error)
}
