//For now, player is ENTIRELY defined by current glue socket.

package main

import (
	"github.com/desertbit/glue"
)

type Player struct {
	// compose with socket instead?
	socket *glue.Socket
	name string
}

func createPlayer(socket *glue.Socket) *Player {
	p := Player{
		socket: socket,
	}
	return &p
}