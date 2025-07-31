// @title           OTP Auth API
// @version         1.0
// @description     API for sending and verifying OTP codes via phone
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8080
// @BasePath  /
// @schemes   http
package main

import (
	"context"
	"log"

	"github.com/NeginSal/otp-auth-api/internal/config"
	"github.com/NeginSal/otp-auth-api/internal/routes"
	"github.com/gin-gonic/gin"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to MongoDB
	client := config.ConnectMongoDB()
	defer client.Disconnect(context.Background())

	// Setup router
	router := gin.Default()
	routes.SetupRoutes(router, client)

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := config.GetEnv("PORT", "8080")
	log.Println("Server running on port " + port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
