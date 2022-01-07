package currency

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CurrencyHandler struct {
	l       *log.Logger
	service *CurrencyService
}

func Handlers(l *log.Logger) *CurrencyHandler {
	return &CurrencyHandler{l, &CurrencyService{}}
}

func (c *CurrencyHandler) GetCurrency(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		c.l.Println("[CURRENCY] ERROR ", err)
		http.Error(w, "[CURRENCY] Unable to convert id", http.StatusBadRequest)
		return
	}
	c.l.Println("[CURRENCY] GET - currency by id ", id)
	w.Write([]byte("get currency"))
}

func (c *CurrencyHandler) AddCurrency(w http.ResponseWriter, r *http.Request) {
	c.l.Println("[CURRENCY] POST - add currency")
	w.Write([]byte("add currency"))
}

func (c *CurrencyHandler) UpdateCurrency(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		c.l.Println("[CURRENCY] ERROR ", err)
		http.Error(w, "[CURRENCY] Unable to convert id", http.StatusBadRequest)
		return
	}
	c.l.Println("[CURRENCY] PUT - update currency ", id)
	w.Write([]byte("add currency"))
}

// testes
// adicionar comentários
// grpc
// os serviços vao ficar em services
// criar dto
func (c *CurrencyHandler) DeleteCurrency(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		c.l.Println("[CURRENCY] ERROR ", err)
		http.Error(w, "[CURRENCY] Unable to convert id", http.StatusBadRequest)
		return
	}
	c.l.Println("[CURRENCY] DELETE - delete currency ", id)
	del := c.service.delete(1)
	w.Write([]byte(del))
}

func (c *CurrencyHandler) MiddlewareValidationCurrency(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cur := &Currency{}

		err := cur.Unmarshal(r.Body)
		if err != nil {
			c.l.Println("[CURRENCY] ERROR ", err)
			http.Error(w, "[CURRENCY] Error reading currency", http.StatusBadRequest)
			return
		}

		err = cur.Validate()
		if err != nil {
			c.l.Println("[CURRENCY] ERROR ", err)
			http.Error(w, "[CURRENCY] Error validate currency: ", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
