package main

import (
	"fmt"
	"math"
)

const (
	l = math.Pi
	a = iota
	b = iota
	c = iota
)

const (
	d, e, f = iota, iota, iota
)

func main() {
	fmt.Println(a, b, c, d, e, f)
}
