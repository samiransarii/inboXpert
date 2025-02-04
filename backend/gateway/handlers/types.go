package handlers

// EmailRequest represents the structure of an email being processed by the handlers.
// It includes various metadata such as subject, sender, recipients, and headers.
type EmailRequest struct {
	ID         string            `json:"id"`
	Subject    string            `json:"subject"`
	Body       string            `json:"body"`
	Sender     string            `json:"sender"`
	Recipients []string          `json:"recipients"`
	Headers    map[string]string `json:"headers"`
}

// CategorizeServiceRequest represents the payload sent to a categorization service.
// It contains a list of emails to be categorized.
type CategorizeServiceRequest struct {
	Emails []EmailRequest `json:"emails"`
}
