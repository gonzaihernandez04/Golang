package main

import (
	"fmt"
	"strings"

	"github.com/gonzajs04/Golang/tree/master/GUIA0/estructuras_interfaces/figuras"
)

func main() {
	cantidadCreadas := 0
	elegida := ""
	var figurasArray []figuras.Figura = make([]figuras.Figura, 5)

	for cantidadCreadas < 5 {
		fmt.Println("Elija que figura desea crear a-Rectangulo | b - Cuadrado | c- Circulo ")
		_, err := fmt.Scanln(&elegida)
		elegidaLower := strings.ToLower(elegida)

		switch {
		case elegidaLower == "a":
			p := figuras.NewPunto(30, 30)
			p2 := figuras.NewPunto(50, 25)
			figurasArray[cantidadCreadas] = figuras.NewRectangulo(p, p2)

			cantidadCreadas++

		case elegidaLower == "b":
			p := figuras.NewPunto(30, 30)
			figurasArray[cantidadCreadas] = figuras.NewCuadrado(p, 25)
			cantidadCreadas++

		case elegidaLower == "c":
			fmt.Println("Ingrese las coordenadas del centro del círculo (x y): ")
			var x, y, rdio int
			fmt.Scanln(&x, &y)
			fmt.Println("Ingrese el radio del círculo: ")
			fmt.Scanln(&rdio)
			p := figuras.NewPunto(x, y)
			figurasArray[cantidadCreadas] = figuras.NewCirculo(rdio, p)
			cantidadCreadas++

		default:
			fmt.Println("No existe dicha opcion", err)

		}

	}

	fmt.Println("Figuras almacenadas:")
	for i, figura := range figurasArray {
		fmt.Printf("Figura %d: %s\n", i+1, figura.String())
	}

}
