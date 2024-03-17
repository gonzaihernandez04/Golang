package structs

type Figura interface {
	GetArea() float64
	Perimetro() float64 // Cambiado a float64
	String() string     // Cambiado a string
}
