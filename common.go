package main

import "time"

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)
