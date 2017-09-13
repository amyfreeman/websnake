package main

import (
	"github.com/desertbit/glue"
)

type Player struct {
	socket *glue.Socket
}

func createPlayer(socket *glue.Socket) *Player {
	p := Player{}
	p.socket = socket;
	return &p
}