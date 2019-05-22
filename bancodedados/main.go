package main

import (
	"fmt"
	"net/http"

	"github.com/ccasado/cursodego/bancodedados/manipulador"
	"github.com/ccasado/cursodego/bancodedados/repo"
)

func init() {
	fmt.Println("Vamos começar a subir o server...")
}

func main() {

	err := repo.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println("Erro ao abrir o banco de dados: ", err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá!")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local/", manipulador.Local)
	fmt.Println("Estou escutando ...")
	http.ListenAndServe(":8181", nil)
}
