package main

import (
	"fmt"
	"strconv"
	"time"
)

func go1(c1 chan string) {
	for i := 0; i < 5; i++ {
		x := "hello string c1" + strconv.Itoa(i)
		c1 <- x
		time.Sleep(time.Second) // if don't have this line, just you go in line 35

	}
	close(c1)
}
func go2(c2 chan string) {
	for i := 0; i < 5; i++ {
		x := "hello string c2" + strconv.Itoa(i)
		c2 <- x
		time.Sleep(time.Second) // if don't have this line, just you go in line 35
	}
	close(c2)
}

func main() {
	c1, c2 := make(chan string), make(chan string)
	go go1(c1)
	go go2(c2)
	for {
		select {
		case x, ok1 := <-c1:
			fmt.Println(x)
			if !ok1 {
				c1 = nil
			}
		case y, ok2 := <-c2:
			fmt.Println(y)
			if !ok2 {
				c2 = nil
			}
		}

		if c1 == nil && c2 == nil {
			break
		}
	}
}
