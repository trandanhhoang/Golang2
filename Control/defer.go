package main

import "fmt"

func main() {
	fmt.Println("counting")
	defer fmt.Println()
	for i := 0; i < 10; i++ {
		defer fmt.Print(i,"\t")
	}

	fmt.Println("done")
}
