package entity

import (
	"github.com/google/uuid"
	"time"
)

type Ticket struct {
	ID            uuid.UUID    `json:"id"`
	NoTicket      string       `json:"no_ticket"`
	Name          string       `json:"name"`
	PersonalNo    *string      `json:"personal_no,omitempty"`
	Birthdate     time.Time    `json:"birthdate"`
	Phone         *string      `json:"phone,omitempty"`
	Email         string       `json:"email"`
	Gender        *Gender      `json:"gender,omitempty"`
	Price         int64        `json:"price"`
	IsValid       bool         `json:"is_valid"`
	TransactionID uuid.UUID    `json:"-"`
	Transaction   *Transaction `gorm:"foreignKey:TransactionID;references:ID" json:"transaction,omitempty"`
	TimetableID   uuid.UUID    `json:"-"`
	Timetable     *Timetable   `gorm:"foreignKey:TimetableID;references:ID" json:"timetable,omitempty"`
	Auditable
}

type NewTicketParams struct {
	Name          string
	NoTicket      string
	PersonalNo    *string
	Birthdate     time.Time
	Phone         *string
	Email         string
	Gender        *Gender
	Price         int64
	TimetableID   uuid.UUID
	TransactionID uuid.UUID
}

func NewTicket(params NewTicketParams) *Ticket {
	return &Ticket{
		ID:            uuid.New(),
		NoTicket:      params.NoTicket,
		Name:          params.Name,
		PersonalNo:    params.PersonalNo,
		Birthdate:     params.Birthdate,
		Phone:         params.Phone,
		Email:         params.Email,
		Gender:        params.Gender,
		Price:         params.Price,
		IsValid:       true,
		TimetableID:   params.TimetableID,
		TransactionID: params.TransactionID,
	}
}

func UsedTicket(id uuid.UUID) *Ticket {
	return &Ticket{
		ID:      id,
		IsValid: false,
	}
}
