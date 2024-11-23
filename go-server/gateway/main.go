package main

import (
	"fmt"
	"log"

	_ "github.com/joho/godotenv/autoload"
	utils "github.com/samiransarii/inboXpert/backend/common/utils"
	handlers "github.com/samiransarii/inboXpert/gateway/handlers"

	"github.com/gin-gonic/gin"
)

var GATEWAY_PORT = utils.GetEnv("GATEWAY_PORT", "8080")

func main() {
	gateway := gin.Default()

	// nil for now, update with the proxies of microservices when microservices starts running
	gateway.SetTrustedProxies(nil)

	// create handlers
	categorizationHandler := handlers.NewCategorizationHandler()

	// Service Routes
	gateway.POST("/categorize", categorizationHandler.Handle)
	// gateway.GET("/spam-filter", handle.SpamFilterService)
	// gateway.GET("/priority", handle.PriorityFilterService)

	// Starting the gateway server
	err := gateway.Run("localhost:" + GATEWAY_PORT)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	fmt.Printf("API Gateway is running on port: %s\n", GATEWAY_PORT)
}
