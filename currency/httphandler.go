package currency

import (
	"log"
	"net/http"
)

type CurrencysHandler struct{}

func (c *CurrencysHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.getCurrency(w, r)
	case http.MethodPost:
		c.addCurrency(w, r)
	case http.MethodPut:
		c.updateCurrency(w, r)
	case http.MethodDelete:
		c.deleteCurrency(w, r)
	default:
		log.Println("[CURRENCY] " + r.Method + " - method not allowed")
		http.Error(w, "404 page not found", http.StatusNotFound)
	}
}

func (c *CurrencysHandler) getCurrency(w http.ResponseWriter, r *http.Request) {
	log.Println("[CURRENCY] GET - currency by id")
	w.Write([]byte("get currency"))
}

func (c *CurrencysHandler) addCurrency(w http.ResponseWriter, r *http.Request) {
	log.Println("[CURRENCY] POST - add currency")
	w.Write([]byte("add currency"))
}

func (c *CurrencysHandler) updateCurrency(w http.ResponseWriter, r *http.Request) {
	log.Println("[CURRENCY] PUT - update currency")
	w.Write([]byte("add currency"))
}

func (c *CurrencysHandler) deleteCurrency(w http.ResponseWriter, r *http.Request) {
	log.Println("[CURRENCY] DELETE - delete currency")
	w.Write([]byte("add currency"))
}
