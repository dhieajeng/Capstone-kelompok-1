package builder

import (
	"github.com/bloomingbug/depublic/internal/http/handler"
	"github.com/bloomingbug/depublic/internal/http/router"
	"github.com/bloomingbug/depublic/internal/repository"
	"github.com/bloomingbug/depublic/internal/service"
	"github.com/bloomingbug/depublic/pkg/jwt_token"
	"github.com/bloomingbug/depublic/pkg/payment"
	"github.com/bloomingbug/depublic/pkg/route"
	"github.com/bloomingbug/depublic/pkg/scheduler"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

func BuildAppPublicRoutes(db *gorm.DB, redisDB *redis.Pool, jwtToken jwt_token.JwtToken, scheduler scheduler.Scheduler, paymentGateway payment.PaymentGateway) []*route.Route {
	handlers := make(map[string]interface{})

	otpRepository := repository.NewOneTimePasswordRepository(db)
	otpService := service.NewOneTimePasswordService(otpRepository, scheduler)
	otpHandler := handler.NewOneTimePasswordHandler(otpService)
	handlers["otp"] = &otpHandler

	tokenRepository := repository.NewTokenRepository(db)
	userRepository := repository.NewUserRepository(db)

	tokenService := service.NewTokenService(otpRepository, tokenRepository, userRepository, scheduler)
	tokenHandler := handler.NewTokenHandler(tokenService)
	handlers["token"] = &tokenHandler

	eventRepository := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepository)
	eventHandler := handler.NewEventHandler(eventService)
	handlers["event"] = &eventHandler

	timetableRepository := repository.NewTimetableRepository(db)
	timetableService := service.NewTimetableService(timetableRepository)

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)

	ticketRepository := repository.NewTicketRepository(db)
	ticketService := service.NewTicketService(ticketRepository)
	ticketHandler := handler.NewTicketHandler(ticketService, transactionService)
	handlers["ticket"] = &ticketHandler

	paymentService := service.NewPaymentService(paymentGateway)

	notificationRepository := repository.NewNotificationRepository(db)
	notificationService := service.NewNotificationService(notificationRepository)

	transactionHandler := handler.NewTransactionHandler(eventService, timetableService, transactionService, ticketService, notificationService, paymentService)
	handlers["transaction"] = &transactionHandler

	userService := service.NewUserService(tokenRepository, userRepository, jwtToken)
	userHandler := handler.NewUserHandler(userService, transactionService, notificationService)
	handlers["user"] = &userHandler

	return router.AppPublicRoutes(handlers)
}

func BuildAppPrivateRoutes(db *gorm.DB, redisDB *redis.Pool, jwtToken jwt_token.JwtToken, scheduler scheduler.Scheduler, paymentGateway payment.PaymentGateway) []*route.Route {
	handlers := make(map[string]interface{})

	tokenRepository := repository.NewTokenRepository(db)
	userRepository := repository.NewUserRepository(db)

	eventRepository := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepository)
	eventHandler := handler.NewEventHandler(eventService)
	handlers["event"] = &eventHandler

	timetableRepository := repository.NewTimetableRepository(db)
	timetableService := service.NewTimetableService(timetableRepository)

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)

	ticketRepository := repository.NewTicketRepository(db)
	ticketService := service.NewTicketService(ticketRepository)
	ticketHandler := handler.NewTicketHandler(ticketService, transactionService)
	handlers["ticket"] = &ticketHandler

	paymentService := service.NewPaymentService(paymentGateway)

	notificationRepository := repository.NewNotificationRepository(db)
	notificationService := service.NewNotificationService(notificationRepository)

	transactionHandler := handler.NewTransactionHandler(eventService, timetableService, transactionService, ticketService, notificationService, paymentService)
	handlers["transaction"] = &transactionHandler

	userService := service.NewUserService(tokenRepository, userRepository, jwtToken)
	userHandler := handler.NewUserHandler(userService, transactionService, notificationService)
	handlers["user"] = &userHandler

	return router.AppPrivateRoutes(handlers)
}
