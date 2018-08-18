//For now, player is ENTIRELY defined by current glue socket.

package main

import (
	"github.com/desertbit/glue"
)

type Player struct {
	*glue.Socket
	Nickname string
}

func createPlayer(socket *glue.Socket) *Player {
	p := Player{
		socket,
		"Nickname",
	}
	return &p
}
