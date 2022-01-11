package cryptocurrency

import (
	"log"
	"os"
	"testing"

	"github.com/xxdannilinxx/klv/pgsql"
	ccpb "github.com/xxdannilinxx/klv/proto/gen/ccpb"
	"github.com/xxdannilinxx/klv/utils"
	"golang.org/x/net/context"
)

var (
	config utils.Config = utils.Config{
		PORT:              os.Getenv("PORT"),
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		POSTGRES_PORT:     os.Getenv("POSTGRES_PORT"),
	}
)

func BenchmarkGetMostVotedCryptoCurrency(b *testing.B) {
	l := log.New(os.Stdout, "klv-api-benchmark - ", log.LstdFlags)

	dbConn := pgsql.ConnectDB(config)
	repository := NewCryptoCurrencyRepository(dbConn)
	s := NewCryptoCurrencyService(l, repository)

	for i := 0; i < b.N; i++ {
		s.GetMostVotedCryptoCurrency(context.Background(), &ccpb.GetMostVotedCryptoCurrencyRequest{})
	}
}
