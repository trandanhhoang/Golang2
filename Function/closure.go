package main

import(
	"fmt"
)

func clo() func(int) int{
	sumOuter:= 0
	return func(x int) int{
		sumOuter += x 
		return sumOuter
	}
} 

func main(){
	neg, pos := clo(),clo()
	for i:=0;i<10;i++{
		fmt.Print(neg(i*-2),"\t")
		fmt.Print(pos(i),"\n")
	}
	
}