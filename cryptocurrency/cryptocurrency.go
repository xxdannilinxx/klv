package cryptocurrency

import (
	"database/sql"
	"fmt"
	"log"

	ccpb "github.com/xxdannilinxx/klv/proto/gen/ccpb"
	"github.com/xxdannilinxx/klv/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"golang.org/x/net/context"
)

type Server struct {
	l  *log.Logger
	db *sql.DB
	ccpb.UnimplementedCryptoCurrencyServer
}

func NewCryptoCurrency(l *log.Logger, db *sql.DB) *Server {
	return &Server{l, db, ccpb.UnimplementedCryptoCurrencyServer{}}
}

func (s *Server) GetMostVotedCryptoCurrency(ctx context.Context, r *ccpb.GetMostVotedCryptoCurrencyRequest) (*ccpb.GetMostVotedCryptoCurrencyResponse, error) {
	s.l.Printf("[CRYPTOCURRENCY] GetMostVotedCryptoCurrency: %s", r)

	return &ccpb.GetMostVotedCryptoCurrencyResponse{}, nil
}

func (s *Server) GetCryptoCurrency(ctx context.Context, r *ccpb.GetCryptoCurrencyRequest) (*ccpb.GetCryptoCurrencyResponse, error) {
	s.l.Printf("[CRYPTOCURRENCY] GetCryptoCurrency: %s", r)

	return &ccpb.GetCryptoCurrencyResponse{}, nil
}

// This function is used for create a new cryptocurrency
//
// Example of call:
// grpcurl -plaintext --msg-template -d '{ "cryptocurrency": { "name": "Klever", "token": "KLV" } }' localhost:8090 CryptoCurrency.CreateCryptoCurrency
func (s *Server) CreateCryptoCurrency(ctx context.Context, r *ccpb.CreateCryptoCurrencyRequest) (*ccpb.CreateCryptoCurrencyResponse, error) {
	s.l.Printf("[CRYPTOCURRENCY] CreateCryptoCurrency: %s", r)

	crypto := r.GetCryptocurrency()

	cc := &CryptoCurrency{
		Id:    0,
		Name:  crypto.GetName(),
		Token: crypto.GetToken(),
	}

	err := cc.Validate()
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("[CRYPTOCURRENCY] Internal server error: %v.", err),
		)
	}

	//
	//
	//
	rows, err := s.db.Query(`SELECT "name" FROM "cryptocurrencies"`)
	utils.CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var name string

		err = rows.Scan(&name)
		if err != nil {
			panic(err)
		}

		fmt.Println(name)
	}
	//
	//
	//

	crypto.Id = 10
	crypto.Votes = 100
	return &ccpb.CreateCryptoCurrencyResponse{Cryptocurrency: crypto}, nil
}

func (s *Server) UpdateCryptoCurrency(ctx context.Context, r *ccpb.UpdateCryptoCurrencyRequest) (*ccpb.UpdateCryptoCurrencyResponse, error) {
	s.l.Printf("[CRYPTOCURRENCY] UpdateCryptoCurrency: %s", r)

	return &ccpb.UpdateCryptoCurrencyResponse{}, nil
}

func (s *Server) DeleteCryptoCurrency(ctx context.Context, r *ccpb.DeleteCryptoCurrencyRequest) (*ccpb.DeleteCryptoCurrencyResponse, error) {
	s.l.Printf("[CRYPTOCURRENCY] DeleteCryptoCurrency: %s", r)

	return &ccpb.DeleteCryptoCurrencyResponse{}, nil
}

func (s *Server) UpVote(ctx context.Context, r *ccpb.UpVoteRequest) (*ccpb.UpVoteResponse, error) {
	s.l.Printf("[CRYPTOCURRENCY] UpVote: %s", r)

	return &ccpb.UpVoteResponse{}, nil
}

func (s *Server) DownVote(ctx context.Context, r *ccpb.DownVoteRequest) (*ccpb.DownVoteResponse, error) {
	s.l.Printf("[CRYPTOCURRENCY] DownVote: %s", r)

	return &ccpb.DownVoteResponse{}, nil
}
