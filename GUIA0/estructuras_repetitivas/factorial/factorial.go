package main

import (
	"errors"
	"fmt"
)

func main() {
	n := 6
	Factorial(n)
}

func Factorial(n int) (int, error) {
	var factorial int = 1
	if n < 0 {
		return 0, errors.New("el numero no puede ser negativo")
	}

	if n > 0 {
		for n > 0 {
			factorial = factorial * (n)
			n--
		}
	}

	fmt.Println(factorial)
	return factorial, nil
}
