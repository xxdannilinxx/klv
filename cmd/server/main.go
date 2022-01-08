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
	PORT         string
	DATABASE_URL string
}

var Config config

// testes
// fazer repository de alguma forma que de pra trocar o banco
// passar logger por referencia
// exemplos de chamada usando grpcurl
// criar dto?
func main() {
	l := log.New(os.Stdout, "klv-api - ", log.LstdFlags)

	Config.PORT = os.Getenv("PORT")
	if Config.PORT == "" {
		l.Fatal("[MAIN] Unspecified port")
	}
	Config.DATABASE_URL = os.Getenv("DATABASE_URL")
	if Config.DATABASE_URL == "" {
		l.Fatal("[MAIN] Unspecified database connection.")
	}

	lis, err := net.Listen("tcp", ":"+Config.PORT)
	if err != nil {
		log.Fatalf("[MAIN] Failed to listen: %v.", err)
	}

	cc := cryptocurrency.Server{}
	grpcServer := grpc.NewServer()

	ccpb.RegisterCryptoCurrencyServer(grpcServer, &cc)
	reflection.Register(grpcServer)

	l.Printf("[MAIN] Server online in port %s.", Config.PORT)

	if err := grpcServer.Serve(lis); err != nil {
		l.Fatalf("[MAIN] Failed to serve: %s.", err)
	}

}
