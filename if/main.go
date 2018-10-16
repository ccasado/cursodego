package main

import (
	"fmt"
)

func main() {

	meses := 0
	situacao := true
	cidade := "rio de janeiro"

	//< > != == <= >= && ||
	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo")
	}

	if situacao {
		fmt.Println("Ele esta devendo")
	}

	if !situacao {
		fmt.Println("Ele esta adimplente")
	}

	if cidade == "rio de janeiro" {
		fmt.Println("O cliente vive no estado RJ")
	}

	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println("Qual a situacao do cliente? ", descricao)
		return
	}

	fmt.Println("Obrigado pela sua consulta")

}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente esta devendo"
		return
	}
	descricao = "O cliente esta com o carne em dia"
	return
}
