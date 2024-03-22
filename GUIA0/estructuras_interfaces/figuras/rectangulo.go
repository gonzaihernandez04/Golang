package figuras

type Rectangulo struct {
	esquinaInicial, esquinaFinal Punto
}

func NewRectangulo(esquinaInicial, esquinaFinal Punto) Rectangulo {
	return Rectangulo{esquinaInicial: esquinaFinal, esquinaFinal: esquinaFinal}
}
