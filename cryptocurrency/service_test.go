package cryptocurrency

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	ccpb "github.com/xxdannilinxx/klv/proto/gen/ccpb"
	"golang.org/x/net/context"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) GetMostVoted() (*CryptoCurrency, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*CryptoCurrency), args.Error(1)
}

func (mock *MockRepository) GetById(id int64) (*CryptoCurrency, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*CryptoCurrency), args.Error(1)
}

func (mock *MockRepository) Save(cc *CryptoCurrency) (*CryptoCurrency, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*CryptoCurrency), args.Error(1)
}

func (mock *MockRepository) Update(cc *CryptoCurrency) (*CryptoCurrency, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*CryptoCurrency), args.Error(1)
}

func (mock *MockRepository) Delete(id int64) (*CryptoCurrency, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*CryptoCurrency), args.Error(1)
}

func (mock *MockRepository) UpVote(id int64) (*CryptoCurrency, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*CryptoCurrency), args.Error(1)
}
func (mock *MockRepository) DownVote(id int64) (*CryptoCurrency, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*CryptoCurrency), args.Error(1)
}

var (
	l                 *log.Logger = log.New(os.Stdout, "klv-api-test - ", log.LstdFlags)
	fakeCryptoService             = GenerateFakeCrypto("SVC")
)

func TestGetMostVotedCryptoCurrency(t *testing.T) {
	mockRepo := new(MockRepository)
	s := NewCryptoCurrencyService(l, mockRepo)

	crypto := &CryptoCurrency{
		Name:  fakeCryptoService.Name,
		Token: fakeCryptoService.Token,
	}
	mockRepo.On("GetMostVoted").Return(crypto, nil)

	resp, err := s.GetMostVotedCryptoCurrency(context.Background(), &ccpb.GetMostVotedCryptoCurrencyRequest{})

	assert.NotNil(t, resp.Cryptocurrency.Id)
	assert.Equal(t, crypto.Name, resp.Cryptocurrency.Name)
	assert.Equal(t, crypto.Token, resp.Cryptocurrency.Token)
	assert.Nil(t, err)
}

func TestGetCryptoCurrency(t *testing.T) {
	mockRepo := new(MockRepository)
	s := NewCryptoCurrencyService(l, mockRepo)

	crypto := &CryptoCurrency{
		Id:    1,
		Name:  fakeCryptoService.Name,
		Token: fakeCryptoService.Token,
		Votes: 0,
	}
	mockRepo.On("GetById").Return(crypto, nil)

	resp, err := s.GetCryptoCurrency(context.Background(), &ccpb.GetCryptoCurrencyRequest{
		Id: crypto.Id,
	})

	assert.NotNil(t, crypto.Id, resp.Cryptocurrency.Id)
	assert.Equal(t, crypto.Name, resp.Cryptocurrency.Name)
	assert.Equal(t, crypto.Token, resp.Cryptocurrency.Token)
	assert.Nil(t, err)
}

func TestCreateCryptoCurrency(t *testing.T) {
	mockRepo := new(MockRepository)
	s := NewCryptoCurrencyService(l, mockRepo)

	crypto := &CryptoCurrency{
		Name:  fakeCryptoService.Name,
		Token: fakeCryptoService.Token,
		Votes: 0,
	}
	mockRepo.On("Save").Return(crypto, nil)

	resp, err := s.CreateCryptoCurrency(context.Background(), &ccpb.CreateCryptoCurrencyRequest{
		Cryptocurrency: &ccpb.CryptoCurrencyStruct{
			Name:  crypto.Name,
			Token: crypto.Token,
		},
	})

	assert.NotNil(t, resp.Cryptocurrency.Id)
	assert.Equal(t, crypto.Name, resp.Cryptocurrency.Name)
	assert.Equal(t, crypto.Token, resp.Cryptocurrency.Token)
	assert.Nil(t, err)
}

