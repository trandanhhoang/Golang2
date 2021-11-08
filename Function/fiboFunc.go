package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
// (0, 1, 1, 2, 3, 5, ...).
func fibonacci() func(int) int {
	oldSum:= 0
	sumOuter:= 1
	return func(x int) int {
		if x==0 {
			return 0
		} 
		result := oldSum + sumOuter
		oldSum = sumOuter
		sumOuter = result
		return oldSum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(i," = \t")
		fmt.Println(f(i))
	}
}
