package main

import "fmt"

func main() {
	var arreglo []int = []int{3, 5, 1, 2, 10}
	fmt.Printf("La suma del arreglo es %v ", SumaArreglo(arreglo))
}

func SumaArreglo(arreglo []int) int {
	if len(arreglo) <= 0 {
		return 0
	}

	var sum int = 0

	for _, numero := range arreglo {
		sum += numero
	}

	return sum

}
