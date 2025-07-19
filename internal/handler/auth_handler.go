package handler

import (
	"context"
	"net/http"

	"github.com/NeginSal/otp-auth-api/internal/dto"
	"github.com/NeginSal/otp-auth-api/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

// POST /send-otp
func (h *AuthHandler) SendOTP(c *gin.Context) {
	var request dto.SendOTPRequest

	if err := c.ShouldBindJSON(&request); err != nil || request.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number is required"})
		return
	}

	_, err := h.AuthService.SendOTP(context.Background(), request.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP has been sent"})
}

// SendOTP godoc
// @Summary      Send OTP to phone number
// @Description  This API sends a one-time password (OTP) to the user's phone number.
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request  body      dto.SendOTPRequest  true  "Phone number"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /send-otp [post]
// POST /verify-otp
func (h *AuthHandler) VerifyOTP(c *gin.Context) {
	var request dto.VerifyOTPRequest

	if err := c.ShouldBindJSON(&request); err != nil || request.Phone == "" || request.OTP == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number and OTP code are required"})
		return
	}

	token, err := h.AuthService.VerifyOTP(context.Background(), request.Phone, request.OTP)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
