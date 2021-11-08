package main

import (
	"fmt"
	"sync"
)

var sum int

var wg sync.WaitGroup

var m sync.Mutex

func main() {
	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go add(1)
	}

	wg.Wait()
	fmt.Println(sum)
}

func add(x int) {
	m.Lock()
	sum += x
	m.Unlock()
	wg.Done()
}
