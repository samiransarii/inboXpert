package handlers

import "github.com/samiransarii/inboXpert/common/utils"

// These variables define the base URLs for various backend services.
// These URLs can be customized via environment variables. If the corresponding
// environment variable is not set, each variable defaults to a local endpoint.
var (
	// CATEGORIZE_SERVICE_URL is the endpoint for the Categorization service.
	// It defaults to "https://localhost/50051" if the CATEGORIZE_SERVICE environment variable is not set.
	CATEGORIZE_SERVICE_URL = utils.GetEnv("CATEGORIZE_SERVICE", "https://localhost/50051")

	// SPAM_FILTER_SERVICE_URL is the endpoint for the Spam Filter service.
	// It defaults to "https://localhost/3002" if the SPAM_FILTER_SERVICE environment variable is not set.
	SPAM_FILTER_SERVICE_URL = utils.GetEnv("SPAM_FILTER_SERVICE", "https://localhost/3002")

	// PRIORITY_FILTER_SERVICE_URL is the endpoint for the Priority service.
	// It defaults to "https://localhost/3003" if the PRIORITY_FILTER_SERVICE environment variable is not set.
	PRIORITY_FILTER_SERVICE_URL = utils.GetEnv("PRIORITY_FILTER_SERVICE", "https://localhost/3003")
)
