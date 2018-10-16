package operadora

import (
	"strconv"

	"github.com/ccasado/cursodego/pacotes/prefixo"
)

//NomeDaOperadora Nome da operadora
var NomeDaOperadora = "XPTO Telecom"

//PrefixoDaCapitalOperadora prefixo mais o nome da operadora
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeDaOperadora
