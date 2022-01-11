package entity

import (
	"github.com/go-playground/validator"
)

type CryptoCurrency struct {
	Id    int64  `json:"id" bson:"id" db:"id"`
	Name  string `json:"name" bson:"name" db:"name" validate:"required"`
	Token string `json:"token" bson:"token" db:"token" validate:"required"`
	Votes int64  `json:"votes" bson:"votes" db:"votes"`
}

// Validate cryptocurrency structure
func (c *CryptoCurrency) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
