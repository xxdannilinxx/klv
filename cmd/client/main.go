package main

import (
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	ccpb "github.com/xxdannilinxx/klv/proto/gen/ccpb"
)

func main() {
	l := log.New(os.Stdout, "klv-api - ", log.LstdFlags)

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("[CLIENT] Did not connect: %s.", err)
	}
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
