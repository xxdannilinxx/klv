package cryptocurrency

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/xxdannilinxx/klv/pgsql"
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
	dbConn               *sql.DB                  = pgsql.ConnectDB(Config)
	repository           CryptoCurrencyRepository = NewCryptoCurrencyRepository(dbConn)
	fakeCryptoRepository *CryptoCurrency          = GenerateFakeCrypto("REPO")
)

func GenerateFakeCrypto(id string) *CryptoCurrency {
	return &CryptoCurrency{
		Name:  fmt.Sprintf("%s%d", id, rand.Intn(99999)),
		Token: fmt.Sprintf("%s%d", id, rand.Intn(99999)),
	}
}

func TestSave(t *testing.T) {
	result, err := repository.Save(fakeCryptoRepository)

	fakeCryptoRepository = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCryptoRepository.Validate())
}

func TestSaveDuplicate(t *testing.T) {
	_, err := repository.Save(fakeCryptoRepository)

	pqErr := err.(*pq.Error)

	assert.NotNil(t, err)
	assert.Equal(t, "23505", string(pqErr.Code))
}

func TestUpdate(t *testing.T) {
	result, err := repository.Update(fakeCryptoRepository)

	fakeCryptoRepository = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCryptoRepository.Validate())
}

func TestUpdateNotExists(t *testing.T) {
	_, err := repository.Update(&CryptoCurrency{
		Name:  fakeCryptoRepository.Name,
		Token: fakeCryptoRepository.Token,
	})

	assert.NotNil(t, err)
	assert.EqualError(t, err, "sql: no rows in result set")
}

func TestGetById(t *testing.T) {
	result, err := repository.GetById(fakeCryptoRepository.Id)

	fakeCryptoRepository = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCryptoRepository.Validate())
}

func TestGetMostVoted(t *testing.T) {
	result, err := repository.GetMostVoted()

	cryptoMostVoted := result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, cryptoMostVoted.Validate())
}

func TestUpVote(t *testing.T) {
	result, err := repository.UpVote(fakeCryptoRepository.Id)

	fakeCryptoRepository = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCryptoRepository.Validate())
}

func DownVote(t *testing.T) {
	result, err := repository.DownVote(fakeCryptoRepository.Id)

	fakeCryptoRepository = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCryptoRepository.Validate())
}

func TestDelete(t *testing.T) {
	result, err := repository.Delete(fakeCryptoRepository.Id)

	fakeCryptoRepository = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCryptoRepository.Validate())

	fakeCryptoRepository = GenerateFakeCrypto("REPO")
}

func TestAllMethodsDependsGetById(t *testing.T) {
	_, err := repository.GetById(0)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "sql: no rows in result set")
}
