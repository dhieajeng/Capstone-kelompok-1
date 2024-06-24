package service

import (
	"context"
	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/pkg/payment"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type paymentService struct {
	snapClient payment.PaymentGateway
}

func (s *paymentService) CreateTransaction(c context.Context, payment *entity.Payment) (*string, error) {
	request := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  payment.OrderID,
			GrossAmt: payment.Amount,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: payment.FirstName,
			LName: payment.LastName,
			Email: payment.Email,
		},
	}

	snapUrl, err := s.snapClient.SnapClient().CreateTransaction(request)
	if err != nil {
		return nil, err
	}

	return &snapUrl.RedirectURL, nil
}

type PaymenService interface {
	CreateTransaction(c context.Context, payment *entity.Payment) (*string, error)
}

func NewPaymentService(snapClient payment.PaymentGateway) PaymenService {
	return &paymentService{snapClient: snapClient}
}
