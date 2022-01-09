package utils

import (
	"github.com/go-playground/validator"
)

// Validate incoming requests
func IsInt(id int64) (bool, error) {
	validate := validator.New()
	err := validate.Var(id, "required")
	return err == nil, err
}
