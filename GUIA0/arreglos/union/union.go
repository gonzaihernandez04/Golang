package main

import "fmt"

func main() {
	var vector1 []int = []int{3, 5, 4, 1}
	var vector2 []int = []int{6, 5, 4, 7}

	UnionInterseccion(vector1, vector2)
	//Union2(vector1, vector2)
}

func UnionInterseccion(vector1 []int, vector2 []int) ([]int, []int) {

	var union []int = make([]int, (len(vector1) + len(vector2)))
	var interseccion []int = make([]int, 0)
	var ultimaPosicion int = 0

	// Recorrer primer array
	for i := 0; i < len(vector1); i++ {
		union[i] = vector1[i]
		ultimaPosicion = i
	}
	ultimaPosicion++

	// Recorrer segundo array
	for i := 0; i < len(vector2); i++ {
		union[ultimaPosicion] = vector2[i]
		ultimaPosicion++
	}

	// Ordeno array para buscar valores que se repitan
	Ordenar(union)

	// Interseccion entre 2 arrays
	for i := 0; i < len(union)-1; i++ {
		if union[i] == union[i+1] {

			// Agrego a interseccion el valor repetido
			interseccion = append(interseccion, union[i])
		}
	}

	fmt.Println(union)
	fmt.Println(interseccion)

	return union, interseccion

}

// Ordenamiento por seleccion
func Ordenar(union []int) {

	// Nuestro contador de pasos
	stepCounter := 1

	// Iteramos de la primera a la penúltima posición del slice
	for i := 0; i < len(union)-1; i++ {
		menor := i
		// iteramos los elementos restantes del slice, buscándo el número menor
		for j := i + 1; j < len(union); j++ {

			stepCounter++

			if union[menor] > union[j] {
				menor = j
			}
		}

		v := union[i]
		union[i] = union[menor]
		union[menor] = v

	}

}

func Union2(vector1 []int, vector2 []int) []int {

	// ... agarra los elementos individuales del 2 vector y lo agrega a vector1
	union := append(vector1, vector2...)

	return union

}
