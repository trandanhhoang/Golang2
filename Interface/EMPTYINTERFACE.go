package main
import(
	"fmt"
)
// An empty interface may hold values of any type.
// (Every type implements at least zero methods.)
// Empty interfaces are used by code that handles values of unknown type.
// For example, fmt.Print takes any number of arguments of type interface{}.
func main(){
	var i interface{}
	describe(i)
	i = "hello"
	describe(i)
	i = 5
	describe(i)
	//type assertion
	s,ok := i.(string)
	fmt.Print(s,ok)
}
func describe(i interface{}){
	fmt.Printf("%v, %v\n",i,i)
}