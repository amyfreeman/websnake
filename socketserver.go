package main

import (
    "net/http"
	"github.com/desertbit/glue"
    "github.com/nu7hatch/gouuid"
	"fmt"
)

type SocketServer struct {
    port string
    players map[string] *Player
    games map[*Player] *Game
    nextStranger *Player
}

func (ss *SocketServer) Listen() {
    //make this a link
    fmt.Println("SocketServer listening at localhost" + ss.port)

	http.Handle("/", http.FileServer(http.Dir("public/dist")))
	server := glue.NewServer(glue.Options{
		HTTPListenAddress: ss.port,
	})
    defer server.Release()
    server.OnNewSocket(ss.OnNewSocket)
    err := server.Run()
	if err != nil{
        //find a better way to handler errors
		fmt.Println(err)
    }
}

func (ss *SocketServer) OnNewSocket(s *glue.Socket) {
    s.OnRead(func(data string){
        p := ss.players[s.ID()]

        // convert code below to use channels. extract switch statement to game class, simply passing the message along

        switch data{
        case "READY":
            fmt.Println(s.ID() + " is ready to start")
            if (ss.nextStranger == nil){
                ss.nextStranger = p
            } else{
                var p2 *Player = ss.nextStranger
                ss.nextStranger = nil
                ss.makeGame(p2, p)
            }
        case "LEFT":
            ss.games[p].keyPress(p, 2)
        case "RIGHT":
            ss.games[p].keyPress(p, 0)
        case "UP":
            ss.games[p].keyPress(p, 1)
        case "DOWN":
            ss.games[p].keyPress(p, 3)
        }
    })
    
    s.OnClose(func(){
        fmt.Println("socket closed with remote address:", s.RemoteAddr())
    })

    p := createPlayer(s)
    ss.players[s.ID()] = p

    fmt.Println("socket open with remote address:", s.RemoteAddr())

    //write on GAMESTATE channel
    s.Write("0................F.................................................................................1")
}

func (ss *SocketServer) makeGame(p1 *Player, p2 *Player) {
    fmt.Println("Making New Game")
    u4, err := uuid.NewV4()
    if err != nil {
        // better error handling
        fmt.Println("error:", err)
    }
    gameId := u4.String()
    game := createGame(gameId, p1, p2)
    ss.games[p1] = game
    ss.games[p2] = game
}

func createSocketServer(port string) {
	ss := SocketServer{
        port: port,
        players: make(map[string] *Player),
        games: make(map[*Player] *Game),
        nextStranger: nil,
    }
	ss.Listen()
}