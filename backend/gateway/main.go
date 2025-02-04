package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload" // Automatically load environment variables from a .env file, if present.

	utils "github.com/samiransarii/inboXpert/common/utils"
	handlers "github.com/samiransarii/inboXpert/gateway/handlers"
)

// GATEWAY_PORT defines the port on which the API Gateway will listen.
// If not provided as an environment variable, it defaults to "8080".
var GATEWAY_PORT = utils.GetEnv("GATEWAY_PORT", "8080")

func main() {
	// Create a new Gin engine instance for routing HTTP requests.
	gateway := gin.Default()

	// Add a middleware to manage CORS (Cross-Origin Resource Sharing) headers,
	// allowing requests from the specified Chrome extension.
	gateway.Use(func(ctx *gin.Context) {
		// Set CORS headers for requests originating from the Chrome extension.
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "chrome-extension://limgejhkljoadkclajoeijlojaanpebl")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

		// Handle preflight OPTIONS requests by returning a 204 No Content status.
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	})

	// Specify that no specific trusted reverse proxy addresses are known.
	// This is part of a security measure to avoid IP spoofing through proxy headers.
	gateway.SetTrustedProxies(nil)

	// Create instances of request handlers for different services.
	categorizationHandler := handlers.NewCategorizationHandler()

	// Define the routes exposed by the API Gateway.
	// POST /categorize: Routes incoming categorization requests to the CategorizationHandler.
	gateway.POST("/categorize", categorizationHandler.Handle)

	// Future routes for spam filtering and priority filtering could be added here:
	// gateway.GET("/spam-filter", spamFilterHandler)
	// gateway.GET("/priority", priorityFilterHandler)

	// Start the API Gateway server on the configured port, listening for incoming requests.
	err := gateway.Run("localhost:" + GATEWAY_PORT)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	fmt.Printf("API Gateway is running on port: %s\n", GATEWAY_PORT)
}
