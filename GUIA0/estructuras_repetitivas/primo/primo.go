package main

import "fmt"

func main() {

	fmt.Println(EsPrimo(14))

}

func EsPrimo(n int) string {
	var result string
	if n <= 1 {
		result = "El numero no puede ser menor o igual a 1"
		return result
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			result = "El numero no es primo"
			return result
		}
	}
	result = "El numero si es primo"
	return result
}
