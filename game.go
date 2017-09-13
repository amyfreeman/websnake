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
	fmt.Println("Game listening")
	for i := 0; ; i++  {

		select {
		case msg, ok := <-(*g).ch:
			if ok {
				for _, player := range (*g).players {
					player.socket.Write(msg)
				}
			} else {
				fmt.Println("Channel closed!")
			}
		default:
		}
		if i % 100000000 == 0 {
			for _, player := range (*g).players {
				player.socket.Write("Game state changed")
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
	go gameListener(&g)
	return &g
}