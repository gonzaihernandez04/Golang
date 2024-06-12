package rutinas

//TODO:Buscar forma de hacer rutas dinamicas para los archivos para TEST
import (
	"ayp2_tp/ejercicios"
	"ayp2_tp/funcionesCSV"

	"log"
	"math"
	"reflect"
	"strings"

	"github.com/gocarina/gocsv"
)

type Rutina struct {
	Nombre                 string                  `csv:"nombre"`
	DuracionEstimada       int                     `csv:"tiempoEstimado"`
	CaloriasQuemadas       int                     `csv:"caloriasQuemadas"`
	TipoEjericio           []string                `csv:"tipoEjercicio"`
	PuntosPorTipoEjercicio []int                   `csv:"puntosPorTipoEjercicio"`
	Dificultad             int                     `csv:"dificultad"`
	Ejercicios             []*ejercicios.Ejercicio `csv:"ejercicio"`
}

// Todo todo
func NuevaRutina(nombre string, ejercicios []*ejercicios.Ejercicio) *Rutina {
	var duracionEstimada, caloriasQuemadas, dificultad int
	var tipoEjercicio []string
	var puntosPorTipoEjercicio []int
	for _, ejercicio := range ejercicios {
		duracionEstimada += ejercicio.TiempoEstimado
		caloriasQuemadas += ejercicio.CaloriasQuemadas
		dificultad += ejercicio.Dificultad
		tipoEjercicio = append(tipoEjercicio, ejercicio.TipoEjericio...)
		puntosPorTipoEjercicio = append(puntosPorTipoEjercicio, ejercicio.PuntosPorTipoEjercicio...)

	}
	// Lo transformo a un valor entero a su valor mas cercano:
	// 3+2+2+1+1 = 9/5 = 1.8
	// 1.8 = 2
	dificultad = int(math.Round(float64(dificultad) / float64(len(ejercicios))))

	return &Rutina{
		Nombre:                 nombre,
		DuracionEstimada:       duracionEstimada,
		CaloriasQuemadas:       caloriasQuemadas,
		TipoEjericio:           tipoEjercicio,
		PuntosPorTipoEjercicio: puntosPorTipoEjercicio,
		Dificultad:             dificultad,
		Ejercicios:             ejercicios,
	}
}

// Devuelve un slice con todos los ejercicios dentro de un archivo csv abierto
func obtenerSliceRutinas() []*Rutina {

	rutinas := []*Rutina{}
	/*ESPECIFICO CORRECTAMENTE LA RUTA.*/
	// ruta := "rutinas/rutinas.csv"
	ruta, _ := funcionesCSV.BuscarArchivoCSV("rutinas.csv")
	archivoRutinas, err := funcionesCSV.AbrirArchivoCSV(ruta)

	defer funcionesCSV.CerrarArchivoCSV(archivoRutinas)

	if err != nil {
		panic(err)
	}

	if err := gocsv.UnmarshalFile(archivoRutinas, &rutinas); err != nil {
		panic(err)
	}

	return rutinas
}

// Devuelve true o false si un ejercicio ya existe o no en el csv de ejercicios
func RutinasExiste(nuevaRutina *Rutina) bool {
	rutinas := obtenerSliceRutinas()

	for _, rutina := range rutinas {
		if reflect.DeepEqual(rutina, nuevaRutina) { //Compara dentro del struct todos los parametros cada atributo
			return true
		}
	}

	return false
}

// Devuelve true o false, segun si dos ejercicios son iguales o no
func RutinasIguales(rutina1, rutina2 *Rutina) bool {
	return reflect.DeepEqual(rutina1, rutina2)
}

// Devuelve el indice de un ejercicio en la lista de ejercicios del csv, o -1 si el ejercicio no esta
func buscarIndiceRutina(rutinaBuscada *Rutina) int {
	ejercicios := obtenerSliceRutinas()

	indiceEjercicio := -1
	for i, ejercicio := range ejercicios {
		if RutinasIguales(ejercicio, rutinaBuscada) {
			indiceEjercicio = i
		}
	}

	return indiceEjercicio
}

