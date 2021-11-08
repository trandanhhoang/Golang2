package main

import (
	"fmt"
	"runtime"
)

func main() {
	numberP := runtime.NumCPU()
	fmt.Println(numberP)
}
