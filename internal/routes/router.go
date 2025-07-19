package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/NeginSal/otp-auth-api/internal/handler"
	"github.com/NeginSal/otp-auth-api/internal/repository"
	"github.com/NeginSal/otp-auth-api/internal/service"
)

func SetupRoutes(r *gin.Engine, db *mongo.Client) {
	// Init repos
	userRepo := repository.NewUserRepository(db)
	otpRepo := repository.NewOTPRepository(db)

	// Init service
	authService := service.NewAuthService(userRepo, otpRepo)

	// Init handler
	authHandler := handler.NewAuthHandler(authService)

	// Routes
	r.POST("/send-otp", authHandler.SendOTP)
	r.POST("/verify-otp", authHandler.VerifyOTP)
}
