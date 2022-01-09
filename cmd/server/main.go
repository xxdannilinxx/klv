package main

import (
	"log"
	"net"
	"os"

	_ "github.com/lib/pq"
	"github.com/xxdannilinxx/klv/cryptocurrency"
	"github.com/xxdannilinxx/klv/db"
	ccpb "github.com/xxdannilinxx/klv/proto/gen/ccpb"
	"github.com/xxdannilinxx/klv/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	Config utils.Config
)

// Role responsible for uploading the server and connecting to the database
func main() {
	l := log.New(os.Stdout, "klv-api - ", log.LstdFlags)

	Config.PORT = os.Getenv("PORT")
	Config.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	Config.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	Config.POSTGRES_DB = os.Getenv("POSTGRES_DB")
	Config.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	Config.POSTGRES_PORT = os.Getenv("POSTGRES_PORT")

	db := db.ConnectDB(Config)

	listener, err := net.Listen("tcp", ":"+Config.PORT)
	utils.CheckError(err)

	ccRepository := cryptocurrency.NewCryptoCurrencyRepository(db)
	cc := cryptocurrency.NewCryptoCurrency(l, ccRepository)
	grpcServer := grpc.NewServer()

	ccpb.RegisterCryptoCurrencyServer(grpcServer, cc)
	reflection.Register(grpcServer)

	l.Printf("[MAIN] Server online in port %s.", Config.PORT)

	err = grpcServer.Serve(listener)
	utils.CheckError(err)
}
