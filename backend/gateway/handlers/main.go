package handlers

import "github.com/samiransarii/inboXpert/common/utils"

var (
	CATEGORIZE_SERVICE_URL      = utils.GetEnv("CATEGORIZE_SERVICE", "https://localhost/50051")
	SPAM_FILTER_SERVICE_URL     = utils.GetEnv("SPAM_FILTER_SERVICE", "https://localhost/3002")
	PRIORITY_FILTER_SERVICE_URL = utils.GetEnv("PRIORITY_FILTER_SERVICE", "https://localhost/3003")
)
