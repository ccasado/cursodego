package model

import "errors"

//Imovel estrutura de dados do imovel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

/*
ErrValorDeImovelInvalido é um erro que é apresentado quando é atribuido ao imovel
um valor muito baixo
*/
var ErrValorDeImovelInvalido = errors.New("O valor informado não é válido")

/*
ErrValorDeImovelMuitoAlto é um erro quando é atribuido ao imovel
um valor muito alto
*/
var ErrValorDeImovelMuitoAlto = errors.New("O valor informado é muito alto")

//SetValor define o valor do imovel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 10000000 {
		err = ErrValorDeImovelMuitoAlto
		return
	}
	i.valor = valor
	return
}

//GetValor retorna o valor do imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
