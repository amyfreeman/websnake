package main

import (
	"github.com/cdalizadeh/websnake/snake"
	"fmt"
	"time"
)

type Game struct {
    gameId string
	players []*Player
	snake *snake.Snake
	gameover bool
}

func gameListener(g *Game) {
	fmt.Println("Game listening")
	g.notifyAll(g.snake.GetStateString())
	t := time.Now()
	for !g.gameover{
		if time.Since(t) > 1000000000 {
			t = time.Now()
			g.snake.Step()
			g.notifyAll(g.snake.GetStateString())
		}
	}
}

func (g *Game) notifyAll(msg string){
	for _, player := range (*g).players {
		player.socket.Write(msg)
	}
}

func (g *Game) keyPress(p *Player, dir int){
	for i := 0; i < len(g.players); i++{
		if g.players[i] == p {
			g.snake.Move(i, dir)
		}
	}
}

func createGame(gameId string, p1 *Player, p2 *Player) *Game {
	g := Game{}
	g.gameId = gameId
	g.players = make([]*Player, 2, 4)
	g.players[0] = p1
	g.players[1] = p2
	g.gameover = false
	g.snake = snake.CreateSnake()
	g.notifyAll("game beginning")
	go gameListener(&g)
	return &g
}