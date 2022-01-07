package currency

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-playground/validator"
)

type Currency struct {
	Id     int    `json:"id" bson:"id"`
	Name   string `json:"name" bson:"name" validate:"required"`
	Ticker string `json:"ticker" bson:"ticker" validate:"required"`
}

func (c *Currency) Unmarshal(r io.Reader) error {
	body, err := ioutil.ReadAll(r)
	err = json.Unmarshal(body, &c)
	return err
}

func (c *Currency) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
