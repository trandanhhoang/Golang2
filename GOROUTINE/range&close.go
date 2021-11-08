package main

import (
	"fmt"
)

func fibo(len int, c chan int) {
	x, y := 0, 1
	for i := 0; i <= 10; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// v,ok := <-c
// fmt.Println(v,ok)

func main() {
	c := make(chan int, 1)
	go fibo(cap(c), c)
	// 1 khi thay c con` la van~ nhan. khi nao close ms dong'.
	for i := range c {
		fmt.Printf("c = %v,len = %v\n", c, len(c))
		fmt.Println("i", i)
	}
	x, ok := <-c
	fmt.Printf("x = %v, ok = %v", x, ok)

}