func AltaRutina(nuevaRutina *Rutina) {
	// Checkear si ya existe el ejercicio en la lista
	if RutinasExiste(nuevaRutina) {
		return
	}

	// Abrir el archivo y cerrarlo automaticamente al salir de la funcion
	// ruta := "rutinas/rutinas.csv"
	ruta, _ := funcionesCSV.BuscarArchivoCSV("rutinas.csv")
	archivoRutinas, _ := funcionesCSV.AbrirArchivoCSV(ruta)

	defer funcionesCSV.CerrarArchivoCSV(archivoRutinas)

	rutinas := obtenerSliceRutinas()

	// Agregar el nuevo ejercicio al slice de ejercicios
	rutinas = append(rutinas, nuevaRutina)

	// Guardar los datos del slice en el archivo csv
	if _, err := archivoRutinas.Seek(0, 0); err != nil {
		panic(err)
	}

	gocsv.MarshalFile(&rutinas, archivoRutinas)

}

func BajaRutina(rutinaParaRemover *Rutina) {
	if !RutinasExiste(rutinaParaRemover) {
		// fmt.Println("EL EJERCICIO NO ESTA EN LA LISTA")
		return
	}

	// Abrir el archivo y cerrarlo automaticamente al salir de la funcion
	// ruta := "rutinas/rutinas.csv"
	ruta, _ := funcionesCSV.BuscarArchivoCSV("rutinas.csv")
	archivoRutina, _ := funcionesCSV.AbrirArchivoCSV(ruta)

	defer funcionesCSV.CerrarArchivoCSV(archivoRutina)

	rutinas := obtenerSliceRutinas()

	indiciceRutina := buscarIndiceRutina(rutinaParaRemover)

	if indiciceRutina >= 0 {
		rutinas = append(rutinas[:indiciceRutina], rutinas[indiciceRutina+1:]...)
	}

	// Elimina todos los datos del csv
	if err := archivoRutina.Truncate(0); err != nil {
		panic(err)
	}

	// Va hasta arriba del archivo csv
	if _, err := archivoRutina.Seek(0, 0); err != nil {
		panic(err)
	}

	// Guarda los datos del slice en el csv
	if err := gocsv.MarshalFile(&rutinas, archivoRutina); err != nil {
		log.Fatalf("Error al escribir en el archivo CSV: %s", err)
	}
}

func ConsultarRutina(nombreBuscado string) []*Rutina {
	// Abrir el archivo y cerrarlo automaticamente al salir de la funcion
	// ruta := "rutinas/rutinas.csv"
	ruta, _ := funcionesCSV.BuscarArchivoCSV("rutinas.csv")
	archivoRutinas, _ := funcionesCSV.AbrirArchivoCSV(ruta)
	defer funcionesCSV.CerrarArchivoCSV(archivoRutinas)

	// Crear el slice que contiene los ejercicios buscados
	var rutinasBuscadas []*Rutina
	rutinas := obtenerSliceRutinas()

	// Si el parametro es un string, busca por nombre de ejercicio, sino busca
	// por calorias

	for _, rutina := range rutinas {
		// Si el nombre del ejericicio contiene el ejercicio buscado, se agrega al slice
		if strings.Contains(strings.ToLower(rutina.Nombre), strings.ToLower(nombreBuscado)) {
			rutinasBuscadas = append(rutinasBuscadas, rutina)
		}
	}

	return rutinasBuscadas
}

// Devuelve un array con las rutinas disponibles
func ListarRutinas() []*Rutina {
	// ruta := "rutinas/rutinas.csv"
	ruta, _ := funcionesCSV.BuscarArchivoCSV("rutinas.csv")
	archivoRutinas, _ := funcionesCSV.AbrirArchivoCSV(ruta)
	defer funcionesCSV.CerrarArchivoCSV(archivoRutinas)

	// Crear el slice que contiene los rutinas buscados
	rutinas := obtenerSliceRutinas()

	/*
		for _, rutina := range rutinas {
			fmt.Println(rutina.Nombre)
		}
	*/

	return rutinas
}

