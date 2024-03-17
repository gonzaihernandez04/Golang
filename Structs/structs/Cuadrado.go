package structs

import "fmt"

type Cuadrado struct {
	Lado float64
}

func NewCuadrado(ladoC float64) Cuadrado {
	return Cuadrado{Lado: ladoC}
}

func (c Cuadrado) GetArea() float64 {
	return c.Lado * c.Lado
}

func (c Cuadrado) Perimetro() float64 {
	return 4 * c.Lado
}

func (c Cuadrado) String() string {
	return fmt.Sprintf("Cuadrado de lado %.2f", c.Lado)
}
