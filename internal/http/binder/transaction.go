package binder

import "github.com/google/uuid"

type TransactionRequest struct {
	EventID uuid.UUID       `param:"id" json:"event_id" validate:"required,uuid"`
	Tickets []TicketRequest `json:"tickets" validate:"required"`
}
