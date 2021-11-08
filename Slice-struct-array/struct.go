package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	x := Vertex{X: 1, Y: 2}
	y := x
	x.Y = 4
	y.Y = 8
	fmt.Println(x.X + x.Y)
	// x is x, y still y
	fmt.Println(x)
	fmt.Println(y)
	p := &x
	p.X = 1e9 // no need (*p).X
	// p.X = x.X
	fmt.Println(x)
}
