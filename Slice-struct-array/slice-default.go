package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13, 12, 13, 14, 10}

	fmt.Println(s[0:10])

	s = s[:5]
	fmt.Println(s)

	fmt.Println(s[0:])

	fmt.Println(s[:])

}
