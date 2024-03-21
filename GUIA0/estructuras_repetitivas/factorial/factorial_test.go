package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	n := 6
	resultadoEsperado := 720

	factorial, err := Factorial(n)
	assert.Equal(t, factorial, resultadoEsperado)
	assert.Equal(t, err, nil)
}
