package handler

import (
	"net/http"

	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/http/form_validator"
	"github.com/bloomingbug/depublic/internal/service"
	"github.com/bloomingbug/depublic/pkg/response"
	"github.com/labstack/echo/v4"
)

type TokenHandler struct {
	tokenService service.TokenService
}

func (h *TokenHandler) GenerateForRegister(c echo.Context) error {
	req := new(binder.VerifyOTPRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(
			http.StatusBadRequest,
			false,
			form_validator.ValidatorErrors(err)))
	}

	token, err := h.tokenService.GenerateTokenRegistration(c, req.OTPCode, req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error(http.StatusInternalServerError, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK, true, "Success", echo.Map{
		"token":      token.ID,
		"email":      token.Email,
		"action":     token.Action,
		"expires_at": token.ExpiresAt,
	}))
}

func (h *TokenHandler) GenerateForForgotPassword(c echo.Context) error {
	req := new(binder.ForgotPasswordRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusOK, response.Error(
			http.StatusOK,
			true,
			form_validator.ValidatorErrors(err)))
	}

	token, err := h.tokenService.GenerateTokenForgotPassword(c, req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error(http.StatusInternalServerError, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK, true, "Success", echo.Map{
		"email":      token.Email,
		"action":     token.Action,
		"expires_at": token.ExpiresAt,
	}))
}

func NewTokenHandler(tokenService service.TokenService) TokenHandler {
	return TokenHandler{
		tokenService: tokenService,
	}
}
