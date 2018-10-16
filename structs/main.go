package main

import (
	"fmt"
)

//Imovel é uma struct que armazena dados de um imovel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {

	casa := Imovel{}
	fmt.Printf("a casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu ap", 750000}
	fmt.Printf("o apartamento é: %+v\r\n", apartamento)

	chacara := Imovel{
		Y:     85,
		X:     26,
		Nome:  "Chacara",
		valor: 150000,
	}
	fmt.Printf("a chacara é: %+v\r\n", chacara)

	casa.X = 18
	casa.Y = 23
	casa.Nome = "Lar doce lar"
	casa.valor = 100000
	fmt.Printf("a casa é: %+v\r\n", casa)

}
