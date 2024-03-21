package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnion(t *testing.T) {
	var union []int = []int{1, 3, 4, 4, 5, 5, 6, 7}
	var interseccion []int = []int{4, 5}
	var vector1 []int = []int{3, 5, 4, 1}
	var vector2 []int = []int{6, 5, 4, 7}

	res, res2 := UnionInterseccion(vector1, vector2)

	assert.Equal(t, res, union)
	assert.Equal(t, res2, interseccion)

}
