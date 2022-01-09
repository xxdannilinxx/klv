package utils

import (
	"github.com/go-playground/validator"
)

func IsInt(id int64) (bool, error) {
	validate := validator.New()
	err := validate.Var(id, "required")
	return err == nil, err
}
