package main

import (
	"fmt"
)

func main() {
	meses := 6
	situacao := true
	cidade := "Teste"

	// < > != == <= >= && ||
	if meses <= 6 {
		fmt.Println("Esse credor deve a pouco tempo")
	}

	if situacao {
		fmt.Println("Ele esta devendo")
	}

	if !situacao {
		fmt.Println("Ele esta adimplente")
	}

	if cidade == "Teste" {
		fmt.Println("O cliente vive no estado Teste")
	}

	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println("Qual a situação do cliente? ", descricao)
		return
	}
	// A opção abaixo não compila
	//fmt.Println("Me passa o status", descricao)
	fmt.Println("Obrigado por nos consultar")
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