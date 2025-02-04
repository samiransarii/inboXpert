package models

// Email represents the essential properties of an email.
type Email struct {
	ID         string
	Subject    string
	Body       string
	Sender     string
	Recipients []string
	Headers    map[string]string
}

// CategoryResult contains categorization information for a single email.
type CategoryResult struct {
	EmailID         string
	Categories      []string
	ConfidenceScore float32
}

// Alternative is used to store an additional category and confidence score for comparison.
type Alternative struct {
	Category        string
	ConfidenceScore float32
}

// EmailRequest defines a request structure for processing a single email,
// possibly including metadata that may influence how the request is handled.
type EmailRequest struct {
	Email    Email
	Metadata map[string]string
}

// EmailResponse holds the result of a single email categorization request,
// indicating success, the result, and any associated error message.
type EmailResponse struct {
	Result  CategoryResult
	Error   string
	Success bool
}

// BatchEmailRequest represents a request to process multiple emails in a single batch.
// BatchSize and Priority can influence processing behavior.
type BatchEmailRequest struct {
	Emails    []Email
	BatchSize int
	Priority  string
}

// BatchEmailResponse provides aggregated results after processing a batch of emails,
// including the total number processed, how many failed, and any error messages.
type BatchEmailResponse struct {
	Results []CategoryResult
	Failed  int32
	Total   int32
	Errors  []string
}
