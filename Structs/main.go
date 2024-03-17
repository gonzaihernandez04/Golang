package main

import (
	"fmt"

	"github.com/gonzajs04/Golang/tree/master/Structs/structs"
)

func main() {
	fmt.Println("Hola mundo")

	// Reserva memoria para todos los campos y las setea en 0
	var c *structs.Circle = new(structs.Circle)
	c.Radio = 10
	c.Color = "Verde"

	fmt.Println("Radio:", c.Radio)
	fmt.Println("Área:", c.Color)

	// Segunda forma
	var circulo structs.Circle = structs.Circle{Radio: 10, Color: "Verde"}

	fmt.Println("Radio:", circulo.Radio)
	fmt.Println("Área:", circulo.Color)

	// Tercera forma

	var circulo2 = structs.NewCircle(30, "Azul")
	fmt.Println("Área:", circulo2.Color)
	fmt.Println("Radio:", circulo2.Radio)

	// METHODS

	fmt.Println(circulo2.GetArea())

	cuadrado := structs.NewCuadrado(5)

	fmt.Println(cuadrado.GetArea())

	array_figuras := [2]structs.Figura{cuadrado, circulo2}

	for _, f := range array_figuras {
		fmt.Println("Área de la figura:", f.GetArea())

	}

}
