package payment

import (
	"github.com/bloomingbug/depublic/configs"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"strings"
)

type PaymentGateway struct {
	cfg *configs.Config
}

func (p *PaymentGateway) SnapClient() *snap.Client {
	var snapClient = new(snap.Client)
	if strings.ToLower(p.cfg.Env) == "production" {
		snapClient.New(p.cfg.Midtrans.ServerKey, midtrans.Production)
	} else {
		snapClient.New(p.cfg.Midtrans.ServerKey, midtrans.Sandbox)
	}

	return snapClient
}

func InitPaymentGateway(cfg *configs.Config) PaymentGateway {
	return PaymentGateway{cfg}
}
