package converter

import (
	"encoding/json"
	"log"

	"github.com/samiransarii/inboXpert/services/email-categorization/internal/models"
	"github.com/samiransarii/inboXpert/services/email-categorization/internal/models/db"
)

func ToServiceModel(e *db.EmailDB) models.Email {
	var headers map[string]string
	var recipients []string

	if e.Headers != "" {
		err := json.Unmarshal([]byte(e.Headers), &headers)
		if err != nil {
			log.Printf("Failed to unmarshal headers: %v", err)
		}
	}

	if e.Recipients != "" {
		err := json.Unmarshal([]byte(e.Recipients), &recipients)
		if err != nil {
			log.Printf("Failed to deserialize recipients list: %v", err)
		}
	}

	return models.Email{
		ID:         e.ID,
		Sender:     e.Sender,
		Subject:    e.Subject,
		Body:       e.Body,
		Recipients: recipients,
		Headers:    headers,
	}
}

func FromServiceModel(email models.Email) db.EmailDB {
	headerJSON, err := json.Marshal(email.Headers)
	if err != nil {
		log.Printf("Failed to marshal headers: %v", err)
		headerJSON = []byte("{}")
	}

	recipientsJSON, err := json.Marshal(email.Recipients)
	if err != nil {
		log.Printf("Failed to serialize recipients: %v", err)
		recipientsJSON = []byte("{}")
	}

	return db.EmailDB{
		ID:         email.ID,
		Headers:    string(headerJSON),
		Subject:    email.Subject,
		Sender:     email.Sender,
		Recipients: string(recipientsJSON),
		Body:       email.Body,
	}
}
