package binder

import (
	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/google/uuid"
)

type TicketRequest struct {
	Name        string         `json:"name" validate:"required"`
	PersonalNo  *string        `json:"personal_no,omitempty"`
	Birthdate   string         `json:"birthdate" validate:"required"`
	Phone       *string        `json:"phone,omitempty"`
	Email       string         `json:"email,omitempty" validate:"required,email"`
	Gender      *entity.Gender `json:"gender" validate:"required,oneof=M F"`
	TimetableID uuid.UUID      `json:"timetable_id" validate:"required,uuid"`
}

type UseTicketRequest struct {
	NoTicket string `query:"no_ticket" json:"no_ticket" validate:"required"`
}
