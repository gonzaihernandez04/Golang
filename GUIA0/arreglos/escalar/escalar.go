package main

import (
	"errors"
	"fmt"
)

func main() {
	var vector1 []int = []int{3, 2, 1}
	var vector2 []int = []int{6, 5, 4}
	Escalar(vector1, vector2)
}

func Escalar(vector1 []int, vector2 []int) (int, error) {
	if len(vector1) != len(vector2) {
		fmt.Println("Los vectores deben ser del mismo tamaño")
		return 0, errors.New("los vectores deben ser del mismo tamaño")
	}
	var productoEscalar int = 0

	for i := 0; i < len(vector1); i++ {
		productoEscalar += (vector1[i] * vector2[i])
	}

	fmt.Printf("La suma de producto escalar de los vectores es de: %v: ", productoEscalar)
	return productoEscalar, nil
}
