package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Today is ")
	today := time.Now().Weekday()
	fmt.Println(today)
	fmt.Println("Tomorrow is ", today +1)

	switch time.Wednesday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	
	}

}
