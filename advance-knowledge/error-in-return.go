package main

import "fmt"

func main() {
	val, err := errFunc()
	fmt.Println(val)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

}

func errFunc() (int, error) {
	return 5, fmt.Errorf("can't devide by 0")
}
