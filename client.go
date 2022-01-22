package main

import (
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	Hub      *Hub
	Conn     *websocket.Conn
	Send     chan []byte
	Room     *Room
	Nickname string
	Points   int
	HasWon   bool
}

func createClient(hub *Hub, conn *websocket.Conn) *Client {
	client := &Client{
		Hub:    hub,
		Conn:   conn,
		Send:   make(chan []byte, 256),
		Points: 0,
		HasWon: false,
	}

	client.Hub.Register <- client

	return client
}

func findOrGenerateRandomNickname(nickname string) string {
	if len(nickname) > 0 {
		return nickname
	}

	return strconv.Itoa(rand.Intn(10000000-1000) + 1000)
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, response http.ResponseWriter, request *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(response, request, nil)

	if err != nil {
		log.Println(err)
		return
	}

	client := createClient(hub, conn)

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
