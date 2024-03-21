package opciones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElegirOpciones(t *testing.T) {
	resultadoEsperado := 1
	assert.Equal(t, ElegirOpciones(1), resultadoEsperado)
}
