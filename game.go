package main

import (
	"fmt"
)

type Game struct {
    gameId string
}

func (g *Game) Listen() {
	fmt.Println("game now listening")
}


func createGame(gameId string) {
	g := Game{}
	g.gameId = gameId
	g.Listen()
}