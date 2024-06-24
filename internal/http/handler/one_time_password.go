package handler

import (
	"net/http"

	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/http/form_validator"
	"github.com/bloomingbug/depublic/internal/service"
	"github.com/bloomingbug/depublic/pkg/response"
	"github.com/labstack/echo/v4"
)

type OneTimePasswordHandler struct {
	otpService service.OneTimePasswordService
}

func (h *OneTimePasswordHandler) Generate(c echo.Context) error {
	req := new(binder.GenerateOTPRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(
			http.StatusBadRequest,
			false,
			form_validator.ValidatorErrors(err)))
	}

	otp, err := h.otpService.GenerateForRegister(c.Request().Context(), req.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK, true, "berhasil mengirim otp ke email", echo.Map{
		"email": otp.Email,
	}))
}

func NewOneTimePasswordHandler(otpService service.OneTimePasswordService) OneTimePasswordHandler {
	return OneTimePasswordHandler{otpService: otpService}
}
