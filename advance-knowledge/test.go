package main

import "fmt"

func main() {
	c, quit := make(chan int), make(chan int)
	//async new thread

	func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		quit <- 15
	}()

	for {
		select {
		case v := <-c:
			{
				fmt.Println(v)
			}
		case <-quit:
			{
				fmt.Println("done!")
				return
			}
		}
	}

}
