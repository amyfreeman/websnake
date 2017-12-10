package main

import (
	"github.com/desertbit/glue"
	"github.com/cdalizadeh/websnake/snake"
	"fmt"
)

type Game struct {
    gameId string
	ch chan string
	players []*Player
	snake *snake.Snake
}

func gameListener(g *Game) {
	fmt.Println("Game listening")
	//refactor for time based looping
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
		if i % 100000000 == 0 { //refactor for time-based looping
			for _, player := range (*g).players {
				player.socket.Write("Game state changed")
			}
		}
	}
}

func createGame(gameId string, ch chan string, s1 *glue.Socket, s2 *glue.Socket) *Game {
	g := Game{}
	g.gameId = gameId
	g.ch = ch
	g.players = make([]*Player, 2, 4)
	g.players[0] = createPlayer(s1)
	g.players[1] = createPlayer(s2)
	g.snake = snake.CreateSnake()
	g.snake.Step()
	go gameListener(&g)
	return &g
}