package main

import (
	"log"
	"net"
	"os"

	"github.com/xxdannilinxx/klv/cryptocurrency"
	"github.com/xxdannilinxx/klv/pgsql"
	ccpb "github.com/xxdannilinxx/klv/proto/gen/ccpb"
	"github.com/xxdannilinxx/klv/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	Config utils.Config = utils.Config{
		PORT:              os.Getenv("PORT"),
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		POSTGRES_PORT:     os.Getenv("POSTGRES_PORT"),
	}
)

// Role responsible for uploading the server and connecting to the database
func main() {
	l := log.New(os.Stdout, "klv-api - ", log.LstdFlags)

	dbConn := pgsql.ConnectDB(Config)

	listener, err := net.Listen("tcp", ":"+Config.PORT)
	utils.CheckError(err)

	ccRepository := cryptocurrency.NewCryptoCurrencyRepository(dbConn)
	cc := cryptocurrency.NewCryptoCurrencyService(l, ccRepository)
	grpcServer := grpc.NewServer()

	ccpb.RegisterCryptoCurrencyServer(grpcServer, cc)
	reflection.Register(grpcServer)

	l.Printf("[MAIN] Server online in port %s.", Config.PORT)

	err = grpcServer.Serve(listener)
	utils.CheckError(err)
}
