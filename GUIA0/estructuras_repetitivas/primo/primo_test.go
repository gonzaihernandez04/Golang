package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEsPrimo(t *testing.T) {
	resultadoEsperado := "El numero si es primo"

	assert.Equal(t, EsPrimo(7), resultadoEsperado)
}

func TestNoEsPrimo(t *testing.T) {
	resultadoEsperado := "El numero no es primo"

	assert.Equal(t, EsPrimo(6), resultadoEsperado)
}
