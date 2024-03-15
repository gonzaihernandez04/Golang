package main;


import(
	"fmt"
	"github.com/gonzajs04/Golang/tree/master/Funciones/operaciones"
	"github.com/gonzajs04/Golang/tree/master/Funciones/arrays"
);



func main(){
	fmt.Println("Hola mundo");

	fmt.Printf("El valor de la suma es de %v\n", operaciones.Sumar(8,2));
	fmt.Printf("El valor de la resta es de %v\n", operaciones.Restar(8,2));
	fmt.Printf("El valor de la multiplicacion es de %v\n", operaciones.Multiplicar(8,2));


	// Append: agrega un elemento a la slice usando append

	var x = append(arrays.X,6);

	fmt.Println("Array append ", x)


	// Copy

	copyArray := copy(arrays.Destino, arrays.Origen)


	// Maps o slice asociativo:
	var datos = make(map[string] int);
	datos["edad"] = 19;
	datos["telefono"] = 1232123451;

	// Delete key and value from slices
	delete(datos,"telefono")


	fmt.Println("Map o asociativo en golang", datos["edad"] , datos["telefono"])


	fmt.Println("\n",arrays.Origen)
	fmt.Println("\n",arrays.Destino)
	fmt.Println("\n",copyArray)


}