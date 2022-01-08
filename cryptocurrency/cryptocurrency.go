package cryptocurrency

import (
	"encoding/json"
	"log"

	ccpb "github.com/xxdannilinxx/klv/proto/gen/ccpb"

	"golang.org/x/net/context"
)

type Server struct {
	ccpb.UnimplementedCryptoCurrencyServer
}

func (s *Server) ListCryptoCurrencys(ctx context.Context, r *ccpb.ListCryptoCurrencysRequest) (*ccpb.ListCryptoCurrencysResponse, error) {
	log.Printf("Receive message from client: %s", r)

	return &ccpb.ListCryptoCurrencysResponse{}, nil
}

func (s *Server) GetCryptoCurrency(ctx context.Context, r *ccpb.GetCryptoCurrencyRequest) (*ccpb.GetCryptoCurrencyResponse, error) {
	log.Printf("Receive message from client: %s", r)

	return &ccpb.GetCryptoCurrencyResponse{}, nil
}

func (s *Server) CreateCryptoCurrency(ctx context.Context, r *ccpb.CreateCryptoCurrencyRequest) (*ccpb.CreateCryptoCurrencyResponse, error) {
	log.Printf("Receive message from client: %s", r)

	// passar log pra cá

	var body []byte
	err := json.Unmarshal(body, r)
	if err != nil {
		log.Println("ERROR ", err)
	}

	// err = cur.Validate()
	// if err != nil {
	// 	log.Println("[CURRENCY] ERROR ", err)
	// 	return
	// }

	return &ccpb.CreateCryptoCurrencyResponse{}, nil
}

func (s *Server) UpdateCryptoCurrency(ctx context.Context, r *ccpb.UpdateCryptoCurrencyRequest) (*ccpb.UpdateCryptoCurrencyResponse, error) {
	log.Printf("Receive message from client: %s", r)

	return &ccpb.UpdateCryptoCurrencyResponse{}, nil
}

func (s *Server) DeleteCryptoCurrency(ctx context.Context, r *ccpb.DeleteCryptoCurrencyRequest) (*ccpb.DeleteCryptoCurrencyResponse, error) {
	log.Printf("Receive message from client: %s", r)

	return &ccpb.DeleteCryptoCurrencyResponse{}, nil
}

func (s *Server) UpVote(ctx context.Context, r *ccpb.UpVoteRequest) (*ccpb.UpVoteResponse, error) {
	log.Printf("Receive message from client: %s", r)

	return &ccpb.UpVoteResponse{}, nil
}

func (s *Server) DownVote(ctx context.Context, r *ccpb.DownVoteRequest) (*ccpb.DownVoteResponse, error) {
	log.Printf("Receive message from client: %s", r)

	return &ccpb.DownVoteResponse{}, nil
}

// log.Printf("Receive message from client: %s", in.Username)

// 	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%v", in.Username))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	body, readErr := ioutil.ReadAll(res.Body)
// 	if readErr != nil {
// 		log.Fatal(readErr)
// 	}

// 	usr := User{}
// 	jsonErr := json.Unmarshal(body, &usr)
// 	if jsonErr != nil {
// 		log.Fatal(jsonErr)
// 	}

// 	return &user.UserResponse{
// 		Id:        usr.ID,
// 		Name:      usr.Name,
// 		Username:  usr.Username,
// 		Avatarurl: usr.AvatarURL,
// 		Location:  usr.Location,
// 		Statistics: &user.Statistics{
// 			Followers: usr.Followers,
// 			Following: usr.Following,
// 			Repos:     usr.Repos,
// 			Gists:     usr.Gists,
// 		},
// 		ListURLs: []string{usr.URL, usr.StarredURL, usr.ReposURL},
// 	}, nil
