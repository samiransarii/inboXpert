package models

// MLRequest represents the data sent to the machine learning service for categorization.
type MLRequest struct {
	ID        string
	Subject   string
	Body      string
	Sender    string
	Recipents []string
	Headers   map[string]string
	Priority  string
}

// MLResponse represents the response returned by the machine learning service,
// including the primary category, confidence score, and any alternative categorizations.
type MLResponse struct {
	ID              string
	Category        string
	ConfidenceScore float32
	Alternatives    []Alternative
	Error           string
}
