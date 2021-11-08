package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	number := 50
	number_of_worker := 12
	ch := make(chan int, number)
	taker := make(chan int, number)

	for i := 1; i <= number_of_worker; i++ {
		go worker(ch, taker)
	}

	for i := 1; i <= number; i++ {
		ch <- i
	}

	close(ch) // when don't want ch <-

	for i := 1; i <= number; i++ {
		fmt.Println(i, " : ", <-taker)
	}
	elapsed := time.Since(start)
	log.Println("Programe took ", elapsed)
}

func worker(ch <-chan int, taker chan<- int) {
	for val := range ch {
		taker <- fibo(val)
	}
}

func fibo(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}
