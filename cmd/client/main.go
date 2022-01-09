package main

import (
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	ccpb "github.com/xxdannilinxx/klv/proto/gen/ccpb"
	"github.com/xxdannilinxx/klv/utils"
)

// Examples of server calls
func main() {
	l := log.New(os.Stdout, "klv-client - ", log.LstdFlags)

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8090", grpc.WithInsecure())
	utils.CheckError(err)

	defer conn.Close()

	cc := ccpb.NewCryptoCurrencyClient(conn)

	req := &ccpb.CreateCryptoCurrencyRequest{
		Cryptocurrency: &ccpb.CryptoCurrencyStruct{
			Id:    10,
			Name:  "Klever",
			Token: "KLV",
			Votes: 1000,
		},
	}
	response, err := cc.CreateCryptoCurrency(context.Background(), req)
	if err != nil {
		l.Fatalf("[CLIENT] Error when calling CreateCryptoCurrency: %s.", err)
	}

	l.Printf("[CLIENT] Response from server: %v.", response)
}
