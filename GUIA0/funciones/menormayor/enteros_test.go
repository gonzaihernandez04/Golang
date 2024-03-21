package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMaximoEnteros(t *testing.T) {
	resultadoEsperado1 := 1
	resultadoEsperado2 := 9

	var enteros []int = []int{5, 9, 1, 3, 2, 8}
	min, max := MinMaximoEnteros(enteros)
	assert.Equal(t, resultadoEsperado1, min)
	assert.Equal(t, resultadoEsperado2, max)
}
