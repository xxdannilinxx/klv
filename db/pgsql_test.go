package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xxdannilinxx/klv/utils"
)

var (
	Config utils.Config = utils.Config{
		PORT:              os.Getenv("PORT"),
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		POSTGRES_PORT:     os.Getenv("POSTGRES_PORT"),
	}
)

func TestConnecDB(t *testing.T) {
	db := ConnectDB(Config)
	err := db.Ping()
	assert.Nil(t, err)
}
