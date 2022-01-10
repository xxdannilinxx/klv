package cryptocurrency

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	ccpb "github.com/xxdannilinxx/klv/proto/gen/ccpb"
	"golang.org/x/net/context"
)

func TestCryptoCurrencyServer(t *testing.T) {
	l := log.New(os.Stdout, "klv-api-test - ", log.LstdFlags)

	s := NewCryptoCurrency(l, Repository)

	tests := []struct {
		crypto *CryptoCurrency
	}{
		{
			crypto: GenerateFakeCrypto("SRV"),
		},
	}
	for _, tt := range tests {
		ctx := context.Background()

		respCreate, err := s.CreateCryptoCurrency(ctx, &ccpb.CreateCryptoCurrencyRequest{
			Cryptocurrency: &ccpb.CryptoCurrencyStruct{
				Name:  tt.crypto.Name,
				Token: tt.crypto.Token,
			},
		})

		tt.crypto = &CryptoCurrency{
			Id:    respCreate.Cryptocurrency.Id,
			Name:  respCreate.Cryptocurrency.Name,
			Token: respCreate.Cryptocurrency.Token,
			Votes: respCreate.Cryptocurrency.Votes,
		}

		assert.NotNil(t, respCreate)
		assert.Nil(t, err)

		respUpdate, err := s.UpdateCryptoCurrency(ctx, &ccpb.UpdateCryptoCurrencyRequest{
			Cryptocurrency: &ccpb.CryptoCurrencyStruct{
				Id:    tt.crypto.Id,
				Name:  tt.crypto.Name,
				Token: tt.crypto.Token,
			},
		})

		tt.crypto = &CryptoCurrency{
			Id:    respCreate.Cryptocurrency.Id,
			Name:  respCreate.Cryptocurrency.Name,
			Token: respCreate.Cryptocurrency.Token,
			Votes: respCreate.Cryptocurrency.Votes,
		}

		assert.NotNil(t, respUpdate)
		assert.Nil(t, err)

		respGetById, err := s.GetCryptoCurrency(ctx, &ccpb.GetCryptoCurrencyRequest{Id: tt.crypto.Id})

		assert.NotNil(t, respGetById)
		assert.Nil(t, err)

		respGetMost, err := s.GetMostVotedCryptoCurrency(ctx, &ccpb.GetMostVotedCryptoCurrencyRequest{})

		assert.NotNil(t, respGetMost)
		assert.Nil(t, err)

		respDelete, err := s.DeleteCryptoCurrency(ctx, &ccpb.DeleteCryptoCurrencyRequest{Id: tt.crypto.Id})

		tt.crypto = GenerateFakeCrypto("SRV")

		assert.NotNil(t, respDelete)
		assert.Nil(t, err)
	}
}
