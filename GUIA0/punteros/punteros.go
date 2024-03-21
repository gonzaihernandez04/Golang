package main

import "fmt"

func main() {
	var v1 int = 2
	var v2 int = 5
	Swap(&v1, &v2)

	fmt.Println(v1)
	fmt.Println(v2)
}

func Swap(px, py *int) {
	temp := *px
	*px = *py
	*py = temp

}
