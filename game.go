package main

import (
	"fmt"
	"time"

	"github.com/cdalizadeh/websnake/snake"
	"github.com/nu7hatch/gouuid"
)

// Game is the interface between the server and the internal game mechanics. Each game has an ID, list of players, reference to a
// snake game, and gameover flag.
type Game struct {
	gameID   string
	players  []*Player
	snake    *snake.Snake
	gameover bool
}

func gameListener(g *Game) {
	fmt.Println("Game listening")
	g.notifyOne(0, "GAMESTATE", g.snake.GetStateString())
	g.notifyOne(1, "GAMESTATE", g.snake.GetInvertedStateString())
	t := time.Now()
	var counter = 0
	for !g.gameover {
		if time.Since(t) > 100000000 {
			t = time.Now()
			g.snake.Step()
			g.notifyOne(0, "GAMESTATE", g.snake.GetStateString())
			g.notifyOne(1, "GAMESTATE", g.snake.GetInvertedStateString())
		}
		g.gameover = g.snake.Gameover
		counter++
		fmt.Println(counter)
	}
	g.notifyAll("STATUS", "GAMEOVER")
	fmt.Println("Game over")
}

func (g *Game) notifyAll(channel string, msg string) {
	for _, player := range g.players {
		player.Channel(channel).Write(msg)
	}
}

func (g *Game) notifyOne(playerIndex int, channel string, msg string) {
	(*g).players[playerIndex].Channel(channel).Write(msg)
}

func (g *Game) keyPress(p *Player, dir int) {
	for i := range g.players {
		if g.players[i] == p {
			if i == 0 {
				g.snake.Move(i, dir)
			} else {
				g.snake.Move(i, (dir+2)%4)
			}

		}
	}
}

func newGame(p1 *Player, p2 *Player) *Game {
	fmt.Println("Making New Game")
	u4, err := uuid.NewV4()
	if err != nil {
		// better error handling
		fmt.Println("error:", err)
	}

	g := Game{
		gameID:   u4.String(),
		players:  make([]*Player, 2, 4),
		gameover: false,
		snake:    snake.New(),
	}
	// todo: add the two lines below to initializer
	g.players[0] = p1
	g.players[1] = p2
	p1.Game = &g
	p2.Game = &g
	g.notifyAll("STATUS", "OPPONENT_FOUND")
	time.Sleep(4 * time.Second)
	g.notifyAll("STATUS", "BEGIN")
	go gameListener(&g)
	return &g
}
