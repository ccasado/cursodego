package main

import (
	"fmt"

	"github.com/ccasado/cursodego/matematica/operacoes"
)

func main() {
	resultado := operacoes.Calculo(operacoes.Multiplicador, 2, 3)
	fmt.Printf("O resultado da multiplicação foi %d\r\n", resultado)
	resultado = operacoes.Calculo(operacoes.Soma, 2, 3)
	fmt.Printf("O resultado da soma foi %d\r\n", resultado)
	resultado = operacoes.Calculo(operacoes.Divisor, 2, 2)
	fmt.Printf("O resultado da soma foi %d\r\n", resultado)
	resultado, resto := operacoes.DivisorComResto(2, 2)
	fmt.Printf("O resultado da divisão foi %d\r\n", resultado)
	fmt.Printf("O resultado do resto da divisão foi %d\r\n", resto)
}
