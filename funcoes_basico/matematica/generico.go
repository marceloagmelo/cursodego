package matematica

/*
Calculo que executa qualquer tipo de calculo, basta enviar a função desejada
*/
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

// Multiplicador multiplica número x * y
func Multiplicador(x int, y int) int {
	return x * y
}

// Divisor efetua a divisão
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

// DivisorComResto retorna o resultado e o resto da divisão
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA % numeroB
	return
}