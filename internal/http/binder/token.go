package binder

type VerifyOTPRequest struct {
	Email   string `json:"email" query:"email" validate:"required,email"`
	OTPCode string `json:"otp_code" query:"otp_code" validate:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}
