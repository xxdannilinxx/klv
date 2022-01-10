package cryptocurrency

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/xxdannilinxx/klv/db"
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

var (
	fakeCrypto *CryptoCurrency = &CryptoCurrency{
		Name:  fmt.Sprintf("Klever%d", rand.Intn(99999)),
		Token: fmt.Sprintf("KLV%d", rand.Intn(99999)),
	}
)

func TestSave(t *testing.T) {
	db := db.ConnectDB(Config)
	repository := NewCryptoCurrencyRepository(db)
	result, err := repository.Save(fakeCrypto)

	fakeCrypto = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCrypto.Validate())
}

func TestSaveDuplicate(t *testing.T) {
	db := db.ConnectDB(Config)
	repository := NewCryptoCurrencyRepository(db)
	_, err := repository.Save(fakeCrypto)

	pqErr := err.(*pq.Error)

	assert.NotNil(t, err)
	assert.Equal(t, "23505", string(pqErr.Code))
}

func TestUpdate(t *testing.T) {
	db := db.ConnectDB(Config)
	repository := NewCryptoCurrencyRepository(db)
	result, err := repository.Update(fakeCrypto)

	fakeCrypto = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCrypto.Validate())
}

func TestUpdateNotExists(t *testing.T) {
	db := db.ConnectDB(Config)
	repository := NewCryptoCurrencyRepository(db)
	_, err := repository.Update(&CryptoCurrency{
		Name:  fakeCrypto.Name,
		Token: fakeCrypto.Token,
	})

	assert.NotNil(t, err)
	assert.EqualError(t, err, "sql: no rows in result set")
}

func TestGetById(t *testing.T) {
	db := db.ConnectDB(Config)
	repository := NewCryptoCurrencyRepository(db)
	result, err := repository.GetById(fakeCrypto.Id)

	fakeCrypto = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCrypto.Validate())
}

func TestGetMostVoted(t *testing.T) {
	db := db.ConnectDB(Config)
	repository := NewCryptoCurrencyRepository(db)
	result, err := repository.GetMostVoted()

	cryptoMostVoted := result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, cryptoMostVoted.Validate())
}

func TestUpVote(t *testing.T) {
	db := db.ConnectDB(Config)
	repository := NewCryptoCurrencyRepository(db)
	result, err := repository.UpVote(fakeCrypto.Id)

	fakeCrypto = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCrypto.Validate())
}

func DownVote(t *testing.T) {
	db := db.ConnectDB(Config)
	repository := NewCryptoCurrencyRepository(db)
	result, err := repository.DownVote(fakeCrypto.Id)

	fakeCrypto = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCrypto.Validate())
}

func TestDelete(t *testing.T) {
	db := db.ConnectDB(Config)
	repository := NewCryptoCurrencyRepository(db)
	result, err := repository.Delete(fakeCrypto.Id)

	fakeCrypto = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCrypto.Validate())

	fakeCrypto = &CryptoCurrency{}
}

func TestAllMethodsDependsGetById(t *testing.T) {
	db := db.ConnectDB(Config)
	repository := NewCryptoCurrencyRepository(db)
	_, err := repository.GetById(0)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "sql: no rows in result set")
}
