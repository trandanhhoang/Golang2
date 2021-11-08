package main

import (
	"fmt"
	"time"
)

func fibo(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			{
				fmt.Println("i;m Here ", x)
				x, y = y, x+y
			}
		case <-quit:
			{
				return
			}
		}
	}
}

func async(c, quit chan int) {
	fmt.Println("i;m Here")
	for i := 0; i < 10; i++ { // if don't have c, this for will't run
		fmt.Printf("i %v %v \n", i, <-c) // this is a trigger for select.
	}
	time.Sleep(time.Minute) // this line for test what is trigger, it done, yup!
	quit <- 15
}

func main() {
	c, quit := make(chan int), make(chan int)
	//async new thread
	go async(c, quit)

	go fibo(c, quit)
	time.Sleep(time.Second) // if don't have this line, just you go in line 35
	time.Sleep(time.Minute) // if don't have this line, just you go in line 35
}
