package main

import "fmt"

const (
	roomChallengeFmt = "room:%s:challenge"
)

func main() {
	key := fmt.Sprintf(roomChallengeFmt, "hoang", "hhe")
	fmt.Println(key)
}
