package server

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// CustomValidator holds custom validator
type CustomValidator struct {
	V *validator.Validate
}

// NewValidator creates new custom validator
func NewValidator() *CustomValidator {
	V := validator.New()
	V.RegisterValidation("date", validateDate)
	return &CustomValidator{V}
}

// Validate validates the request
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.V.Struct(i)
}

func validateDate(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	re := regexp.MustCompile(`^\d{4}-\d{1,2}-\d{1,2}(T00:00:00Z)?$`)
	return re.MatchString(val)
}
