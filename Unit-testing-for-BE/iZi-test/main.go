package main

import (
	"fmt"
	"time"
)

var ZeroString string
var ZeroMap map[string]interface{}

func SetTimeout(f func(), d time.Duration) *time.Timer {
	t := time.NewTimer(d)

	go func() {
		<-t.C
		f()
	}()

	return t
}

func hoang() {
	fmt.Println("kk")
}

func main() {
	i := SetTimeout(hoang, 3)
	fmt.Println(i)
}
