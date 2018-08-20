package main

import (
	"github.com/desertbit/glue"
)

// Player represents a snake game player. It is composed of a glue socket, and contains a nickname
// and reference to the current game the player is in.
type Player struct {
	*glue.Socket
	Nickname string
	Game     *Game
}

func newPlayer(socket *glue.Socket) *Player {
	p := Player{
		socket,
		"Nickname",
		nil,
	}
	return &p
}
