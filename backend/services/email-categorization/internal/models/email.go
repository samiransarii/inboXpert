package models

type Email struct {
	ID         string
	Subject    string
	Body       string
	Sender     string
	Recipients []string
	Headers    map[string]string
}

type CategoryResult struct {
	EmailID         string
	Categories      []string
	ConfidenceScore float32
}

type Alternative struct {
	Category        string
	ConfidenceScore float32
}

type EmailRequest struct {
	Email    Email
	Metadata map[string]string
}

type EmailResponse struct {
	Result  CategoryResult
	Error   string
	Success bool
}

type BatchEmailRequest struct {
	Emails    []Email
	BatchSize int
	Priority  string
}

type BatchEmailResponse struct {
	Results []CategoryResult
	Failed  int32
	Total   int32
	Errors  []string
}