func ModificarRutina(nombre string, nuevaRutina *Rutina) {
	archivoRutina, _ := funcionesCSV.AbrirArchivoCSV("rutinas.csv")
	defer funcionesCSV.CerrarArchivoCSV(archivoRutina)

	// Crear el slice que contiene los ejercicios buscados
	rutinas := obtenerSliceRutinas()
	indiceRutina := buscarIndiceRutina(ConsultarRutina(nombre)[0])

	rutinas[indiceRutina] = nuevaRutina

	if _, err := archivoRutina.Seek(0, 0); err != nil {
		panic(err)
	}

	gocsv.MarshalFile(&rutinas, archivoRutina)
}

// Función para verificar si un ejercicio contiene alguno de los tipos deseados
func ContieneTipoEjercicio(ejercicioTipos []string, tiposDeseados []string) bool {
	for _, tipo := range ejercicioTipos {
		for _, deseado := range tiposDeseados {
			if tipo == deseado {
				return true
			}
		}
	}
	return false
}

// -------------------------------------------------- GeneracionAutomagicaDeRutinas2 --------------------------------------------------

func GeneracionAutomagicaDeRutinas2_Aux(nombre string, caloriasTotales int) *Rutina {
	var mejorRutina *Rutina
	mejorDuracion := int(^uint(0) >> 1) // Representa el valor máximo de un int

	// Crear la rutina inicial
	rutinaInicial := &Rutina{
		Nombre: nombre,
	}

	ejerciciosDisponibles := ejercicios.ListarTodosEjercicios()
	// Llamar a la función de backtracking
	backtrack(0, caloriasTotales, 0, rutinaInicial, ejerciciosDisponibles, &mejorRutina, &mejorDuracion)

	rutinaDefinitiva := NuevaRutina(mejorRutina.Nombre, mejorRutina.Ejercicios)

	return rutinaDefinitiva
}

// Funcion para encontrar mejor combinacion para los ejercicios, donde la pri
func backtrack(indice int, caloriasRestantes int, duracionActual int, rutinaActual *Rutina, ejerciciosFiltrados []*ejercicios.Ejercicio, mejorRutina **Rutina, mejorDuracion *int) {
	if caloriasRestantes <= 0 {
		if duracionActual < *mejorDuracion {
			*mejorRutina = &Rutina{
				Nombre:                 rutinaActual.Nombre,
				DuracionEstimada:       duracionActual,
				CaloriasQuemadas:       rutinaActual.CaloriasQuemadas,
				TipoEjericio:           rutinaActual.TipoEjericio,
				PuntosPorTipoEjercicio: rutinaActual.PuntosPorTipoEjercicio,
				Dificultad:             rutinaActual.Dificultad,
				Ejercicios:             append([]*ejercicios.Ejercicio(nil), rutinaActual.Ejercicios...),
			}
			*mejorDuracion = duracionActual
		}
		return
	}

	if indice == len(ejerciciosFiltrados) {
		return
	}

	// No incluir el ejercicio actual
	backtrack(indice+1, caloriasRestantes, duracionActual, rutinaActual, ejerciciosFiltrados, mejorRutina, mejorDuracion)

	// Incluir el ejercicio actual
	ejercicio := ejerciciosFiltrados[indice]
	rutinaActual.Ejercicios = append(rutinaActual.Ejercicios, ejercicio)
	rutinaActual.CaloriasQuemadas += ejercicio.CaloriasQuemadas

	backtrack(indice+1, caloriasRestantes-ejercicio.CaloriasQuemadas, duracionActual+ejercicio.TiempoEstimado, rutinaActual, ejerciciosFiltrados, mejorRutina, mejorDuracion)
	// Borra ejercicio actual de la rutina para probar nuevas combinaciones
	rutinaActual.Ejercicios = rutinaActual.Ejercicios[:len(rutinaActual.Ejercicios)-1]
	rutinaActual.CaloriasQuemadas -= ejercicio.CaloriasQuemadas
}

// -------------------------------------------------- GeneracionAutomagicaDeRutinas3 --------------------------------------------------

func GenerarRutinaAutomatica3(nombre string) {

}
