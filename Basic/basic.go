package main

import (
	"fmt"
)

func sum(ele int) (x, y int) {
	x = ele + 10
	y = ele - 10
	return
}

func split(ele int) (int, int) {
	x := ele + 5
	y := ele - 5
	return x, y
}

func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println(sum(5))
	fmt.Println(split(5))

}
