package model

//Polimorfismo

//Galinha representa uma ave do tipo Galinha
type Galinha interface {
	Cacareja() string
}

//Pata representa uma ave do tipo Pata
type Pata interface {
	Quack() string
}

//Ave representa um animal
type Ave struct {
	Nome string
}

//Cacareja retorna o som emitido por uma galinha
func (a Ave) Cacareja() string {
	return "Cócóricó..."
}

//Quack retorna o som emitido por uma pata
func (a Ave) Quack() string {
	return "Quack..."
}
