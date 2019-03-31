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
	casa := new(Imovel)
	fmt.Printf("A casa é: %+v\r\n", casa)
	fmt.Printf("A casa com endereço de memória é: %p - %+v\r\n", &casa, casa)

	chacara := Imovel{17, 28, "Chacara linda", 280000}
	fmt.Printf("A chacara é: %p - %+v\r\n", &chacara, chacara)
	mudaImovelComPonteiro(&chacara)
	fmt.Printf("A chacara com ponteiro é: %p - %+v\r\n", &chacara, chacara)
	mudaImovel(chacara)
	fmt.Printf("A chacara sem ponteiro é: %p - %+v\r\n", &chacara, chacara)
}

func mudaImovelComPonteiro(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5
}

func mudaImovel(imovel Imovel) {
	imovel.X = imovel.X + 20
	imovel.Y = imovel.Y - 1
}

