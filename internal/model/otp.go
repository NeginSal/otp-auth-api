package model

import (
	"time"
)

type OTPRequest struct {
	ID        string    `bson:"_id,omitempty"` 
	Phone     string    `bson:"phone"`          
	OTP       string    `bson:"otp"`           
	ExpiresAt time.Time `bson:"expires_at"`    
	Verified  bool      `bson:"verified"`      
	CreatedAt time.Time `bson:"created_at"`    
}