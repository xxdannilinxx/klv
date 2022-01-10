package cryptocurrency

import (
	"database/sql"
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
	DbConn     *sql.DB                   = db.ConnectDB(Config)
	Repository *CryptoCurrencyRepository = NewCryptoCurrencyRepository(DbConn)
	fakeCrypto *CryptoCurrency           = GenerateFakeCrypto("REPO")
)

func GenerateFakeCrypto(id string) *CryptoCurrency {
	return &CryptoCurrency{
		Name:  fmt.Sprintf("%s%d", id, rand.Intn(99999)),
		Token: fmt.Sprintf("%s%d", id, rand.Intn(99999)),
	}
}

func TestSave(t *testing.T) {
	result, err := Repository.Save(fakeCrypto)

	fakeCrypto = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCrypto.Validate())
}

func TestSaveDuplicate(t *testing.T) {
	_, err := Repository.Save(fakeCrypto)

	pqErr := err.(*pq.Error)

	assert.NotNil(t, err)
	assert.Equal(t, "23505", string(pqErr.Code))
}

func TestUpdate(t *testing.T) {
	result, err := Repository.Update(fakeCrypto)

	fakeCrypto = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCrypto.Validate())
}

func TestUpdateNotExists(t *testing.T) {
	_, err := Repository.Update(&CryptoCurrency{
		Name:  fakeCrypto.Name,
		Token: fakeCrypto.Token,
	})

	assert.NotNil(t, err)
	assert.EqualError(t, err, "sql: no rows in result set")
}

func TestGetById(t *testing.T) {
	result, err := Repository.GetById(fakeCrypto.Id)

	fakeCrypto = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCrypto.Validate())
}

func TestGetMostVoted(t *testing.T) {
	result, err := Repository.GetMostVoted()

	cryptoMostVoted := result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, cryptoMostVoted.Validate())
}

func TestUpVote(t *testing.T) {
	result, err := Repository.UpVote(fakeCrypto.Id)

	fakeCrypto = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCrypto.Validate())
}

func DownVote(t *testing.T) {
	result, err := Repository.DownVote(fakeCrypto.Id)

	fakeCrypto = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCrypto.Validate())
}

func TestDelete(t *testing.T) {
	result, err := Repository.Delete(fakeCrypto.Id)

	fakeCrypto = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCrypto.Validate())

	fakeCrypto = GenerateFakeCrypto("REPO")
}

func TestAllMethodsDependsGetById(t *testing.T) {
	_, err := Repository.GetById(0)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "sql: no rows in result set")
}
