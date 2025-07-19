package dto

type SendOTPRequest struct {
	Phone string `json:"phone" example:"09123456789"`
}

type VerifyOTPRequest struct {
	Phone string `json:"phone" example:"09123456789"`
	OTP   string `json:"otp" example:"123456"`
}
