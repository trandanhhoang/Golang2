package main

import (
	"fmt"
	"sync"
	// "time"
)

func getInputChan() <-chan int {
	c := make(chan int, 100)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	go func() {
		for _, num := range arr {
			c <- num
		}
		close(c)
	}()

	return c
}

func getSquareChan(input <-chan int) <-chan int {
	c := make(chan int, 100)

	go func() {
		for num := range input {
			c <- num * num
		}
		close(c)
	}()

	return c
}

func merge(chanLink ...<-chan int) <-chan int {
	c := make(chan int, 100)
	// need WAIT GROUP FOR MORE TASK
	var wg sync.WaitGroup
	for _, output := range chanLink {
		wg.Add(1)
		for num := range output {
			c <- num
		}
		wg.Done()
	}
	go func() {
		close(c)
		wg.Wait()
	}()

	return c
}

func main() {
	chanInputNums := getInputChan()

	chanOptSqr1 := getSquareChan(chanInputNums)
	chanOptSqr2 := getSquareChan(chanInputNums)

	chanMergedSqr := merge(chanOptSqr1, chanOptSqr2)

	sqrSum := 0

	for num := range chanMergedSqr {
		sqrSum += num
	}

	fmt.Println("Sum of squares between 0-9 is", sqrSum)
}
