package socket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("w", w)
	fmt.Println("r", r)
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			break
		}
		// fmt.Println("mt", mt)
		// fmt.Println("message", message)
		// fmt.Printf("message %s \n", message)

		err = c.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}
