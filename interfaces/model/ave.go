package model

/*
Galinha representa uma ave do tipo Galinha
*/
type Galinha interface {
	Cacareja() string
}

/*
Pata representa uma ave do tipo Pata
*/
type Pata interface {
	Quack() string
}

//Ave representa um anival
type Ave struct {
	Nome string
}

//Caracarja retorna o som de uma galinha
func (a Ave) Cacareja() string {
	return "Cócóricó"
}

//Quack retorna o som de uma pata
func (a Ave) Quack() string {
	return "Quack, quack..."
}