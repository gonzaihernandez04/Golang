package figuras

type Circulo struct {
	centro Punto
	rdio   float32
}

func NewCirculo(rdio float32, centro Punto) Circulo {
	return Circulo{centro: centro, rdio: rdio}
}
