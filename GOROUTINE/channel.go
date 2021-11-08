package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	c := make(chan int) //Like maps and slices, channels must be created before use:
	s := []int{2, -5, 1, -2, 9}

	// it has order , WHY ???????????
	go sum(s[:len(s)/2], c) //0,1
	go sum(s[len(s)/2:], c) //2,3,4

	x := <-c
	y := <-c

	fmt.Println(x, y, x+y, len(s))
}
