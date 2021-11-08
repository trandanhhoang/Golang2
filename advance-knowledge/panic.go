package main

import (
	"fmt"
	"log"
)

func main() {
	// a := 10
	// b := 0
	fmt.Println("Start MAIN\n")
	panic("divide by zero2")

	panicker()

	fmt.Println("end MAIN")

}

func panicker() {
	panic("divide by zero2")
	panic("divide by zero")
	log.Fatal("haha")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("chia cho 0")
			fmt.Println(err, "\n")
		}
	}()
	fmt.Println("panicker")
}
