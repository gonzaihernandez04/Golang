package opciones

import "fmt"

func ElegirOpciones(opcionPrueba int) int {
	bandera := false
	var opcion int

	for !bandera {
		fmt.Println("Elija una opcion: 1, 2, 3, 4: ")

		_, err := fmt.Scanln(&opcion)

		if err != nil {
			fmt.Println("Error al leer la entrada:", err)
			opcion = opcionPrueba

		}

		switch {
		case opcion == 1:
			fmt.Printf("\n Elegiste la opcion %v", opcion)
			bandera = true

		case opcion == 2:
			fmt.Printf("\n Elegiste la opcion %v", opcion)
			bandera = true

		case opcion == 3:
			fmt.Printf("\n Elegiste la opcion %v", opcion)
			bandera = true

		default:
			fmt.Println("\n No existe la opcion elegida")

		}

	}

	return opcion

}
