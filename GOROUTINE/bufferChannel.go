package main

import "fmt"

func main() {
	//the second argument is buffer length
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3 // Just 3 times cause buffer = 3
	fmt.Println(<-ch)
	ch <- 4 // Just 3 times cause buffer = 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch) // buffer = 3, so this line error, we can
	// make it right by add "close(ch)" at line 13
}
