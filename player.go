package main

import (
	"github.com/desertbit/glue"
)

type Player struct {
	*glue.Socket
	Nickname string
	Game     *Game
}

func createPlayer(socket *glue.Socket) *Player {
	p := Player{
		socket,
		"Nickname",
		nil,
	}
	return &p
}
