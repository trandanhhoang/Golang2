package main

import (
	"fmt"
	"time"
	"math/rand"

)

func main() {
	fmt.Println("The time is", time.Now())
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(11))
}
