package main

import "fmt"

func ImprimirPolinomio(coeficientes []float64) string {
	resultado := ""
	grado := len(coeficientes) - 1
	signo := ""

	for i := len(coeficientes) - 1; i >= 0; i-- {
		coeficiente := coeficientes[i]
		var strCadena string
		if coeficiente > 0 && i != len(coeficientes)-1 {
			signo = "+"
		} else if coeficiente < 0 {
			signo = "-"
			coeficiente = -coeficiente
		}
		if i < grado {
			if grado-i == 1 {
				strCadena = "x"
			} else {
				strCadena = fmt.Sprintf("x^%d", grado-i)

			}

		}
		resultado += fmt.Sprintf(" %s %.1f%s", signo, coeficiente, strCadena)
		if resultado == "" {
			resultado = "VACIO"
		}

	}

	return resultado

}
