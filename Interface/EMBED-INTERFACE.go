package main

import (
	"fmt"
)

type animal1 interface {
	speak()
}

type animal2 interface {
	say()
}

type animal3 interface {
	animal1
	animal2
}

type Dog struct{}

func (dog Dog) say() {
	fmt.Println("woop woop!")
}

func (dog Dog) speak() {
	fmt.Println("hello Hoang!")
}

func main() {
	dog := Dog{}

	var milu animal1 = dog

	milu.speak()

	var kiki animal2 = dog
	// fmt.Println(*kiki)
	kiki.say()
	// kiki.speak()

	milu.speak()
}
