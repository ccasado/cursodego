package main

import (
	"encoding/json"
	"fmt"

	"github.com/ccasado/cursodego/erro/model"
)

func main() {

	casa := model.Imovel{}
	casa.Nome = "Casa amarela"
	casa.X = 18
	casa.Y = 25
	if err := casa.SetValor(100000000); err != nil {
		fmt.Println("[main] Houve um erro na atribuicao do valor: ", err)
		if err == model.ErrValorDeImovelMuitoAlto {
			fmt.Println("[main] Corretor, verifique sua avaliação")
		}
		return
	}

	fmt.Printf("Casa é: %+v\r\n", casa)

	objJSON, err := json.Marshal(casa)
	if err != nil {
		fmt.Println("[main] Houve um erro na geração do objeto JSON: ", err.Error)
		return
	}

	fmt.Println("A casa em JSON: ", string(objJSON))

}
