package main

/*%d: Este es un marcador de posición específico para enteros. Se usa para imprimir valores enteros en el formato decimal. Por ejemplo, si tienes un entero x, %d imprimirá su valor en decimal.

%v: Este es un marcador de posición más general que se utiliza para imprimir el valor de cualquier tipo de dato de manera predeterminada. %v intenta representar el valor de una manera adecuada para el tipo de dato que estás imprimiendo. Por ejemplo, para un entero, %v imprimirá el valor en su forma decimal. Para una cadena, imprimirá la cadena. Es muy útil cuando no estás seguro del tipo exacto de dato que estás imprimiendo o cuando quieres imprimir valores de diferentes tipos en la misma cadena de formato.
*/
import (
	"fmt"
	"reflect"
	"unsafe"
)

// bool: Devuelve true o false
var boolean = false

// string
var cadena = "Hola mundo!"

//int o entero ocupa de 32 o 64 bits ALMACENA rangos de numero DESDE -2^31 hasta 2^30

var numero int = 2

// int 8 bits almacena numeros del -128 a 127. Ahorra memoria
var numero3 int8 = 3

// int 16 de -32768 a 32767

var numero2 int16 = 120

//int32  desde -9223372036854775808 a 9223372036854775807

var numeroLong int32 = 10000000

//uint: es similar a int, 32 a 64

var numeroUint uint = 3

//Uint8 almacena de 0 a 255

var numeroUint8 uint8 = 254

/*   uint16     // 0 a 65535
uint32     // 0 a 4294967295
uint64     // 0 a 18446744073709551615
float32    // 32 bits
float64    // 64 bits
complex64  // 32 bits
complex128 // 64 bits
byte       // alias para uint8
rune       // alias para int32, representa una posición en código Unicode
uintptr    // utilizado para guardar una dirección de puntero

*/

func main() {
	fmt.Printf("El tamaño en Bytes del numero %v es de %v\n", numeroLong, reflect.TypeOf(numeroLong).Size())
	fmt.Printf("Tamaño en bytes de la variable numeroLong: %d\n", unsafe.Sizeof(numeroLong))

}
