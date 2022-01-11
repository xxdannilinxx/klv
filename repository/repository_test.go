package repository

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/xxdannilinxx/klv/entity"
	"github.com/xxdannilinxx/klv/pgsql"
	"github.com/xxdannilinxx/klv/utils"
)

type UnitTestSuite struct {
	suite.Suite
}

var (
	Config               utils.Config
	dbConn               *sql.DB                  = pgsql.ConnectDB(Config)
	repository           CryptoCurrencyRepository = NewCryptoCurrencyRepository(dbConn)
	fakeCryptoRepository *entity.CryptoCurrency   = GenerateFakeCrypto("REPO")
)

func GenerateFakeCrypto(id string) *entity.CryptoCurrency {
	return &entity.CryptoCurrency{
		Name:  fmt.Sprintf("%s%d", id, rand.Intn(99999)),
		Token: fmt.Sprintf("%s%d", id, rand.Intn(99999)),
	}
}

func (s *UnitTestSuite) BeforeTest(suiteName, testName string) {
	err := godotenv.Load(".env")
	utils.CheckError(err)

	log.Print("--------")
	log.Print("--------")
	log.Print("--------")

	Config = utils.Config{
		PORT:              os.Getenv("PORT"),
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		POSTGRES_PORT:     os.Getenv("POSTGRES_PORT"),
	}
}

func (s *UnitTestSuite) TestSave(t *testing.T) {
	s.BeforeTest(".env", "TestSave")
	result, err := repository.Save(fakeCryptoRepository)

	fakeCryptoRepository = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCryptoRepository.Validate())
}

func (s *UnitTestSuite) TestSaveDuplicate(t *testing.T) {
	s.BeforeTest(".env", "TestSaveDuplicate")
	_, err := repository.Save(fakeCryptoRepository)

	pqErr := err.(*pq.Error)

	assert.NotNil(t, err)
	assert.Equal(t, "23505", string(pqErr.Code))
}

func (s *UnitTestSuite) TestUpdate(t *testing.T) {
	s.BeforeTest(".env", "TestUpdate")
	result, err := repository.Update(fakeCryptoRepository)

	fakeCryptoRepository = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCryptoRepository.Validate())
}

func (s *UnitTestSuite) TestUpdateNotExists(t *testing.T) {
	s.BeforeTest(".env", "TestUpdateNotExists")
	_, err := repository.Update(&entity.CryptoCurrency{
		Name:  fakeCryptoRepository.Name,
		Token: fakeCryptoRepository.Token,
	})

	assert.NotNil(t, err)
	assert.EqualError(t, err, "sql: no rows in result set")
}

func (s *UnitTestSuite) TestGetById(t *testing.T) {
	s.BeforeTest(".env", "TestGetById")
	result, err := repository.GetById(fakeCryptoRepository.Id)

	fakeCryptoRepository = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCryptoRepository.Validate())
}

func (s *UnitTestSuite) TestGetMostVoted(t *testing.T) {
	s.BeforeTest(".env", "TestGetMostVoted")
	result, err := repository.GetMostVoted()

	cryptoMostVoted := result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, cryptoMostVoted.Validate())
}

func (s *UnitTestSuite) TestUpVote(t *testing.T) {
	s.BeforeTest(".env", "TestUpVote")
	result, err := repository.UpVote(fakeCryptoRepository.Id)

	fakeCryptoRepository = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCryptoRepository.Validate())
}

func (s *UnitTestSuite) DownVote(t *testing.T) {
	s.BeforeTest(".env", "DownVote")
	result, err := repository.DownVote(fakeCryptoRepository.Id)

	fakeCryptoRepository = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCryptoRepository.Validate())
}

func (s *UnitTestSuite) TestDelete(t *testing.T) {
	s.BeforeTest(".env", "TestDelete")
	result, err := repository.Delete(fakeCryptoRepository.Id)

	fakeCryptoRepository = result

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Nil(t, fakeCryptoRepository.Validate())

	fakeCryptoRepository = GenerateFakeCrypto("REPO")
}

func (s *UnitTestSuite) TestAllMethodsDependsGetById(t *testing.T) {
	s.BeforeTest(".env", "TestAllMethodsDependsGetById")
	_, err := repository.GetById(0)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "sql: no rows in result set")
}
