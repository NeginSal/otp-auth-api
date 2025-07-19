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
	"github.com/joho/godotenv"

	_ "github.com/NeginSal/otp-auth-api/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	client := config.ConnectDB()
	defer client.Disconnect(context.TODO())

	router := gin.Default()
	routes.SetupRoutes(router, client)

	port := config.GetEnv("PORT", "8080")
	log.Println("Server running on port " + port)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server")
	}
}
