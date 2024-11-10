package main

import (
	"fmt"
	"log"

	_ "github.com/joho/godotenv/autoload"
	utils "github.com/samiransarii/inboXpert/backend/shared"

	"github.com/gin-gonic/gin"
)

var (
	GATEWAY_PORT            = utils.GetEnv("GATEWAY_PORT", "8080")
	CATEGORIZE_SERVICE      = utils.GetEnv("CATEGORIZE_SERVICE", "https://localhost/3001")
	SPAM_FILTER_SERVICE     = utils.GetEnv("SPAM_FILTER_SERVICE", "https://localhost/3002")
	PRIORITY_FILTER_SERVICE = utils.GetEnv("PRIORITY_FILTER_SERVICE", "https://localhost/3003")
)

func main() {
	gateway := gin.Default()

	// nil for now, update with the proxies of microservices when microservices starts running
	gateway.SetTrustedProxies([]string{"CATEGORIZE_SERVICE", "SPAM_FILTER_SERVICE", "PRIORITY_FILTER_SERVICE", "https://localhost/3001", "https://localhost/3002", "https://localhost/3003"})

	gateway.GET("/categorize", func(c *gin.Context) {
		utils.ProxyToService(c, CATEGORIZE_SERVICE)
	})

	gateway.GET("/spam-filter", func(c *gin.Context) {
		utils.ProxyToService(c, SPAM_FILTER_SERVICE)
	})

	gateway.GET("/priority", func(c *gin.Context) {
		utils.ProxyToService(c, PRIORITY_FILTER_SERVICE)
	})

	err := gateway.Run("localhost:" + GATEWAY_PORT)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	fmt.Printf("API Gateway is running on port: %s\n", GATEWAY_PORT)
}
