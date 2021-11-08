package main

import "https://golang.org/x/tour/pic"

// Implement Pic. It should return a slice of length dy,
//  each element of which is a slice of dx 8-bit unsigned integers.
//   When you run the program, it will display your picture, 
//   interpreting the integers as grayscale (well, bluescale) values.
func Pic(dx, dy int) [][]uint8 {
	ss := make([][]uint8,dy)
	for y:=0;y<dy;y++{
		s := make([]uint8,dx)
		ss[y] = s
	}

	return ss
}

func main() {
	pic.Show(Pic)
}
