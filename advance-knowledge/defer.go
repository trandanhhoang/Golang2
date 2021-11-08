package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Line 0") // close connect when open
	defer fmt.Println("Line 1")
	fmt.Println("Line 2")
	fmt.Println("Line 3")

	a := 12
	defer fmt.Println(a)
	a = 24
}
