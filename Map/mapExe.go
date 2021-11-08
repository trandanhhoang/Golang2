package main

import (
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _,v := range strings.Split(s," "){
		m[v]++
	}
	
	return m
}

func main() {
	var s = "I ate a donut. Then I ate another donut."
	fmt.Println(WordCount(s))
}
