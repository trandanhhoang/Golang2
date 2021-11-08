package main

import (
	"fmt"
	"runtime"
	"time"
)

func squares(c chan int) {
	for {
		fmt.Printf("Length of channel c is %v and capacity of channel c is %v\n", len(c), cap(c))
	}
}

func main() {
	fmt.Println("main() started")
	c := make(chan int, 3)
	go squares(c)

	time.Sleep(time.Second)
	fmt.Println("active goroutines", runtime.NumGoroutine())
	c <- 1
	c <- 2
	c <- 3
	c <- 4 // blocks here
	c <- 5 // blocks here

	fmt.Println("main() stopped")
}
