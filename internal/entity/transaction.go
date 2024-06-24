package entity

import (
	"github.com/google/uuid"
)

type Transaction struct {
	ID         uuid.UUID `json:"id"`
	Invoice    string    `json:"invoice"`
	GrandTotal int64     `json:"grand_total"`
	SnapToken  *string   `json:"snap_token,omitempty"`
	Status     string    `json:"status"`
	UserID     uuid.UUID `json:"-"`
	User       *User     `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Tickets    []Ticket  `gorm:"foreignKey:TransactionID;references:ID" json:"tickets,omitempty"`
	Auditable
}

type NewTransactionParams struct {
	UserID     uuid.UUID
	Invoice    string
	GrandTotal int64
}

func NewTransaction(params NewTransactionParams) *Transaction {
	return &Transaction{
		ID:         uuid.New(),
		Invoice:    params.Invoice,
		GrandTotal: params.GrandTotal,
		SnapToken:  nil,
		UserID:     params.UserID,
	}
}

type UpdateTransactionParams struct {
	ID         uuid.UUID
	Invoice    *string
	GrandTotal *int64
	SnapToken  *string
	Status     *string
}

func UpdateTransaction(params UpdateTransactionParams) *Transaction {
	transaction := &Transaction{
		ID: params.ID,
	}

	if params.Invoice != nil {
		transaction.Invoice = *params.Invoice
	}
	if params.GrandTotal != nil {
		transaction.GrandTotal = *params.GrandTotal
	}
	if params.SnapToken != nil {
		transaction.SnapToken = params.SnapToken
	}
	if params.Status != nil {
		transaction.Status = *params.Status
	}

	return transaction
}

func UpdateStatusTransaction(id uuid.UUID, status string) *Transaction {
	return &Transaction{
		ID:     id,
		Status: status,
	}
}
