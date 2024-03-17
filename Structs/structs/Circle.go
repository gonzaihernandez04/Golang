package structs

import (
	"fmt"
	"math"
)

type Circle struct {
	Radio float64
	Color string
}

func NewCircle(radio float64, color string) Circle {
	return Circle{Radio: radio, Color: color}
}

// GetArea es un método de la estructura Circle
func (c Circle) GetArea() float64 {
	return math.Pi * (c.Radio * c.Radio)
}

func (c Circle) Perimetro() float64 {
	return 2 * math.Pi * c.Radio
}

func (c Circle) String() string {
	return fmt.Sprintf("Círculo de radio %.2f", c.Radio)
}
