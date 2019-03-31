package main

import (
	"fmt"
)

// Imovel é uma structs dados do imóvel
type Imovel struct {
	X int
	Y int
	Nome string
	valor int
}

func main() {
	casa := Imovel{}
	fmt.Printf("A casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu AP", 760000}
	fmt.Printf("A aparamento é: %+v\r\n", apartamento)

	chacara := Imovel {
		Y: 85,
		Nome: "Chacara",
		X: 22,
		valor: 55,
	}
	fmt.Printf("A chacara é: %+v\r\n", chacara)

	casa.Nome = "Lar doce lar" 
	casa.valor = 450000
	casa.X = 18
	casa.Y = 31
	fmt.Printf("A casa é: %+v\r\n", casa)
}