// Pacote principal
package main

import (
	"fmt"
	"net/http"
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

// funcao main
func main() {
	http.HandleFunc("/hello", HelloHandle)
	http.HandleFunc("/cryptos", CryptoHandler)
	http.ListenAndServe(":8080", nil)
	print("Servidor online")
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
