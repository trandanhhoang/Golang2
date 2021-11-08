package main

// type mess struct
import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var upgrader = websocket.Upgrader{
	// CheckOrigin: func(r *http.Request) bool {
	// return true
	// },
}

type ChatMessages struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcaster = make(chan ChatMessages)
var rdb *redis.Client

func handleConn(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer ws.Close()
	if rdb.Exists("chat_messages").Val() != 0 {
		fullData, err := rdb.LRange("chat_messages", 0, -1).Result()
		if err != nil {
			panic(err)
		}

		for _, val := range fullData {
			var prevMsg ChatMessages
			if err := json.Unmarshal([]byte(val), &prevMsg); err != nil {
				log.Print(err)
			}
			messageClient(ws, prevMsg)
		}
	}
	clients[ws] = true

	log.Println("websocket.Conn", &ws)
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s, type of recv: %T", message, message)
		// convert before send message
		var msg ChatMessages
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println(err)
		}
		broadcaster <- msg
		log.Printf("mt: %T", mt)
	}
}

func handleMessage() {
	for {
		select {
		case msg := <-broadcaster:
			messageToAllClient(msg)
			saveIntoRedis(msg)
		}
	}
}

func saveIntoRedis(msg ChatMessages) {
	arrByteOfMsg, err := json.Marshal(msg)
	if err != nil {
		log.Print(err)
	}
	if err := rdb.RPush("chat_messages", arrByteOfMsg).Err(); err != nil {
		panic(err)
	}
}

func messageToAllClient(msg ChatMessages) {
	for client := range clients {
		messageClient(client, msg)
	}
}

func messageClient(wsOfClient *websocket.Conn, msg ChatMessages) {
	arrByteOfMsg, err := json.Marshal(msg)
	if err != nil {
		log.Print(err)
	}

	if err := wsOfClient.WriteMessage(1, arrByteOfMsg); err != nil {
		log.Println("write:", err)
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	redisURL := os.Getenv("REDIS_URL")
	//server a static html using go web server
	http.Handle("/", http.FileServer(http.Dir("./public")))
	// handle when have a client join
	http.HandleFunc("/websocket", handleConn)
	// connect redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//
	go handleMessage()
	fmt.Println("server is serve at http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
