package form_validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type FormValidator struct {
	formValidator *validator.Validate
}

func (fv *FormValidator) Validate(i interface{}) error {
	return fv.formValidator.Struct(i)
}

func NewFormValidator() *FormValidator {
	validate := validator.New(validator.WithRequiredStructEnabled())

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		fmt.Println(name)
		if name == "-" {
			return ""
		}
		return name
	})

	return &FormValidator{validate}
}

func ValidatorErrors(err error) map[string]string {
	fields := map[string]string{}

	var castedObject validator.ValidationErrors
	if errors.As(err, &castedObject) {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				fields[err.Field()] = fmt.Sprintf("field %s harus di isi", err.Field())
			case "oneof":
				fields[err.Field()] = fmt.Sprintf("field %s harus memiliki salah satu dari nilai: %s", err.Field(), err.Param())
			case "email":
				fields[err.Field()] = fmt.Sprintf("field %s harus berupa email", err.Field())
			case "length":
				fields[err.Field()] = fmt.Sprintf("field %s harus berisi %s karakter", err.Field(), err.Param())
			case "eqfield":
				fields[err.Field()] = fmt.Sprintf("field %s harus sama dengan field %s", err.Field(), err.Param())
			case "gte":
				fields[err.Field()] = fmt.Sprintf("field %s harus lebih dari atau sama dengan %s", err.Field(), err.Param())
			case "lte":
				fields[err.Field()] = fmt.Sprintf("field %s harus kurang dari atau sama dengan %s", err.Field(), err.Param())
			case "gt":
				fields[err.Field()] = fmt.Sprintf("field %s harus lebih dari %s", err.Field(), err.Param())
			case "lt":
				fields[err.Field()] = fmt.Sprintf("field %s harus kurang dari %s", err.Field(), err.Param())
			case "uuid":
				fields[err.Field()] = fmt.Sprintf("field %s harus berupa UUID", err.Field())
			default:
				fields[err.Field()] = fmt.Sprintf("%s error with tag %s should be %s", err.Field(), err.Tag(), err.Param())
			}
		}
	}

	return fields
}
