package main

import (
	"github.com/bloomingbug/depublic/configs"
	"github.com/bloomingbug/depublic/internal/builder"
	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/http/form_validator"
	"github.com/bloomingbug/depublic/pkg/cache"
	"github.com/bloomingbug/depublic/pkg/jwt_token"
	paymentGateway "github.com/bloomingbug/depublic/pkg/payment"
	"github.com/bloomingbug/depublic/pkg/postgres"
	"github.com/bloomingbug/depublic/pkg/scheduler"
	"github.com/bloomingbug/depublic/pkg/server"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := configs.NewConfig(".env")
	checkError(err)

	pg, err := postgres.InitProgres(&cfg.Postgres)
	checkError(err)

	redis := cache.InitCache(&cfg.Redis)

	jwtToken := jwt_token.NewJwtToken(cfg.JWT.SecretKey)
	sch := scheduler.NewScheduler(redis, cfg.Namespace)

	payment := paymentGateway.InitPaymentGateway(cfg)

	publicRoutes := builder.BuildAppPublicRoutes(pg, redis, jwtToken, sch, payment)
	privateRoutes := builder.BuildAppPrivateRoutes(pg, redis, jwtToken, sch, payment)

	echoBinder := &echo.DefaultBinder{}
	formValidator := form_validator.NewFormValidator()
	customBinder := binder.NewBinder(echoBinder, formValidator)

	srv := server.NewServer(cfg, customBinder, publicRoutes, privateRoutes)
	srv.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
