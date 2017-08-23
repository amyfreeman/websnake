package main

import (
	"github.com/desertbit/glue"
	"fmt"
)

type Game struct {
    gameId string
	ch chan string
	p1 *glue.Socket
	p2 *glue.Socket
}

func (g *Game) Listen() {
	fmt.Println("game " + g.gameId + " now listening")
	var msg string = ""
	for true {
		msg = <- g.ch
		if msg != "" {
			fmt.Println("message received: " + msg)
		}
	}
}


func createGame(gameId string, ch chan string) Game {
	g := Game{}
	g.gameId = gameId
	g.ch = ch
	go g.Listen()
	return g
}