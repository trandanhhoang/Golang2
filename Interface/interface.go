package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float32
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	// v := Vertex{3, 4}

	a = f // a MyFloat implements Abser
	// a = &v // a *Vertex implements Abser
	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a need an object, not pointer
	// a = v, if want u this line, make a func (v Vertex)
	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float32 {
	if f < 0 {
		return float32(-f)
	}
	return float32(f)
}

// type Vertex struct {
// X, Y float64
// }
//
// func (v *Vertex) Abs() float64 {
// return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }
