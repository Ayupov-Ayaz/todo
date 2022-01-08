package validator

import "github.com/go-playground/validator/v10"

type BasicValidator struct {
	val *validator.Validate
}

func NewBasicValidator() *BasicValidator {
	return &BasicValidator{
		val: validator.New(),
	}
}

func (v *BasicValidator) Struct(s interface{}) error {
	return v.val.Struct(s)
}
