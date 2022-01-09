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

// This function is used for return most voted cryptocurrency
//
// Example of call:
// grpcurl -plaintext -d '' localhost:8090 CryptoCurrency.GetMostVotedCryptoCurrency
func (s *Server) GetMostVotedCryptoCurrency(ctx context.Context, r *ccpb.GetMostVotedCryptoCurrencyRequest) (*ccpb.GetMostVotedCryptoCurrencyResponse, error) {
	s.l.Printf("[CRYPTOCURRENCY] GetMostVotedCryptoCurrency: %s", r)

	repository := &CryptoCurrencyRepository{s.db}
	crypto, err := repository.GetMostVoted()
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("[CRYPTOCURRENCY] Currency not found: %s.", err),
		)
	}

	return &ccpb.GetMostVotedCryptoCurrencyResponse{Cryptocurrency: &ccpb.CryptoCurrencyStruct{
		Id:    crypto.Id,
		Name:  crypto.Name,
		Token: crypto.Token,
		Votes: crypto.Votes,
	}}, nil
}

// This function is used for get specific cryptocurrency
//
// Example of call:
// grpcurl -plaintext -d '{"id": 1}' localhost:8090 CryptoCurrency.GetCryptoCurrency
func (s *Server) GetCryptoCurrency(ctx context.Context, r *ccpb.GetCryptoCurrencyRequest) (*ccpb.GetCryptoCurrencyResponse, error) {
	s.l.Printf("[CRYPTOCURRENCY] GetCryptoCurrency: %s", r)

	_, err := utils.IsInt(r.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("[CRYPTOCURRENCY] Invalid field validation: %s", err))
	}

	repository := &CryptoCurrencyRepository{s.db}
	crypto, err := repository.Get(r.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("[CRYPTOCURRENCY] Currency not found: %s.", err),
		)
	}

	return &ccpb.GetCryptoCurrencyResponse{Cryptocurrency: &ccpb.CryptoCurrencyStruct{
		Id:    crypto.Id,
		Name:  crypto.Name,
		Token: crypto.Token,
		Votes: crypto.Votes,
	}}, nil
}

// This function is used for create a new cryptocurrency
//
// Example of call:
// grpcurl -plaintext -d '{ "cryptocurrency": { "name": "Klever", "token": "KLV" } }' localhost:8090 CryptoCurrency.CreateCryptoCurrency
func (s *Server) CreateCryptoCurrency(ctx context.Context, r *ccpb.CreateCryptoCurrencyRequest) (*ccpb.CreateCryptoCurrencyResponse, error) {
	s.l.Printf("[CRYPTOCURRENCY] CreateCryptoCurrency: %s", r)

	crypto := &CryptoCurrency{
		Name:  r.Cryptocurrency.Name,
		Token: r.Cryptocurrency.Token,
	}

	err := crypto.Validate()
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("[CRYPTOCURRENCY] Internal server error: %s.", err),
		)
	}

	repository := &CryptoCurrencyRepository{s.db}
	crypto, err = repository.Insert(crypto)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("[CRYPTOCURRENCY] Internal server error: %s.", err),
		)
	}

	return &ccpb.CreateCryptoCurrencyResponse{Cryptocurrency: &ccpb.CryptoCurrencyStruct{
		Id:    crypto.Id,
		Name:  crypto.Name,
		Token: crypto.Token,
		Votes: crypto.Votes,
	}}, nil
}

// This function is used for update existent cryptocurrency
//
// Example of call:
// grpcurl -plaintext -d '{ "cryptocurrency": { "id": 1, "name": "Klever", "token": "KLV" } }' localhost:8090 CryptoCurrency.UpdateCryptoCurrency
func (s *Server) UpdateCryptoCurrency(ctx context.Context, r *ccpb.UpdateCryptoCurrencyRequest) (*ccpb.UpdateCryptoCurrencyResponse, error) {
	s.l.Printf("[CRYPTOCURRENCY] UpdateCryptoCurrency: %s", r)

	crypto := &CryptoCurrency{
		Id:    r.Cryptocurrency.Id,
		Name:  r.Cryptocurrency.Name,
		Token: r.Cryptocurrency.Token,
	}

	err := crypto.Validate()
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("[CRYPTOCURRENCY] Internal server error: %s.", err),
		)
	}

	repository := &CryptoCurrencyRepository{s.db}
	crypto, err = repository.Update(crypto)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("[CRYPTOCURRENCY] Internal server error: %s.", err),
		)
	}

	return &ccpb.UpdateCryptoCurrencyResponse{Cryptocurrency: &ccpb.CryptoCurrencyStruct{
		Id:    crypto.Id,
		Name:  crypto.Name,
		Token: crypto.Token,
		Votes: crypto.Votes,
	}}, nil
}

// This function is used for delete cryptocurrency
//
// Example of call:
// grpcurl -plaintext -d '{"id": 1}' localhost:8090 CryptoCurrency.DeleteCryptoCurrency
func (s *Server) DeleteCryptoCurrency(ctx context.Context, r *ccpb.DeleteCryptoCurrencyRequest) (*ccpb.DeleteCryptoCurrencyResponse, error) {
	s.l.Printf("[CRYPTOCURRENCY] DeleteCryptoCurrency: %s", r)

	_, err := utils.IsInt(r.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("[CRYPTOCURRENCY] Invalid field validation: %s", err))
	}

	repository := &CryptoCurrencyRepository{s.db}
	_, err = repository.Delete(r.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("[CRYPTOCURRENCY] Currency not deleted: %s.", err),
		)
	}

	return &ccpb.DeleteCryptoCurrencyResponse{Success: true}, nil
}

// This function is used for return most voted cryptocurrency
//
// Example of call:
// grpcurl -plaintext -d '{"id": 1}' localhost:8090 CryptoCurrency.UpVote
func (s *Server) UpVote(ctx context.Context, r *ccpb.UpVoteRequest) (*ccpb.UpVoteResponse, error) {
	s.l.Printf("[CRYPTOCURRENCY] UpVote: %s", r)

	_, err := utils.IsInt(r.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("[CRYPTOCURRENCY] Invalid field validation: %s", err))
	}

	repository := &CryptoCurrencyRepository{s.db}
	_, err = repository.UpVote(r.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("[CRYPTOCURRENCY] Currency not voted: %s.", err),
		)
	}

	return &ccpb.UpVoteResponse{Success: true}, nil
}

// This function is used for return most voted cryptocurrency
//
// Example of call:
// grpcurl -plaintext -d '{"id": 1}' localhost:8090 CryptoCurrency.DownVote
func (s *Server) DownVote(ctx context.Context, r *ccpb.DownVoteRequest) (*ccpb.DownVoteResponse, error) {
	s.l.Printf("[CRYPTOCURRENCY] DownVote: %s", r)

	_, err := utils.IsInt(r.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("[CRYPTOCURRENCY] Invalid field validation: %s", err))
	}

	repository := &CryptoCurrencyRepository{s.db}
	_, err = repository.DownVote(r.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("[CRYPTOCURRENCY] Currency not voted: %s.", err),
		)
	}

	return &ccpb.DownVoteResponse{Success: true}, nil
}
