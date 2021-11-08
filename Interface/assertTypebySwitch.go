package main

import(
	"fmt"
)

func main(){
	assert("hello")
	assert(5)
	assert(5.3)
	assert(true)
}

func assert(i interface{}){
	switch v:=i.(type) {
	case string: fmt.Printf("string ne : %v,%T\n",v,v)
	case int: fmt.Printf("Int ne : %v,%T\n",v,v)
	default: fmt.Printf("%v,%T\n",v,v)
	}
}
