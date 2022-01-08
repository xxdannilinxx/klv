package main

import (
	"log"
	"net"
	"os"

	"github.com/xxdannilinxx/klv/cryptocurrency"
	ccpb "github.com/xxdannilinxx/klv/proto/gen/ccpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	PORT              string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
}

var Config config

// bson
// fazer repository de alguma forma que de pra trocar o banco
// a regra de negocio fica no repository

// testes
// readme
// comentarios
func main() {
	l := log.New(os.Stdout, "klv-api - ", log.LstdFlags)

	Config.PORT = os.Getenv("PORT")
	Config.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	Config.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	Config.POSTGRES_DB = os.Getenv("POSTGRES_DB")

	lis, err := net.Listen("tcp", ":"+Config.PORT)
	if err != nil {
		log.Fatalf("[MAIN] Failed to listen: %v.", err)
	}

	cc := cryptocurrency.NewCryptoCurrency(l)
	grpcServer := grpc.NewServer()

	ccpb.RegisterCryptoCurrencyServer(grpcServer, cc)
	reflection.Register(grpcServer)

	l.Printf("[MAIN] Server online in port %s.", Config.PORT)

	if err := grpcServer.Serve(lis); err != nil {
		l.Fatalf("[MAIN] Failed to serve: %s.", err)
	}

}
