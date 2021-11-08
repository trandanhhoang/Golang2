package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	k := math.Sqrt(float64(x*x + y*y)) // cab't use Float32 or Empty Type here
	var z uint = uint(k)
	fmt.Println(x, y, z)

	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Println(i,f,u)

	hoang := true
	fmt.Printf("hoang is of type %T\n", hoang)


}
