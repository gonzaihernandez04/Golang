package figuras

type Cuadrado struct {
	posicion Punto
	lado     int
}

func NewCuadrado(posicion Punto, lado int) Cuadrado {
	return Cuadrado{posicion: posicion, lado: lado}
}
