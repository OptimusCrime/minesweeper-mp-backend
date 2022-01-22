package main

import "encoding/json"

func handleHandleIncomingLobbyJoinMessage(client *Client, message string) {
	var joinLobbyIncomingMessage JoinLobbyIncomingMessage
	json.Unmarshal([]byte(message), &joinLobbyIncomingMessage)

	room, err := findOrCreateRoom(client, joinLobbyIncomingMessage.RoomCode)
	nickname := findOrGenerateRandomNickname(joinLobbyIncomingMessage.Nickname)

	client.Nickname = nickname

	if err != nil {
		newMessageStruct := LobbyFailedJoinOutboundMessage{
			Type:     "JOIN_LOBBY_FAILED",
			RoomCode: joinLobbyIncomingMessage.RoomCode,
		}
		newMessageMarshal, _ := json.Marshal(newMessageStruct)

		client.Send <- newMessageMarshal
	} else {
		for _, clientInRoom := range room.Clients {
			newMessageStruct := LobbySuccessJoinOutboundMessage{
				Type:     "JOIN_LOBBY_SUCCESS",
				Nickname: nickname,
				Room: MessageRoom{
					RoomCode:    room.Name,
					CurrentTurn: room.CurrentTurn,
					HasStarted:  room.Open,
				},
				Players: getPlayersInRoom(*clientInRoom),
			}

			newMessageMarshal, _ := json.Marshal(newMessageStruct)

			// Send message back to client
			clientInRoom.Send <- newMessageMarshal
		}
	}
}
