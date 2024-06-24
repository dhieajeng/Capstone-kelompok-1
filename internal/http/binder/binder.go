package binder

import (
	"errors"
	"github.com/bloomingbug/depublic/internal/http/form_validator"
	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Binder struct {
	defaultBinder *echo.DefaultBinder
	*form_validator.FormValidator
}

func NewBinder(
	dbr *echo.DefaultBinder,
	vdr *form_validator.FormValidator) *Binder {
	return &Binder{dbr, vdr}
}

func (b *Binder) Bind(i interface{}, c echo.Context) error {
	if err := b.defaultBinder.Bind(i, c); err != nil {
		return err
	}

	if err := defaults.Set(i); err != nil {
		return err
	}

	if err := b.Validate(i); err != nil {
		var errs validator.ValidationErrors
		errors.As(err, &errs)
		return errs
	}

	return nil
}
