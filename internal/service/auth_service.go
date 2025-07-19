package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/NeginSal/otp-auth-api/internal/model"
	"github.com/NeginSal/otp-auth-api/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	UserRepo *repository.UserRepository
	OTPRepo  *repository.OTPRepository
}

func NewAuthService(userRepo *repository.UserRepository, otpRepo *repository.OTPRepository) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
		OTPRepo:  otpRepo,
	}
}

// SendOTP generates a new OTP, saves it in DB, and logs it
func (s *AuthService) SendOTP(ctx context.Context, phone string) (string, error) {
	code := generateOTP()
	expirationMinutes, _ := strconv.Atoi(os.Getenv("OTP_EXPIRATION_MINUTES"))
	expiresAt := time.Now().Add(time.Duration(expirationMinutes) * time.Minute)

	otp := &model.OTPRequest{
		Phone:     phone,
		OTP:       code,
		ExpiresAt: expiresAt,
		Verified:  false,
	}

	err := s.OTPRepo.SaveOTP(ctx, otp)
	if err != nil {
		return "", err
	}

	fmt.Printf("📱 OTP for %s is: %s (expires in %d minutes)\n", phone, code, expirationMinutes)

	return code, nil
}

func (s *AuthService) VerifyOTP(ctx context.Context, phone string, code string) (string, error) {
	otp, err := s.OTPRepo.GetLatestByPhone(ctx, phone)
	if err != nil {
		return "", errors.New("OTP not found")
	}

	if otp.Verified {
		return "", errors.New("OTP already used")
	}

	if otp.ExpiresAt.Before(time.Now()) {
		return "", errors.New("OTP expired")
	}

	if otp.OTP != code {
		return "", errors.New("invalid OTP")
	}

	s.OTPRepo.MarkVerified(ctx, otp.ID)

	_, err = s.UserRepo.FindByPhone(ctx, phone)
	if err != nil {
		newUser := &model.User{
			Phone: phone,
		}
		err = s.UserRepo.CreateUser(ctx, newUser)
		if err != nil {
			return "", errors.New("failed to create user")
		}
	}

	token, err := generateJWT(phone)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

// generateOTP generates a random 5-digit OTP code
func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%05d", rand.Intn(100000))
}

// generateJWT creates a JWT token for the user
func generateJWT(phone string) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"phone": phone,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
