package main

import (
	"github.com/marceloagmelo/cursodego/funcoes_basico/matematica"
	"fmt"
)

func main() {
	//resultado := multiplicador(2,2)
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("O resultado do multiplicador foi: %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Divisor, 12, 3)
	fmt.Printf("O resultado da divisão foi: %d\r\n", resultado)
	resultado, resto := matematica.DivisorComResto(12, 5)
	fmt.Printf("O resultado da divisão foi: %d e o resto foi %d\r\n", resultado, resto)
}

func multiplicador(x int, y int) int {
	return x * y
}