//For now, player is ENTIRELY defined by current glue socket.

package main

import (
	"github.com/desertbit/glue"
)

type Player struct {
	socket *glue.Socket
	name string
}

func createPlayer(socket *glue.Socket) *Player {
	p := Player{}
	p.socket = socket;
	return &p
}