package declaraciones

import (
	"fmt"
	"math/cmplx"

	// Permite ver tipos de datos
	"reflect"
	"time"
)

// Declaraciones globales
// Declaraciones globales
var numero = 2
var numero2, numero3, numero4 = 8, 5, 7
var nombre, apellido, edad = "Gonzalo", "Hernandez", 20

// Declaracion y especificacion de tipo de dato
var entero1, entero2, entero3 int = 2, 2, 2

// Declaracion multiple
var (
	esNum                = false
	MaxUInt64 uint64     = 1<<64 - 1
	complejoZ complex128 = cmplx.Sqrt(-5 + 12i)
)

func MostrarVariables() {

	// %v, Que signifca? Coloca el valor con el tipo de dato correspondiente al declarado
	fmt.Printf("Los valores de los numeros son %v %v %v ", entero1, entero2, entero3)

	currentTime := time.Now()

	fmt.Println("El tipo de `time.Now()` es:", reflect.TypeOf(currentTime))
	fmt.Println("El tipo de `complejoZ` es:", reflect.TypeOf(complejoZ))

	fmt.Println("El tipo de `numero` es:", reflect.TypeOf(nombre))

}
