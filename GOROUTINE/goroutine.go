package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}

func main() {
	// 2 go not working, go in line 18 still not working, JUST only "go" in line 17 working
	// create new thread
	go say("world") // remove "go" to see change
	say("hello")
}
