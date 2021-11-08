package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	result := x
	sub := x
	for start := 1.0;start < x;start +=  0.1{
		newSub := start*start - x
		if math.Abs(newSub) < sub{
			sub = math.Abs(newSub)
			result = start
		}
	}
	return result
}

func main() {
	fmt.Println(Sqrt(2))
}
