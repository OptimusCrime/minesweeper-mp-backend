package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const maxMessageSize = 512

func (client *Client) readPump() {
	defer func() {
		client.Hub.Unregister <- client
		client.Conn.Close()
	}()
	client.Conn.SetReadLimit(maxMessageSize)
	client.Conn.SetReadDeadline(time.Now().Add(pongWait))
	client.Conn.SetPongHandler(func(string) error { client.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		messageString := string(message[:])
		log.Println(messageString)

		var result map[string]interface{}
		jsonErr := json.Unmarshal([]byte(messageString), &result)
		if jsonErr != nil {
			// TODO shiet
			return
		}

		if val, ok := result["type"]; ok {
			switch val {
			case "JOIN_LOBBY":
				handleHandleIncomingLobbyJoinMessage(client, messageString)
				break
			}
		}
	}
}
