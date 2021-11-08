package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	fmt.Println(m)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["hehe"] = Vertex{
		1, 2.0,
	}
	fmt.Println(m["Bell Labs"])
	ele := m["Bell Labs"]
	ele.Lat = 10
	m["Bell Labs"] = ele

	fmt.Println(m)
	fmt.Printf("%T\n", m)

}
