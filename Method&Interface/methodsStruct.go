package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

type myInt int

func (v myInt) Abs() int {
	if v < 0{
		return int(-v)
	}
	return int(v)
}

func main() {
	x := myInt(-6)
	fmt.Println(x.Abs())
}
