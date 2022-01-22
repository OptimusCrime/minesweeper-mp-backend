package main

import (
	"log"
	"net/http"
)

func main() {
	createBoard()
	hub := newHub()
	go hub.run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	err := http.ListenAndServe(":8094", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
