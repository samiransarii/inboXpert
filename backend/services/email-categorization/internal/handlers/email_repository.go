package handlers

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samiransarii/inboXpert/services/email-categorization/internal/models"
	"github.com/samiransarii/inboXpert/services/email-categorization/internal/models/db"
	"github.com/samiransarii/inboXpert/services/email-categorization/internal/utils/converter"
)

type EmailRepository struct {
	DB *pgxpool.Pool
}

func NewEmailRepository(db *pgxpool.Pool) *EmailRepository {
	return &EmailRepository{DB: db}
}

// SaveEmail stores an email in the database
func (r *EmailRepository) SaveEmail(ctx context.Context, email models.Email) error {
	// Convert service model to database model
	emailDB := converter.FromServiceModel(email)

	query := `
		INSERT INTO emails (id, headers, subject, sender, recipients, body, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.DB.Exec(ctx, query, emailDB.ID, emailDB.Headers, emailDB.Subject, emailDB.Sender, emailDB.Recipients, emailDB.Body, time.Now())
	if err != nil {
		log.Printf("Failed to save email: %v", err)
		return err
	}

	log.Println("Email saved successfully")
	return nil
}

// SaveCategory stores categorization results in the database
func (r *EmailRepository) SaveCategory(ctx context.Context, record db.CatgegoryRecord) error {
	query := `
		INSERT INTO categories (id, email_id, categories, confidence_score, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	categoriesJSON, err := json.Marshal(record.Categories)
	if err != nil {
		log.Printf("Failed to serialize categories: %v", err)
		return err
	}

	_, dbErr := r.DB.Exec(ctx, query, record.ID, record.EmailID, categoriesJSON, record.ConfidenceScore, time.Now())
	if dbErr != nil {
		log.Printf("Failed to save categorization record: %v", err)
		return err
	}

	log.Println("Categorization record saved successfully.")
	return nil
}

// GetEmails retrieves all emails from the database
func (r *EmailRepository) GetEmails(ctx context.Context) ([]models.Email, error) {
	query := `
		SELECT id, headers, subject, sender, recipients, body, created_at
		FROM emails
	`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		log.Printf("Failed to retrieve emails: %v", err)
		return nil, err
	}
	defer rows.Close()

	var emails []models.Email
	for rows.Next() {
		var emailDB db.EmailDB
		err := rows.Scan(&emailDB.ID, &emailDB.Headers, &emailDB.Subject, &emailDB.Sender, &emailDB.Recipients, &emailDB.Body, &emailDB.CreatedAt)
		if err != nil {
			log.Printf("Failed to scan email: %v", err)
			continue
		}

		emails = append(emails, converter.ToServiceModel(&emailDB))
	}

	return emails, nil
}
