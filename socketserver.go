package main

import (
	"fmt"
	"net/http"

	"github.com/desertbit/glue"
	"github.com/nu7hatch/gouuid"
	//"reflect"
)

type SocketServer struct {
	port         string
	players      map[string]*Player // Maps (*glue.Socket).ID() to *Player
	games        map[*Player]*Game  // Maps *Player to *Game
	nextStranger *Player
}

func (ss *SocketServer) Listen() {
	fmt.Println("SocketServer listening at http://localhost" + ss.port)

	http.Handle("/", http.FileServer(http.Dir("public/dist")))
	server := glue.NewServer(glue.Options{
		HTTPListenAddress: ss.port,
	})
	defer server.Release()
	server.OnNewSocket(ss.OnNewSocket)
	err := server.Run()
	if err != nil {
		//find a better way to handle errors
		fmt.Println(err)
	}
}

func (ss *SocketServer) OnNewSocket(s *glue.Socket) {
	s.OnRead(func(data string) {
		fmt.Println("Unknown socket command detected: " + data)
	})

	s.Channel("STATUS").OnRead(func(data string) {
		p := ss.players[s.ID()]

		switch data {
		case "READY":
			fmt.Println(p.Nickname + " is ready to start")
			if ss.nextStranger == nil {
				ss.nextStranger = p
			} else {
				var p2 = ss.nextStranger
				ss.nextStranger = nil
				ss.makeGame(p2, p)
			}
		default:
			fmt.Println("Unknown STATUS command detected: " + data)
		}
	})

	s.Channel("NICKNAME").OnRead(func(data string) {
		ss.players[s.ID()].Nickname = data
	})

	s.Channel("GAMEPLAY").OnRead(func(data string) {
		p := ss.players[s.ID()]
		game := ss.games[p]
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

	s.OnClose(func() {
		fmt.Println("socket closed with remote address:", s.RemoteAddr())
	})

	p := createPlayer(s)
	ss.players[s.ID()] = p

	// run on a server other than localhost, find out if this works
	fmt.Println("socket open with remote address:", s.RemoteAddr())
}

func (ss *SocketServer) makeGame(p1 *Player, p2 *Player) {
	fmt.Println("Making New Game")
	u4, err := uuid.NewV4()
	if err != nil {
		// better error handling
		fmt.Println("error:", err)
	}
	gameID := u4.String()
	game := createGame(gameID, p1, p2)
	ss.games[p1] = game
	ss.games[p2] = game
}

func createSocketServer(port string) {
	ss := SocketServer{
		port:         port,
		players:      make(map[string]*Player),
		games:        make(map[*Player]*Game),
		nextStranger: nil,
	}
	ss.Listen()
}
