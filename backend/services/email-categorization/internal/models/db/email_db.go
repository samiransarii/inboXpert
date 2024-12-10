package db

import "time"

// Email represents the structure of the emails table
type EmailDB struct {
	ID         string    `db:"id"`         // Primary Key (UUID)
	Sender     string    `db:"sender"`     // Email sender
	Subject    string    `db:"subject"`    // Email subject
	Body       string    `db:"content"`    // Email content
	Recipients string    `db:"recipients"` // Email recipients
	Headers    string    `db:"headers"`    // Stringified JSON for headers
	CreatedAt  time.Time `db:"created_at"` // Timestamp
}

// CatgegoryRecord represents a categorization result
type CatgegoryRecord struct {
	ID              string    `db:"id"`       // UUID
	EmailID         string    `db:"email_id"` // UUID
	Categories      string    `db:"categories"`
	ConfidenceScore float32   `db:"confidence_score"`
	CreatedAt       time.Time `db:"created_at"`
}
