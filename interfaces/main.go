package main

import (
	"fmt"
	"github.com/marceloagmelo/cursodego/interfaces/model"
)

func main() {
	jojo := model.Ave{}
	jojo.Nome = "Jojo da Silva"

	queroAcordarComUmCarcarejo(jojo)
	queroOuvirUmaPataNoLago(jojo)
}

func queroAcordarComUmCarcarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func queroOuvirUmaPataNoLago(p model.Pata) {
	fmt.Println(p.Quack())
}