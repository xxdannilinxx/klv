package main

import (
	"log"
	"net/http"

	"github.com/alexflint/go-arg"
	"github.com/xxdannilinxx/klv/currency"
)

type config struct {
	PORT         string `arg:"env:PORT, -P, --PORT" help:"Port of the application" placeholder:"PORT"`
	DATABASE_URL string `arg:"env:DATABASE_URL, -D, --DATABASE_URL" help:"Connection with database" placeholder:"DATABASE_URL"`
}

var Config config

func main() {
	arg.MustParse(&Config)

	log.Println("[MAIN] Server online in port " + Config.PORT + ".")

	mux := http.NewServeMux()
	mux.Handle("/currency", &currency.CurrencysHandler{})

	log.Fatal(http.ListenAndServe(":"+Config.PORT, mux))
}
