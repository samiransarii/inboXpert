package utils

import (
	"log"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

// GetEnv retrieves the value of the environment variable specified by `key`,
//
// If the environment variable is set, the function returns its value.
// If the variable is not set, it returns the provided `fallback` value instead.
//
// Parameters:
//
//	-key: The name of the environment variable to look up.
//	-fallback: The value to return if the environment variable is not set.
//
// Returns:
//
//	The value of the environment variable if it exists, or the `fallback` value if not.
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}

// ProxyToService forwards an incoming HTTP request to a specified target URL.
//
// This function acts as a reverse proxy, taking an HTTP request contecxt from Gin
// an redirecting it to the target URL provided. It parses the target URL, checks for
// any errors, and sets up a single-host reverse proxy to handle the request.
//
// Parameters:
//   - c: Gin context carrying the HTTP request and response.
//   - target: String URL of the service to which the request should be proxied.
func ProxyToService(c *gin.Context, target string) {
	remote, err := url.Parse(target)

	if err != nil {
		log.Fatalf("Invalid request: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(c.Writer, c.Request)
}
