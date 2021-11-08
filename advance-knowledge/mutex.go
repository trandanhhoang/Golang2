package main

// chua nam ro~, delll hieu cai gi ?
// doc lai mutex in c++
// what is thread
// what is deadlock, what is race condition

import (
	"fmt"
	"sync"
)

var counter = 0

func main() {
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}

	for i := 0; i < 10000; i++ {
		wg.Add(1)

		go increase(&wg, &mu)
	}
	wg.Wait()
	fmt.Println("cunter", counter)
}

func read(wg *sync.WaitGroup, mu *sync.RWMutex) {
	fmt.Println("hello ", counter)

	wg.Done()
}

func increase(wg *sync.WaitGroup, mu *sync.RWMutex) {
	// mu.Lock()
	counter++
	// mu.Unlock()
	wg.Done()

}
