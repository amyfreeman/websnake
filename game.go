package main

import (
	"github.com/desertbit/glue"
	"fmt"
)

type Game struct {
    gameId string
	ch chan string
	players []*glue.Socket
}

func (g *Game) Listen() {
	fmt.Println("game " + g.gameId + " now listening")
	var msg string = ""
	for true {
		msg = <- g.ch
		if msg != "" {
			fmt.Println("message received: " + msg)
			g.players[0].Write("pls receive")
		}
	}
}


func createGame(gameId string, ch chan string, p1 *glue.Socket) Game {
	g := Game{}
	g.gameId = gameId
	g.ch = ch
	g.players = make([]*glue.Socket, 1, 4)
	g.players[0] = p1
	go g.Listen()
	return g
}