package main

func main() {
	var enteros []int = []int{5, 9, 1, 3, 2, 8}
	MinMaximoEnteros(enteros)

}

func MinMaximoEnteros(enteros []int) (int, int) {
	if len(enteros) > 0 {
		var min int = enteros[0]
		var max int
		for _, numero := range enteros {
			if numero < min {
				min = numero
			}
			if numero > max {
				max = numero
			}
		}

		return min, max
	}
	return 0, 0
}
