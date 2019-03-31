package operadora

import (
	"strconv"
	"github.com/marceloagmelo/cursodego/pacotes/prefixo"
)

// Nome Operadora representa o nome da operadora
var NomeOperadora = "XPTO Telecom"

// PrefixoDaCapitalOperadora
var PrefixoDaCapitalOperadora =  strconv.Itoa(prefixo.Capital) + " " + NomeOperadora