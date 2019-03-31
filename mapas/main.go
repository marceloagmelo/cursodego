package main

import (
	"fmt"
	"github.com/marceloagmelo/cursodego/structs_avancado/model"
)

var cache map[string]model.Imovel

func main() {
	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	cache["Casa Amarela"] = casa

	fmt.Println("Há uma Casa Amarela no cache?")
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim, e o que achei foi: %+v\r\n", imovel)
	}

	apto := model.Imovel{}
	apto.Nome = "Apto Azul"
	apto.X = 19
	apto.Y = 26
	apto.SetValor(60000)

	cache[apto.Nome] = apto

	fmt.Println("Quantos intens há no cache? ", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave [%s] = %+v\r\n", chave, imovel)
	}

	_, achou = cache["Casa Amarela"]
	if achou {
		delete(cache, imovel.Nome)
	}

	fmt.Println("Quantos intens há no cache? ", len(cache))
}