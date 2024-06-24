package router

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/bloomingbug/depublic/internal/http/handler"
	"github.com/bloomingbug/depublic/pkg/route"
)

const (
	Administrator = "Administrator"
	Buyer         = "Buyer"
)

var (
	allRoles  = []string{Administrator, Buyer}
	onlyAdmin = []string{Administrator}
	onlyBuyer = []string{Buyer}
)

func AppPublicRoutes(h map[string]interface{}) []*route.Route {
	return []*route.Route{
		{
			Method: http.MethodPost,
			Path:   "/request-otp",
			Handler: func(c echo.Context) error {
				return h["otp"].(*handler.OneTimePasswordHandler).Generate(c)
			},
		},
		{
			Method: http.MethodPost,
			Path:   "/verify-otp",
			Handler: func(c echo.Context) error {
				return h["token"].(*handler.TokenHandler).GenerateForRegister(c)
			},
		},
		{
			Method: http.MethodPost,
			Path:   "/auth/register",
			Handler: func(c echo.Context) error {
				return h["user"].(*handler.UserHandler).Registration(c)
			},
		},
		{
			Method: http.MethodPost,
			Path:   "/auth/login",
			Handler: func(c echo.Context) error {
				return h["user"].(*handler.UserHandler).Login(c)
			},
		},
		{
			Method: http.MethodPost,
			Path:   "/forgot-password",
			Handler: func(c echo.Context) error {
				return h["token"].(*handler.TokenHandler).GenerateForForgotPassword(c)
			},
		},
		{
			Method: http.MethodPost,
			Path:   "/reset-password",
			Handler: func(c echo.Context) error {
				return h["user"].(*handler.UserHandler).ResetPassword(c)
			},
		},
		{
			Method: http.MethodGet,
			Path:   "/events",
			Handler: func(c echo.Context) error {
				return h["event"].(*handler.EventHandler).GetAllEvent(c)
			},
		},
		{
			Method: http.MethodGet,
			Path:   "/events/:id",
			Handler: func(c echo.Context) error {
				return h["event"].(*handler.EventHandler).GetDetailEvent(c)
			},
		},
		{
			Method: http.MethodPost,
			Path:   "/payment",
			Handler: func(c echo.Context) error {
				return h["transaction"].(*handler.TransactionHandler).WebHookTransaction(c)
			},
		},
	}
}

func AppPrivateRoutes(h map[string]interface{}) []*route.Route {
	return []*route.Route{
		{
			Method: http.MethodGet,
			Path:   "/user/profile",
			Handler: func(c echo.Context) error {
				return h["user"].(*handler.UserHandler).Profile(c)
			},
			Roles: allRoles,
		},
		{
			Method: http.MethodGet,
			Path:   "/user/transactions",
			Handler: func(c echo.Context) error {
				return h["user"].(*handler.UserHandler).TransactionHistory(c)
			},
			Roles: allRoles,
		},
		{
			Method: http.MethodGet,
			Path:   "/user/notifications",
			Handler: func(c echo.Context) error {
				return h["user"].(*handler.UserHandler).Notifications(c)
			},
			Roles: allRoles,
		},
		{
			Method: http.MethodGet,
			Path:   "/user/notifications/:id",
			Handler: func(c echo.Context) error {
				return h["user"].(*handler.UserHandler).ReadNotification(c)
			},
			Roles: allRoles,
		},
		{
			Method: http.MethodPost,
			Path:   "/events/:id",
			Handler: func(c echo.Context) error {
				return h["transaction"].(*handler.TransactionHandler).CreateTransaction(c)
			},
			Roles: onlyBuyer,
		},
		{
			Method: http.MethodGet,
			Path:   "/admin/ticket",
			Handler: func(c echo.Context) error {
				return h["ticket"].(*handler.TicketHandler).UseTicket(c)
			},
			Roles: onlyAdmin,
		},
		{
			Method: http.MethodPost,
			Path:   "/admin/ticket",
			Handler: func(c echo.Context) error {
				return h["ticket"].(*handler.TicketHandler).UseTicket(c)
			},
			Roles: onlyAdmin,
		},
		{
			Method: http.MethodGet,
			Path:   "/admin/events",
			Handler: func(c echo.Context) error {
				return h["event"].(*handler.EventHandler).GetAllEvent(c)
			},
			Roles: onlyAdmin,
		},
		{
			Method: http.MethodGet,
			Path:   "/admin/events/:id",
			Handler: func(c echo.Context) error {
				return h["event"].(*handler.EventHandler).GetDetailEventWithTicket(c)
			},
			Roles: onlyAdmin,
		},
		{
			Method: http.MethodPost,
			Path:   "/admin/events",
			Handler: func(c echo.Context) error {
				return h["event"].(*handler.EventHandler).CreateEvent(c)
			},
			Roles: onlyAdmin,
		},
		{
			Method: http.MethodPut,
			Path:   "/admin/events/:id",
			Handler: func(c echo.Context) error {
				return h["event"].(*handler.EventHandler).EditEvent(c)
			},
			Roles: onlyAdmin,
		},
		{
			Method: http.MethodDelete,
			Path:   "/admin/events/:id",
			Handler: func(c echo.Context) error {
				return h["event"].(*handler.EventHandler).DeleteEvent(c)
			},
			Roles: onlyAdmin,
		},
	}
}
