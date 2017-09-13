package main

import (
	"github.com/desertbit/glue"
	"fmt"
)

type Game struct {
    gameId string
	ch chan string
	players []*Player
}

func gameListener(g *Game) {
	fmt.Printf("Now listening to game address: %p \n", g)
	var msg string = ""
	for true {
		msg = <- (*g).ch
		if msg != "" { 
			for _, player := range (*g).players {
				player.socket.Write("pls receive")
			}
		}
	}
}

func (g *Game) addPlayer(socket *glue.Socket) {
	g.players = append(g.players, createPlayer(socket))
}

func createGame(gameId string, ch chan string, socket *glue.Socket) *Game {
	g := Game{}
	g.gameId = gameId
	g.ch = ch
	g.players = make([]*Player, 1, 4)
	g.players[0] = createPlayer(socket)
	fmt.Printf("Real game address: %p \n", &g)
	go gameListener(&g)
	return &g
}