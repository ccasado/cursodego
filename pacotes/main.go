package main

import (
	"fmt"

	"github.com/ccasado/cursodego/pacotes/operadora"
	"github.com/ccasado/cursodego/pacotes/prefixo"
)

//Nome do usuario do sistema
var NomeDoUsuario = "Cris"

func main() {
	fmt.Printf("Nome do usuário: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeDaOperadora)
	fmt.Printf("Teste com prefixo: %s\r\n", prefixo.TesteComPrefixo)

}
