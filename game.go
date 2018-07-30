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
	g.notifyAll("GAMEPLAY", g.snake.GetStateString())
	t := time.Now()
	for !g.gameover{
		if time.Since(t) > 1000000000 {
			t = time.Now()
			g.snake.Step()
			g.notifyAll("GAMEPLAY", g.snake.GetStateString())
		}
	}
	g.notifyAll("STATUS", "GAMEOVER")
}

func (g *Game) notifyAll(channel string, msg string){
	for _, player := range (*g).players {
		// write to channel
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
	g := Game{
		gameId: gameId,
		players: make([]*Player, 2, 4),
		gameover: false,
		snake: snake.CreateSnake(),
	}
	// todo: add the two lines below to initializer
	g.players[0] = p1
	g.players[1] = p2
	g.notifyAll("STATUS", "OPPONENT_FOUND")
	// sleep for three seconds
	g.notifyAll("STATUS", "BEGIN")
	go gameListener(&g)
	return &g
}