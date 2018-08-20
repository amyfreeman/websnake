package main

import (
	"fmt"
	"net/http"

	"github.com/desertbit/glue"
)

var nextStranger *Player

func listen() {
	fmt.Println("Server listening at http://localhost" + port)

	http.Handle("/", http.FileServer(http.Dir("public/dist")))
	server := glue.NewServer(glue.Options{
		HTTPListenAddress: port,
	})
	defer server.Release()
	server.OnNewSocket(onNewSocket)
	err := server.Run()
	if err != nil {
		//find a better way to handle errors
		fmt.Println(err)
	}
}

func onNewSocket(s *glue.Socket) {
	p := newPlayer(s)

	p.OnRead(func(data string) {
		fmt.Println("Unknown socket command detected: " + data)
	})

	p.Channel("STATUS").OnRead(func(data string) {
		switch data {
		case "READY":
			fmt.Println(p.Nickname + " is ready to start")
			if nextStranger == nil {
				nextStranger = p
			} else {
				var p2 = nextStranger
				nextStranger = nil
				newGame(p2, p)
			}
		default:
			fmt.Println("Unknown STATUS command detected: " + data)
		}
	})

	p.Channel("NICKNAME").OnRead(func(data string) {
		p.Nickname = data
	})

	p.Channel("GAMEPLAY").OnRead(func(data string) {
		game := p.Game
		if game != nil {
			switch data {
			case "LEFT":
				game.keyPress(p, 2)
			case "RIGHT":
				game.keyPress(p, 0)
			case "UP":
				game.keyPress(p, 1)
			case "DOWN":
				game.keyPress(p, 3)
			default:
				fmt.Println("Unknown GAMEPLAY command detected: " + data)
			}
		}
	})

	p.OnClose(func() {
		if nextStranger == p {
			nextStranger = nil
		}
		fmt.Println("socket closed with remote address:", p.RemoteAddr())
	})

	fmt.Println("socket open with remote address:", p.RemoteAddr())
}
