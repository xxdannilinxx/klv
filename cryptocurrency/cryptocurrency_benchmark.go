package cryptocurrency

import (
	"log"
	"os"
	"testing"

	ccpb "github.com/xxdannilinxx/klv/proto/gen/ccpb"
	"golang.org/x/net/context"
)

func BenchmarkGetMostVotedCryptoCurrency(b *testing.B) {
	l := log.New(os.Stdout, "klv-api-benchmark - ", log.LstdFlags)

	s := NewCryptoCurrency(l, Repository)

	for i := 0; i < b.N; i++ {
		s.GetMostVotedCryptoCurrency(context.Background(), &ccpb.GetMostVotedCryptoCurrencyRequest{})
	}
}
