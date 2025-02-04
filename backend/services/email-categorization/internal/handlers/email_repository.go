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

// EmailRepository provides methods to interact with the email-related data in the database.
// It encapsulates database access, making it easier to test and maintain.
type EmailRepository struct {
	// DB is the pooled database connection used for all queries.
	DB *pgxpool.Pool
}

// NewEmailRepository creates a new instance of EmailRepository with the given database connection pool.
func NewEmailRepository(db *pgxpool.Pool) *EmailRepository {
	return &EmailRepository{DB: db}
}

// SaveEmail inserts a new email record into the database. It first converts the in-memory Email model
// into a database-specific model structure. If successful, the email is stored along with a timestamp
// indicating when it was created.
func (r *EmailRepository) SaveEmail(ctx context.Context, email models.Email) error {
	// Convert from service-level model to database-level model
	emailDB := converter.FromServiceModel(email)

	query := `
		INSERT INTO emails (id, headers, subject, sender, recipients, body, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.DB.Exec(ctx, query,
		emailDB.ID,
		emailDB.Headers,
		emailDB.Subject,
		emailDB.Sender,
		emailDB.Recipients,
		emailDB.Body,
		time.Now(),
	)
	if err != nil {
		log.Printf("Failed to save email: %v", err)
		return err
	}

	log.Println("Email saved successfully")
	return nil
}

// SaveCategory stores a categorization record in the database. The categories are converted to JSON
// before insertion. The record includes the email ID, the categorized labels, the confidence score,
// and a timestamp of when it was created.
func (r *EmailRepository) SaveCategory(ctx context.Context, record db.CatgegoryRecord) error {
	query := `
		INSERT INTO categories (id, email_id, categories, confidence_score, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	// Serialize the categories map to JSON for storage
	categoriesJSON, err := json.Marshal(record.Categories)
	if err != nil {
		log.Printf("Failed to serialize categories: %v", err)
		return err
	}

	_, dbErr := r.DB.Exec(ctx, query,
		record.ID,
		record.EmailID,
		categoriesJSON,
		record.ConfidenceScore,
		time.Now(),
	)
	if dbErr != nil {
		log.Printf("Failed to save categorization record: %v", dbErr)
		return dbErr
	}

	log.Println("Categorization record saved successfully.")
	return nil
}

// GetEmails retrieves all email records from the database and converts them into service-level Email models.
// It returns a slice of Emails and an error if something goes wrong during query or row scanning.
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
		err := rows.Scan(
			&emailDB.ID,
			&emailDB.Headers,
			&emailDB.Subject,
			&emailDB.Sender,
			&emailDB.Recipients,
			&emailDB.Body,
			&emailDB.CreatedAt,
		)
		if err != nil {
			log.Printf("Failed to scan email: %v", err)
			continue
		}

		// Convert the database model back to a service-level model.
		emails = append(emails, converter.ToServiceModel(&emailDB))
	}

	return emails, nil
}
