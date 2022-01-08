package cryptocurrency

import (
	"fmt"
	"log"

	ccpb "github.com/xxdannilinxx/klv/proto/gen/ccpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"golang.org/x/net/context"
)

type Server struct {
	l *log.Logger
	ccpb.UnimplementedCryptoCurrencyServer
}

func NewCryptoCurrency(l *log.Logger) *Server {
	return &Server{l, ccpb.UnimplementedCryptoCurrencyServer{}}
}

func (s *Server) ListCryptoCurrencys(ctx context.Context, r *ccpb.ListCryptoCurrencysRequest) (*ccpb.ListCryptoCurrencysResponse, error) {
	s.l.Printf("[CRYPTOCURRENCY] ListCryptoCurrencys: %s", r)

	return &ccpb.ListCryptoCurrencysResponse{}, nil
}

func (s *Server) GetCryptoCurrency(ctx context.Context, r *ccpb.GetCryptoCurrencyRequest) (*ccpb.GetCryptoCurrencyResponse, error) {
	s.l.Printf("[CRYPTOCURRENCY] GetCryptoCurrency: %s", r)

	return &ccpb.GetCryptoCurrencyResponse{}, nil
}

func (s *Server) CreateCryptoCurrency(ctx context.Context, r *ccpb.CreateCryptoCurrencyRequest) (*ccpb.CreateCryptoCurrencyResponse, error) {
	s.l.Printf("[CRYPTOCURRENCY] CreateCryptoCurrency: %s", r)

	cr := &CryptoCurrency{
		Id:    r.Cryptocurrency.Id,
		Name:  r.Cryptocurrency.Name,
		Token: r.Cryptocurrency.Token,
		Votes: r.Cryptocurrency.Votes,
	}

	err := cr.Validate()
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("[CRYPTOCURRENCY] Internal error: %v.", err),
		)
	}

	return &ccpb.CreateCryptoCurrencyResponse{Cryptocurrency: &ccpb.CryptoCurrencyStruct{
		Id:    cr.Id,
		Name:  cr.Name,
		Token: cr.Token,
		Votes: cr.Votes,
	}}, nil
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
