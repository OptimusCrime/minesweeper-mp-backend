package main

////////
/// General types
/////////

type MessageRoom struct {
	RoomCode    string `json:"roomCode"`
	CurrentTurn int    `json:"currentTurn"`
	HasStarted  bool   `json:"hasStarted"`
}

type MessagePlayer struct {
	Nickname      string `json:"Nickname"`
	CurrentPlayer bool   `json:"currentPlayer"`
	Points        int    `json:"points"`
	HasWon        bool   `json:"HasWon"`
}

////////
/// Lobby join
/////////

type JoinLobbyIncomingMessage struct {
	Type     string `json:"type"`
	Nickname string `json:"Nickname"`
	RoomCode string `json:"roomCode"`
}

type LobbySuccessJoinOutboundMessage struct {
	Type     string          `json:"type"`
	Nickname string          `json:"Nickname"`
	Room     MessageRoom     `json:"room"`
	Players  []MessagePlayer `json:"players"`
}

type LobbyFailedJoinOutboundMessage struct {
	Type     string `json:"type"`
	RoomCode string `json:"roomCode"`
}
