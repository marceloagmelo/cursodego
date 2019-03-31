package main

import (
	"github.com/marceloagmelo/cursodego/pacotes/operadora"
	"fmt"
	"github.com/marceloagmelo/cursodego/pacotes/prefixo"
)

// NomedDousuario é um nome de usuário de sistema
var NomeDoUsuario = "Marcelo"

func main() {
	fmt.Printf("Nome do usuário: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da Operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Valor de teste: %s\r\n", prefixo.TesteComPrefixo)
}