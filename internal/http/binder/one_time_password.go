package binder

type GenerateOTPRequest struct {
	Email string `json:"email" validate:"required,email"`
}
