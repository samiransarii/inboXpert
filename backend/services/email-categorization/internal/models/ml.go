package models

type MLRequest struct {
	ID        string
	Subject   string
	Body      string
	Sender    string
	Recipents []string
	Headers   map[string]string
	Priority  string
}

type MLResponse struct {
	ID              string
	Category        string
	ConfidenceScore float32
	Alternatives    []Alternative
	Error           string
}
