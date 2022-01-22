package main

import (
	"bytes"
	"errors"
	"math/rand"
)

var roomNameOptions = [...]string{"A", "B", "C", "D", "E", "F", "H", "J", "K", "L", "M", "N", "P", "R", "T", "X", "Y", "3", "4", "7", "8", "9"}
var roomNameLength = 5

type Room struct {
	Name        string
	Open        bool
	Clients     []*Client
	CurrentTurn int
}

func createNewRoom() *Room {
	return &Room{
		Name:        createRandomRoomName(),
		Open:        true,
		Clients:     []*Client{},
		CurrentTurn: 0,
	}
}

func getPlayersInRoom(client Client) []MessagePlayer {
	var players []MessagePlayer

	for _, player := range client.Room.Clients {
		players = append(players, MessagePlayer{
			Nickname:      player.Nickname,
			CurrentPlayer: *player == client,
		})
	}

	return players
}

func createRandomRoomName() string {
	var byteBuffer bytes.Buffer

	for i := 0; i < roomNameLength; i++ {
		byteBuffer.WriteString(roomNameOptions[rand.Intn(len(roomNameOptions))])
	}

	return byteBuffer.String()
}

func findOrCreateRoom(client *Client, roomCode string) (*Room, error) {
	if len(roomCode) > 0 {
		for _, room := range client.Hub.Rooms {
			if room.Name == roomCode {
				return client.attachRoom(room)
			}
		}

		return nil, errors.New("no such room")
	}

	newRoom := createNewRoom()

	client.Hub.Rooms = append(client.Hub.Rooms, newRoom)

	return client.attachRoom(newRoom)
}

func (client *Client) attachRoom(room *Room) (*Room, error) {
	room.Clients = append(room.Clients, client)
	client.Room = room

	return room, nil
}
