package main

import "fmt"

func main() {
	n := 1
	n2 := 6
	fmt.Printf("El producto de n y n2 es: %v ", Sucesivas(n, n2))
}

func Sucesivas(n int, n2 int) int {
	var resultado int = 0
	if n2 < 0 {
		// Cambio el valor por si es negativo
		n2 = -n2
	}
	for i := 0; i < n2; i++ {
		resultado += n
	}

	return resultado
}
