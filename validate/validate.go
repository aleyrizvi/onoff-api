package validate

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Validator struct {
	Validator *validator.Validate
}

func NewValidatorInstance() *Validator {
	v := validator.New()

	return &Validator{
		Validator: v,
	}
}

func (v *Validator) GetInstance() *validator.Validate {
	return v.Validator
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, v.FormatErrors(err.(validator.ValidationErrors)))
	}
	return nil
}

func (v *Validator) FormatErrors(err validator.ValidationErrors) []string {
	var errors []string

	for _, e := range err {
		var err string
		switch e.Tag() {
		case "required":
			err = fmt.Sprintf("%s is required", e.Field())
		case "email":
			err = fmt.Sprintf("%s must be a valid email", e.Field())
		case "gte":
			err = fmt.Sprintf("%s must be greater than or equal to %s", e.Field(), e.Param())
		default:
			err = fmt.Sprintf("%s invalid error", e.Field())
		}
		errors = append(errors, err)
	}

	return errors
}
