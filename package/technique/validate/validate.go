package validate

import (
	"github.com/go-playground/validator/v10"
)

type IValidator interface {
	Validate(i any) error
}
type GestGoValidator struct {
	Validator *validator.Validate
}

func (cv *GestGoValidator) Validate(i any) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func NewGestGoValidator(validator *validator.Validate) IValidator {
	return &GestGoValidator{
		Validator: validator,
	}
}
