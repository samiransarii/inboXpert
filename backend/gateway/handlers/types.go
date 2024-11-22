package handlers

type EmailRequest struct {
	ID        string            `json:"id"`
	Subject   string            `json:"subject"`
	Body      string            `json:"body"`
	Sender    string            `json:"sender"`
	Recipents []string          `json:"recipients"`
	Headers   map[string]string `json:"headers"`
}

type CategorizeServiceRequest struct {
	Emails []EmailRequest `json:"emails"`
}
