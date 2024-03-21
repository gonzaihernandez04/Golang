package main

func main() {
	var v1 int = 2
	var v2 int = 5
	Swap(&v1, &v2)

}

// Recibe los valores con * de la direccion de memoria
func Swap(px, py *int) (int, int) {
	// Accedo al valor de la direccion de memoria de px con *
	temp := *px
	// Accedo al valor de la direccion de memoria de px y de py con * e intercambio sus valores
	*px = *py
	// Accedo al valor de la direccion de memoria de pyy le asigno el valor de temp
	*py = temp

	return *px, *py

}
