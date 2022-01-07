package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/xxdannilinxx/klv/currency"
)

type config struct {
	PORT         string
	DATABASE_URL string
}

var Config config

func main() {
	Config.PORT = os.Getenv("PORT")
	Config.DATABASE_URL = os.Getenv("DATABASE_URL")

	l := log.New(os.Stdout, "klv-api - ", log.LstdFlags)

	l.Println("[MAIN] Server online in port " + Config.PORT + ".")

	ch := currency.Handlers(l)
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{id:[0-9]+}", ch.GetCurrency)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ch.AddCurrency)
	postRouter.Use(ch.MiddlewareValidationCurrency)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ch.UpdateCurrency)
	putRouter.Use(ch.MiddlewareValidationCurrency)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", ch.DeleteCurrency)

	log.Fatal(http.ListenAndServe(":"+Config.PORT, sm))
}
