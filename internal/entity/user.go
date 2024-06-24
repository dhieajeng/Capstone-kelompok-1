package entity

import (
	"time"

	"github.com/google/uuid"
)

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)

type Role string

const (
	Admin Role = "Administrator"
	Buyer Role = "Buyer"
)

type User struct {
	ID            uuid.UUID       `json:"id"`
	Name          string          `json:"name"`
	Email         string          `json:"email"`
	Password      string          `json:"-"`
	Role          Role            `json:"role"`
	Phone         *string         `json:"phone,omitempty"`
	Address       *string         `json:"address,omitempty"`
	Avatar        *string         `json:"avatar,omitempty"`
	Birthdate     *time.Time      `json:"birthdate,omitempty"`
	Gender        *Gender         `json:"gender,omitempty"`
	Notifications *[]Notification `gorm:"foreignKey:UserID;references:ID" json:"notifications,omitempty"`
	Transactions  *[]Transaction  `json:"transactions,omitempty"`
	Auditable
}

func NewUser(name, email, password string, phone, address, avatar *string, birthdate *time.Time, gender *Gender, role Role) *User {
	return &User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Password:  password,
		Role:      role,
		Phone:     phone,
		Address:   address,
		Avatar:    avatar,
		Birthdate: birthdate,
		Gender:    gender,
	}
}

func ChangePassword(id uuid.UUID, password string) *User {
	return &User{
		ID:        id,
		Password:  password,
		Auditable: UpdateAuditable(),
	}
}
