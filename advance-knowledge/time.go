package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()
	log.Println("Start at ", start)
	for i := 1; i <= 42; i++ {
		fibo(i)
	}

	elapsed := time.Since(start)
	log.Println("Programe took ", elapsed)
}

func fibo(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}
