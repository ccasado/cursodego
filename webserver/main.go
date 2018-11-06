package main

import (
	"fmt"
	"net/http"

	"github.com/ccasado/cursodego/webserver/manipulador"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá!")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	fmt.Println("Estou escutando ...")
	http.ListenAndServe(":8181", nil)
}
