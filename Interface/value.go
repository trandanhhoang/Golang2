package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

type my_float float64

func (t *T) M() {
	fmt.Println(t.S)
}
func (f my_float) M() {
	fmt.Println(f)
}

func main() {
	var i I
	t := T{"Hello World"}
	i = &t
	describe((i))
	i.M()

	f := my_float(3.14)
	i = f
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("value = %v, Type is %T\n", i, i)
}
