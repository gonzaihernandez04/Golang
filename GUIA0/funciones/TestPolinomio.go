package funciones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImprimirPolinomio(t *testing.T) {
	var coeficientes []float64 = []float64{-3.0, -2.0, 5.0, 6.0}
	resultadoEsperado := "6.0 + 5.0x - 2.0x^2 - 3.0x^3"
	resultado := ImprimirPolinomio(coeficientes)
	assert.Equal(t, resultadoEsperado, resultado)

}
