package handlers

import (
	"context"

	"github.com/samiransarii/inboXpert/services/email-categorization/internal/models"
)

type EmailCategorizer interface {
	CategorizeEmails(ctx context.Context, req *models.EmailRequest) (*models.EmailResponse, error)
}

type BatchEmailCategorizer interface {
	BatchCategorizeEmails(ctx context.Context, batch *models.BatchEmailRequest) (*models.BatchEmailResponse, error)
}
