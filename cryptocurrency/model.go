package cryptocurrency

import (
	"github.com/go-playground/validator"
)

type CryptoCurrency struct {
	Id    int64  `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name" validate:"required"`
	Token string `json:"token" bson:"token" validate:"required"`
	Votes int64  `json:"votes" bson:"votes" validate:"required"`
}

func (c *CryptoCurrency) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
