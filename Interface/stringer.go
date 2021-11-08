package main

import (
	"fmt"
	"strconv"
)

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return "Name is " + p.name + ",Age " + strconv.Itoa(p.age)
}

func main() {
	p1 := person{"hoang", 20}

	p2 := person{"anh", 20}
	fmt.Println(p1)
	fmt.Println(p2)
}
