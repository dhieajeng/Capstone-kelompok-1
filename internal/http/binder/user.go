package binder

import "github.com/bloomingbug/depublic/internal/entity"

type RegisterRequest struct {
	Name                 string         `form:"name" json:"name" validate:"required,alphanum"`
	Email                string         `form:"email" json:"email" validate:"required,email"`
	Password             string         `form:"password" json:"password" validate:"required,min=8"`
	PasswordConfirmation string         `form:"password_confirmation" json:"password_confirmation" validate:"required,min=8,eqfield=Password"`
	Phone                *string        `form:"phone" json:"phone"`
	Avatar               *string        `form:"avatar" json:"avatar"`
	Address              *string        `form:"address" json:"address"`
	Birthdate            *string        `form:"birthdate" json:"birthdate"`
	Gender               *entity.Gender `form:"gender" json:"gender" validate:"required,oneof=M F"`
	Token                string         `form:"token" json:"token" validate:"required,uuid"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ResetPasswordRequest struct {
	Token                string `json:"token" validate:"required,uuid"`
	Password             string `json:"password" validate:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,min=8,eqfield=Password"`
}
