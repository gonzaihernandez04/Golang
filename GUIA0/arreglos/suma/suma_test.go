package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumaArreglo(t *testing.T) {
	var arreglo []int = []int{3, 5, 1, 2, 10}
	suma := 21

	assert.Equal(t, SumaArreglo(arreglo), suma)
}
