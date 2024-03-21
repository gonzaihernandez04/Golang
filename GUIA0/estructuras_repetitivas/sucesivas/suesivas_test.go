package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSucesivas(t *testing.T) {
	n := 1
	n2 := 6
	assert.Equal(t, Sucesivas(n, n2), 6)
}
