package config

import "github.com/go-playground/validator/v10"

// get validator instance
func NewValidator() *validator.Validate {
	return validator.New()
}
