package handlers

import (
	"github.com/gin-gonic/gin"
	utils "github.com/samiransarii/inboXpert/backend/common/utils"
)

var (
	CATEGORIZE_SERVICE_URL      = utils.GetEnv("CATEGORIZE_SERVICE", "https://localhost/3001")
	SPAM_FILTER_SERVICE_URL     = utils.GetEnv("SPAM_FILTER_SERVICE", "https://localhost/3002")
	PRIORITY_FILTER_SERVICE_URL = utils.GetEnv("PRIORITY_FILTER_SERVICE", "https://localhost/3003")
)

// CategorizationService proxies incoming categorize requests to the categorization service.
//
// This function serves as a hanlder for categorization request in the Gin framework.
// It recieves a context paramenter from Gin, and then proxies the request to the Categorization Service
// by calling utils.ProxyToService with the provided context and categorization service URL.
//
// Parameters:
//   - c: *gin.Context - the context for the current HTTP request, which includes request data.
//
// Usage:
//
//	Pass this function as a handler for the "/categorize" route to proxy requests
//	to the appropriate backend service.
func CategorizationService(c *gin.Context) {
	utils.ProxyToService(c, CATEGORIZE_SERVICE_URL)
}

// // SpamFilterService proxies incoming spam filter requests to the spam filtering service.
//
// This function handles spam filtering requests by utilizing the Gin context to capture
// request data. It calls utils.ProxyToService, passing the request context along with
// the URL of the spam filtering service to forward the request.
//
// Parameters:
//   - c: *gin.Context - the context for the current HTTP request, including request data.
//
// Usage:
//
//	Register this function as a handler for the spam filter route to enable proxying
//	requests to the backend spam filtering service.
func SpamFilterService(c *gin.Context) {
	utils.ProxyToService(c, SPAM_FILTER_SERVICE_URL)
}

// PriorityFilterService proxies incoming priority filter requests to the priority filtering service.
//
// This function manages priority filtering requests in the Gin framework by capturing
// request context data and forwarding the request to the Priority Filter Service via
// utils.ProxyToService.
//
// Parameters:
//   - c: *gin.Context - the context for the current HTTP request, providing access to request data.
//
// Usage:
//
//	Attach this function as a handler for the priority filter route to enable request
//	proxying to the backend priority filter service.
func PriorityFilterService(c *gin.Context) {
	utils.ProxyToService(c, PRIORITY_FILTER_SERVICE_URL)
}
