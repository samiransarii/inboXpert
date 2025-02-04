package db

import "time"

// EmailDB represents the database schema for storing emails.
type EmailDB struct {
	ID         string    `db:"id"`         // Unique identifier for the email (UUID).
	Sender     string    `db:"sender"`     // The email sender address.
	Subject    string    `db:"subject"`    // The subject line of the email.
	Body       string    `db:"content"`    // The body/content of the email.
	Recipients string    `db:"recipients"` // A comma-separated list of recipient email addresses.
	Headers    string    `db:"headers"`    // JSON-encoded string of email headers.
	CreatedAt  time.Time `db:"created_at"` // Timestamp indicating when the email was stored.
}

// CatgegoryRecord represents a record of categorization results for a given email.
type CatgegoryRecord struct {
	ID              string    `db:"id"`              // Unique identifier for this record (UUID).
	EmailID         string    `db:"email_id"`        // The associated email's unique identifier.
	Categories      string    `db:"categories"`      // JSON-encoded categorization results.
	ConfidenceScore float32   `db:"confidence_score"` // The model’s confidence score for the categorization.
	CreatedAt       time.Time `db:"created_at"`      // Timestamp indicating when the record was created.
}
