package entity

type Payment struct {
	OrderID   string
	Amount    int64
	FirstName string
	LastName  string
	Email     string
}

func NewPayment(orderID string, amount int64, firstName, lastName, email string) *Payment {
	return &Payment{
		OrderID:   orderID,
		Amount:    amount,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}
