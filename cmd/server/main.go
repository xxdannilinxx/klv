// Pacote principal
package main

import (
	"fmt"
	"net/http"

	"github.com/alexflint/go-arg"
)

// Struct
// maiusculo = publico, acessa de outros arquivos
// minusculo = privado, só acessa no mesmo arquivo
type Crypto struct {
	Id       int
	Name     string
	Likes    int
	Deslikes int
}

func (c Crypto) Create() {

}

type config struct {
	PORT         string `arg:"env:PORT, -P, --PORT" help:"Port of the application" placeholder:"PORT"`
	DATABASE_URL string `arg:"env:DATABASE_URL, -D, --DATABASE_URL" help:"Connection with database" placeholder:"DATABASE_URL"`
}

var Config config

// funcao main
func main() {
	arg.MustParse(&Config)

	println("Servidor online na porta", Config.PORT)
	println("Configuração de conexão com o postgre", Config.DATABASE_URL)

	http.HandleFunc("/", HomeHandle)
	http.HandleFunc("/hello", HelloHandle)
	http.HandleFunc("/cryptos", CryptoHandler)

	if err := http.ListenAndServe(":"+Config.PORT, nil); err != nil {
		println("ListenAndServe: ", err)
	}
}

// teste
func HomeHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Início")
}

// teste
func HelloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Teste")
}

// teste
func CryptoHandler(w http.ResponseWriter, r *http.Request) {

	// crypto1 := Crypto{
	// 	Id:       1,
	// 	Name:     "Bitcoin",
	// 	Likes:    8,
	// 	Deslikes: 0,
	// }

	// crypto := []Crypto{
	// 	{1, "Bitcoin", 8, 0},
	// 	{2, "Dogcoin", 3, 5},
	// 	{3, "Klever", 10, 0},
	// 	{4, "Ethereum", 7, 1},
	// }
	// print(crypto)

	// x := Crypto{}
	// x.Create()

	fmt.Fprintf(w, "Ok")
}
