package arrays

//var Array []int = make([]int, 10)

// Array cortado del uno al 5
//var Sliced []int = Array[0:5]

var X = []int{1, 2, 3, 4, 5}

// Copy

// Declaración e inicialización de las slices de origen y destino
var Origen []int = []int{1, 2, 3, 4, 5}
var Destino []int = make([]int, 3) // La longitud de destino debe ser igual o mayor a la de origen

// Copiar elementos desde la slice de origen a la slice de destino
// n := copy(destino, origen)
