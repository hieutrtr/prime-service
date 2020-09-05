package router

import (
	"gopkg.in/go-playground/validator.v9"
)

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

type Validator struct {
	validator *validator.Validate
}

func (f *Validator) Validate(i interface{}) error {
	return f.validator.Struct(i)
}
