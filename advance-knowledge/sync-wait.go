package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go read("sheep", &wg)
	go read("fish", &wg)
	wg.Wait()
	// go read("sheep")
	// go read("fish")
	// time.Sleep(time.Minute)
	fmt.Println("hoang")

}

// func read(x string) {
// for i := 0; i < 5; i++ {
// fmt.Println(x, i)
// time.Sleep(time.Second * 1)
// }
// }
func read(x string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(x, i)
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}
