package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(arr)
	arrStruct := []struct {
		x int
		y bool
	}{{1, true}, {2, false}}

	fmt.Println(arrStruct)

}
