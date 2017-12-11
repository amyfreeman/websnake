package main

import (
	"github.com/desertbit/glue"
	"github.com/cdalizadeh/websnake/snake"
	"fmt"
	"time"
)

type Game struct {
    gameId string
	ch chan string
	players []*Player
	snake *snake.Snake
	gameover bool
}

func gameListener(g *Game) {
	fmt.Println("Game listening")
	str := g.snake.GetStateString();
	fmt.Println(str)
	for _, player := range (*g).players {
		player.socket.Write(str)
	}
	t := time.Now()
	for !g.gameover{
		select {
		case msg, ok := <-(*g).ch:
			if ok {
				for _, player := range (*g).players {
					player.socket.Write(msg)
					fmt.Println(msg)
				}
			} else {
				fmt.Println("Channel closed!")
			}
		default:
		}
		if time.Since(t) > 1000000000 {
			t = time.Now()
			g.snake.Step()
			g.snake.PrintState()
			str = g.snake.GetStateString()
			for _, player := range (*g).players {
				player.socket.Write(str)
			}
		}
	}
}

func (g *Game) leftPress(socketId string){
	for i := 0; i < len(g.players); i++{
		if g.players[i].socket.ID() == socketId {
			g.snake.Move(i, 2)
		}
	}
}

func (g *Game) rightPress(socketId string){
	for i := 0; i < len(g.players); i++{
		if g.players[i].socket.ID() == socketId {
			g.snake.Move(i, 0)
		}
	}
}

func (g *Game) upPress(socketId string){
	for i := 0; i < len(g.players); i++{
		if g.players[i].socket.ID() == socketId {
			g.snake.Move(i, 1)
		}
	}
}

func (g *Game) downPress(socketId string){
	for i := 0; i < len(g.players); i++{
		if g.players[i].socket.ID() == socketId {
			g.snake.Move(i, 3)
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
	g.gameover = false
	g.snake = snake.CreateSnake()
	go gameListener(&g)
	return &g
}