package main

import "fmt"

func main() {
	pow := make([]int, 10)
	//omit the second variable, the "value"
	for i := range pow {
		fmt.Print(i)
		if i == 9 {
			fmt.Print("\n")
		}
		pow[i] = 1 << uint(i) // == 2**i
	}
	// omit by "_"
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	fmt.Printf("%d\n", len(pow))
}
