package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumaArreglo(t *testing.T) {
	var vector1 []int = []int{3, 2, 1}
	var vector2 []int = []int{6, 5, 4}
	result, err := Escalar(vector1, vector2)
	assert.Equal(t, result, 32)
	assert.Equal(t, err, nil)
}

func TestNoSonIguales(t *testing.T) {
	var vector1 []int = []int{3, 2, 1}
	var vector2 []int = []int{6, 5, 9, 4}
	result, err := Escalar(vector1, vector2)
	assert.Equal(t, result, 0)
	assert.Error(t, err, "los vectores deben ser del mismo tamaño")
}
