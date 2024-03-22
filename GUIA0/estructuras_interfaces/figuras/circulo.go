package figuras

import (
	"fmt"
	"math"
)

type Circulo struct {
	centro Punto
	rdio   int
}

func NewCirculo(rdio int, centro Punto) Circulo {
	return Circulo{centro: centro, rdio: rdio}
}

func (c Circulo) Area() int {
	return int((math.Pi) * (float64(c.rdio) * float64(c.rdio)))
}

func (c Circulo) Diametro() int {
	return c.rdio * 2
}

func (c Circulo) Perimetro() int {
	return int(2 * math.Pi * float64(c.rdio))
}

func (c Circulo) String() string {
	return fmt.Sprintf("Círculo de radio %d", c.rdio)
}
