package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwap(t *testing.T) {
	v1 := 2
	v2 := 5

	r1, r2 := Swap(&v1, &v2)
	assert.Equal(t, r1, v1)
	assert.Equal(t, r2, v2)

}
