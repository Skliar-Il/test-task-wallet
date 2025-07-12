package validator

import "github.com/go-playground/validator/v10"

type Validator struct {
	Validator *validator.Validate
}

func (v Validator) Validate(out any) error {
	return v.Validator.Struct(out)
}
