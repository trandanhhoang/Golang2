package main

import "fmt"

// just start with "var" and "func" , so, can't x := 10
var c, python bool
var java bool = true

func main() {
	var i int
	x := 5
	var str = "hoang pro"
	y:= 3 +3i
	var empty string
	fmt.Printf("%v %T\n", y, y)
	fmt.Println(i, c, python, str, java, empty, x)
	fmt.Printf("%v %q\n", x, empty)
}

// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // alias for uint8
// rune // alias for int32     represents a Unicode code point
// float32 float64
// complex64 complex128
