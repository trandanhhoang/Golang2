package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// RANGE use for "slice" or "map"
//first index is index,second index is copy of ele
func main() {
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
		if i == 5 {
			v = 10
		}
		fmt.Printf("2^%d = %d\n", i, v)

	}

	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}
}
