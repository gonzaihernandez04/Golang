package main

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/untref-ayp2/mymodules/math"
	"github.com/untref-ayp2/mymodules/saludos"
)

func main() {
	msg := color.RedString("Saludos desde argentina")
	fmt.Println(msg)
	a := 10
	b := 5
	c := math.Sumar(a, b)

	saludo := saludos.Saludando()

	constante := math.PI

	fmt.Printf("La suma de %d y %d es de %d\n", a, b, c)

	fmt.Println(saludo)

	fmt.Println(constante)

}