func TestCreateInvalidCryptoCurrency(t *testing.T) {
	mockRepo := new(MockRepository)
	s := NewCryptoCurrencyService(l, mockRepo)

	crypto := &CryptoCurrency{
		Name:  "",
		Token: "",
		Votes: 0,
	}
	mockRepo.On("Save").Return(crypto, nil)

	_, err := s.CreateCryptoCurrency(context.Background(), &ccpb.CreateCryptoCurrencyRequest{
		Cryptocurrency: &ccpb.CryptoCurrencyStruct{
			Name:  crypto.Name,
			Token: crypto.Token,
		},
	})

	assert.NotNil(t, err)
}

func TestUpdateCryptoCurrency(t *testing.T) {
	mockRepo := new(MockRepository)
	s := NewCryptoCurrencyService(l, mockRepo)

	crypto := &CryptoCurrency{
		Name:  fakeCryptoService.Name,
		Token: fakeCryptoService.Token,
		Votes: 0,
	}
	mockRepo.On("Update").Return(crypto, nil)

	resp, err := s.UpdateCryptoCurrency(context.Background(), &ccpb.UpdateCryptoCurrencyRequest{
		Cryptocurrency: &ccpb.CryptoCurrencyStruct{
			Name:  crypto.Name,
			Token: crypto.Token,
		},
	})

	assert.NotNil(t, resp.Cryptocurrency.Id)
	assert.Equal(t, crypto.Name, resp.Cryptocurrency.Name)
	assert.Equal(t, crypto.Token, resp.Cryptocurrency.Token)
	assert.Nil(t, err)
}

func TestDeleteCryptoCurrency(t *testing.T) {
	mockRepo := new(MockRepository)
	s := NewCryptoCurrencyService(l, mockRepo)

	crypto := &CryptoCurrency{
		Id:    1,
		Name:  fakeCryptoService.Name,
		Token: fakeCryptoService.Token,
		Votes: 0,
	}
	mockRepo.On("Delete").Return(crypto, nil)

	resp, err := s.DeleteCryptoCurrency(context.Background(), &ccpb.DeleteCryptoCurrencyRequest{
		Id: crypto.Id,
	})

	assert.True(t, resp.Success)
	assert.Nil(t, err)
}

func TestDeleteEmptyCryptoCurrency(t *testing.T) {
	mockRepo := new(MockRepository)
	s := NewCryptoCurrencyService(l, mockRepo)

	crypto := &CryptoCurrency{
		Id:    0,
		Name:  fakeCryptoService.Name,
		Token: fakeCryptoService.Token,
		Votes: 0,
	}
	mockRepo.On("Delete").Return(crypto, nil)

	_, err := s.DeleteCryptoCurrency(context.Background(), &ccpb.DeleteCryptoCurrencyRequest{
		Id: crypto.Id,
	})

	assert.NotNil(t, err)
}

func TestUpVoteCryptoCurrency(t *testing.T) {
	mockRepo := new(MockRepository)
	s := NewCryptoCurrencyService(l, mockRepo)

	crypto := &CryptoCurrency{
		Id:    1,
		Name:  fakeCryptoService.Name,
		Token: fakeCryptoService.Token,
		Votes: 0,
	}
	mockRepo.On("UpVote").Return(crypto, nil)

	resp, err := s.UpVote(context.Background(), &ccpb.UpVoteRequest{
		Id: crypto.Id,
	})

	assert.True(t, resp.Success)
	assert.Nil(t, err)
}

func TestDownVoteCryptoCurrency(t *testing.T) {
	mockRepo := new(MockRepository)
	s := NewCryptoCurrencyService(l, mockRepo)

	crypto := &CryptoCurrency{
		Id:    1,
		Name:  fakeCryptoService.Name,
		Token: fakeCryptoService.Token,
		Votes: 0,
	}
	mockRepo.On("DownVote").Return(crypto, nil)

	resp, err := s.DownVote(context.Background(), &ccpb.DownVoteRequest{
		Id: crypto.Id,
	})

	assert.True(t, resp.Success)
	assert.Nil(t, err)
}
