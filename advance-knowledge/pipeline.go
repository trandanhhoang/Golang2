package main

import "fmt"

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}

	inputChan := genPipe(arr)

	c1 := fanOut(inputChan)
	c2 := fanOut(inputChan)
	c3 := fanOut(inputChan)
	c4 := fanOut(inputChan)

	c := fanIn(c1, c2, c3, c4)

	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += <-c
	}
	fmt.Println(sum)

}

func genPipe(arr []int) <-chan int {
	ch := make(chan int)
	go func() {
		for ele := range arr {
			ch <- ele
		}
		close(ch)
	}()
	return ch
}

func fanOut(ch <-chan int) <-chan int {
	chOut := make(chan int)
	go func() {
		for ele := range ch {
			chOut <- ele * ele
		}
		close(chOut)
	}()
	return chOut
}

func fanIn(chs ...<-chan int) <-chan int {
	chOut := make(chan int)
	go func() {
		for _, ch := range chs {
			for ele := range ch {
				chOut <- ele
			}
		}
	}()
	return chOut
}
