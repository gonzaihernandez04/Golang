package main

import (
	"fmt"
)

func main() {
	var result bool
	var err error
	tienda := CrearTienda(10)
	result, err = AgregarProducto(&tienda, "Pera", 500)
	result, err = AgregarProducto(&tienda, "Manzana", 500)
	result, err = AgregarProducto(&tienda, "Palta", 500)

	fmt.Println(tienda.ArrayProductos)

	fmt.Println(result)
	fmt.Println(err)

}

type error interface {
	Error() string
}
type MyError struct {
	Code    int
	Message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("Error %d: %s ", e.Code, e.Message)
}

type Tienda struct {
	ArrayProductos []Producto
	CantProductos  int
}

func CrearTienda(TotalTiendaProductos int) Tienda {
	tienda := Tienda{}
	tienda.CantProductos = 0
	tienda.ArrayProductos = make([]Producto, TotalTiendaProductos)

	return tienda
}

// Puede retornar de tipo bool y error
func AgregarProducto(t *Tienda, nombre string, precio int) (bool, error) {
	if t.CantProductos < len(t.ArrayProductos) {
		t.ArrayProductos[t.CantProductos] = NewProducto(nombre, precio)

		t.CantProductos++
		return true, nil
	} else {
		return false, MyError{Code: 1, Message: "Cantidad excedida de productos"}

	}

}

type Producto struct {
	Nombre string
	Precio int
}

func NewProducto(nombre string, precio int) Producto {
	return Producto{Nombre: nombre, Precio: precio}
}
