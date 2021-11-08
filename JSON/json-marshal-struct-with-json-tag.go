package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name        string `json:"full_name"`
	Age         int    `json:"age,omitempty"`
	Active      bool   `json:"-"`
	Order       *Order `json:"order,omitempty"`
	lastLoginAt string
	City        string `json:"city,omitempty"`
}

type Order struct {
	Name string
	Add  string
}

func main() {
	u, err := json.Marshal(User{Name: "Hoang", Age: 1, Active: true, lastLoginAt: "today", Order: &Order{}})
	// u, err := json.Marshal(User{Name: "Hoang", Age: 1, Active: true, lastLoginAt: "today", Order: nil, City: ""})
	// u, err := json.Marshal(&User{})
	// u, err := json.Marshal(User{})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(u)) // {"full_name":"Bob"}
	// lastLogin be eliminated because start with lower-case "l"

}
