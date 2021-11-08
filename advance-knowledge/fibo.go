package main

import (
	"fmt"
	// "sync"
	// "time"
)

func fibo(length int) <-chan int {
	// var wg sync.WaitGroup
	c := make(chan int, length)
	// wg.Add(1)
	go func() {
		for i, j, k := 0, 1, 0; k < length; k++ {
			i, j = i+j, i
			c <- i
		}
		close(c)
		// wg.Done()
	}()
	// wg.Wait()

	return c
}

func main() {
	for output := range fibo(10) {
		fmt.Println(output)
	}
	// time.Sleep(2 * time.Second)
}
