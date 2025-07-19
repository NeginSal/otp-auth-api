package main

import (
	"context"
	"log"

	"github.com/NeginSal/otp-auth-api/internal/config"
	"github.com/NeginSal/otp-auth-api/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server")
	}
}
