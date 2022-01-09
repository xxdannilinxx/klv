package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/lib/pq"
	"github.com/xxdannilinxx/klv/cryptocurrency"
	ccpb "github.com/xxdannilinxx/klv/proto/gen/ccpb"
	"github.com/xxdannilinxx/klv/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	PORT              string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	POSTGRES_HOST     string
	POSTGRES_PORT     string
}

var Config config

func main() {
	l := log.New(os.Stdout, "klv-api - ", log.LstdFlags)

	Config.PORT = os.Getenv("PORT")
	Config.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	Config.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	Config.POSTGRES_DB = os.Getenv("POSTGRES_DB")
	Config.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	Config.POSTGRES_PORT = os.Getenv("POSTGRES_PORT")

	psglconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Config.POSTGRES_HOST, Config.POSTGRES_PORT, Config.POSTGRES_USER, Config.POSTGRES_PASSWORD, Config.POSTGRES_DB)
	db, err := sql.Open("postgres", psglconn)
	utils.CheckError(err)

	defer db.Close()

	err = db.Ping()
	utils.CheckError(err)

	listener, err := net.Listen("tcp", ":"+Config.PORT)
	utils.CheckError(err)

	cc := cryptocurrency.NewCryptoCurrency(l, db)
	grpcServer := grpc.NewServer()

	ccpb.RegisterCryptoCurrencyServer(grpcServer, cc)
	reflection.Register(grpcServer)

	l.Printf("[MAIN] Server online in port %s.", Config.PORT)

	err = grpcServer.Serve(listener)
	utils.CheckError(err)
}
