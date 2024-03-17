package structs

type Rectangulo struct {
	LL float32
	LT float32
}

func NewRect(l1 float32, l2 float32) Rectangulo {
	return Rectangulo{LL: l1, LT: l2}
}

/*func (r *Rectangulo) GetArea() float32 {
	l := distance(r.LT, r.LL, r.LT, r.LL)
	w := distance(r.LT, r.LL, r.LT, r.LL)

	return l * w

}
*/
