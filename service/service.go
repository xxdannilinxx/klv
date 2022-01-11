package service

import (
	"fmt"
	"log"

	"github.com/xxdannilinxx/klv/entity"
	_ "github.com/xxdannilinxx/klv/entity"
	ccpb "github.com/xxdannilinxx/klv/proto/gen/ccpb"
	"github.com/xxdannilinxx/klv/repository"
	_ "github.com/xxdannilinxx/klv/repository"
	"github.com/xxdannilinxx/klv/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"golang.org/x/net/context"
)

type CryptoCurrencyService interface {
	GetMostVotedCryptoCurrency(ctx context.Context, r *ccpb.GetMostVotedCryptoCurrencyRequest) (*ccpb.GetMostVotedCryptoCurrencyResponse, error)
	GetCryptoCurrency(ctx context.Context, r *ccpb.GetCryptoCurrencyRequest) (*ccpb.GetCryptoCurrencyResponse, error)
	CreateCryptoCurrency(ctx context.Context, r *ccpb.CreateCryptoCurrencyRequest) (*ccpb.CreateCryptoCurrencyResponse, error)
	UpdateCryptoCurrency(ctx context.Context, r *ccpb.UpdateCryptoCurrencyRequest) (*ccpb.UpdateCryptoCurrencyResponse, error)
	DeleteCryptoCurrency(ctx context.Context, r *ccpb.DeleteCryptoCurrencyRequest) (*ccpb.DeleteCryptoCurrencyResponse, error)
	UpVote(ctx context.Context, r *ccpb.UpVoteRequest) (*ccpb.UpVoteResponse, error)
	DownVote(ctx context.Context, r *ccpb.DownVoteRequest) (*ccpb.DownVoteResponse, error)
}

type Server struct {
	l          *log.Logger
	repository repository.CryptoCurrencyRepository
	ccpb.UnimplementedCryptoCurrencyServer
}

// New cryptocurrency module office
func NewCryptoCurrencyService(l *log.Logger, repository repository.CryptoCurrencyRepository) *Server {
	return &Server{l, repository, ccpb.UnimplementedCryptoCurrencyServer{}}
}

// This function is used for return most voted cryptocurrency
//
// Example of call:
// grpcurl -plaintext -d '' localhost:8090 CryptoCurrency.GetMostVotedCryptoCurrency
func (s *Server) GetMostVotedCryptoCurrency(ctx context.Context, r *ccpb.GetMostVotedCryptoCurrencyRequest) (*ccpb.GetMostVotedCryptoCurrencyResponse, error) {
	s.l.Printf("[CRYPTOCURRENCY] GetMostVotedCryptoCurrency: %s", r)

	crypto, err := s.repository.GetMostVoted()
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

	crypto, err := s.repository.GetById(r.Id)

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

	crypto := &entity.CryptoCurrency{
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

	crypto, err = s.repository.Save(crypto)
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

	crypto := &entity.CryptoCurrency{
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

	crypto, err = s.repository.Update(crypto)
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

	_, err = s.repository.Delete(r.Id)

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

	_, err = s.repository.UpVote(r.Id)

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

	_, err = s.repository.DownVote(r.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("[CRYPTOCURRENCY] Currency not voted: %s.", err),
		)
	}

	return &ccpb.DownVoteResponse{Success: true}, nil
}
