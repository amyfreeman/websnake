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
	g.notifyOne(0, "GAMESTATE", g.snake.GetStateString())
	g.notifyOne(1, "GAMESTATE", g.snake.GetInvertedStateString())
	t := time.Now()
	for !g.gameover{
		if time.Since(t) > 1000000000 {
			t = time.Now()
			g.snake.Step()
			g.notifyOne(0, "GAMESTATE", g.snake.GetStateString())
			g.notifyOne(1, "GAMESTATE", g.snake.GetInvertedStateString())
		}
		g.gameover = g.snake.Gameover()
	}
	g.notifyAll("STATUS", "GAMEOVER")
	fmt.Println("Game over")
}

func (g *Game) notifyAll(channel string, msg string){
	for _, player := range (*g).players {
		player.Channel(channel).Write(msg)
	}
}

func (g *Game) notifyOne(playerIndex int, channel string, msg string){
	(*g).players[playerIndex].Channel(channel).Write(msg)
}

func (g *Game) keyPress(p *Player, dir int){
	for i := 0; i < len(g.players); i++{
		if g.players[i] == p {
			if i == 0{
				g.snake.Move(i, dir)
			} else{
				g.snake.Move(i, (dir + 2) % 4)
			}

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
	time.Sleep(4 * time.Second)
	g.notifyAll("STATUS", "BEGIN")
	go gameListener(&g)
	return &g
}