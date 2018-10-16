package main

import (
	"fmt"
)

func main() {

	var nums []int
	fmt.Println(nums, len(nums), cap(nums))
	nums = make([]int, 5)
	fmt.Println(nums, len(nums), cap(nums))
	capitais := []string{"Lisboa"}
	fmt.Println(capitais, len(capitais), cap(capitais))
	capitais = append(capitais, "Brasilia")
	fmt.Println(capitais, len(capitais), cap(capitais))
	cidades := make([]string, 4)
	cidades[0] = "NY"
	cidades[1] = "LO"
	cidades[2] = "TO"
	cidades[3] = "SG"
	fmt.Println(cidades, len(cidades), cap(cidades))
	cidades[0] = "BR"
	fmt.Println(cidades, len(cidades), cap(cidades))
	for indice, cidade := range cidades {
		fmt.Printf("Cidade[%d] = %s\r\n", indice, cidade)
	}

}
