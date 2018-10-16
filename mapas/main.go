package main

import (
	"fmt"

	"github.com/ccasado/cursodego/mapas/model"
)

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa amarela"
	casa.X = 18
	casa.Y = 54
	casa.SetValor(600000)

	cache["Casa amarela"] = casa

	fmt.Println("Há uma casa amarela no cache?")
	Imovel, achou := cache["Casa amarela"]
	if achou {
		fmt.Printf("Sim, e o que achei foi: %+v\r\n", Imovel)
	}

	apto := model.Imovel{}
	apto.Nome = "Apto azul"
	apto.X = 18
	apto.Y = 54
	apto.SetValor(700000)

	cache[apto.Nome] = apto

	fmt.Println("Quanto itens ha no cache? ", len(cache))

	for chave, Imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\n\r", chave, Imovel)
	}

	Imovel, achou = cache["Casa amarela"]
	if achou {
		delete(cache, Imovel.Nome)
	}

	fmt.Println("Quanto itens ha no cache? ", len(cache))

}
