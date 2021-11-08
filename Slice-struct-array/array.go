package main

import "fmt"

func main() {
	var a [2]string
	fmt.Println(a)
	b := make([]string, 2)
	fmt.Println(b)
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2.0, 3, 5, 7, 11, 13}
	fmt.Printf("%d\n", primes[:len(primes)])
	fmt.Printf("%s", a)
}
