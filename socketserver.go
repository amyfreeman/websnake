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
	http.Handle("/", http.FileServer(http.Dir("public/dist")))
	server := glue.NewServer(glue.Options{
		HTTPListenAddress: ss.port,
	})
    defer server.Release()
    server.OnNewSocket(ss.OnNewSocket)
	err := server.Run()
	if err != nil{
		fmt.Println(err)
	}
}

func (ss *SocketServer) OnNewSocket(s *glue.Socket) {
    s.OnRead(func(data string){
        p := ss.players[s.ID()]

        switch data{
        case "stranger":
            fmt.Println("stranger")
            if (ss.nextStranger == nil){
                ss.nextStranger = p
                s.Write("confirmed. waiting for stranger")
            } else{
                var p2 *Player = ss.nextStranger
                ss.nextStranger = nil
                ss.makeGame(p2, p)
            }
        case "friend":
            fmt.Println("friend")
        case "left":
            ss.games[p].keyPress(p, 2)
        case "right":
            ss.games[p].keyPress(p, 0)
        case "up":
            ss.games[p].keyPress(p, 1)
        case "down":
            ss.games[p].keyPress(p, 3)
        }
    })
    
    s.OnClose(func(){
        fmt.Println("socket closed with remote address:", s.RemoteAddr())
    })

    p := createPlayer(s)
    ss.players[s.ID()] = p

    fmt.Println("socket open with remote address:", s.RemoteAddr())
    s.Write("0................F.................................................................................1")
}

func (ss *SocketServer) makeGame(p1 *Player, p2 *Player) {
    fmt.Println("Making Game")
    u4, err := uuid.NewV4()
    if err != nil {
        fmt.Println("error:", err)
    }
    gameId := u4.String()
    game := createGame(gameId, p1, p2)
    ss.games[p1] = game
    ss.games[p2] = game
}

func createSocketServer(port string) {
	ss := SocketServer{}
	ss.port = port
    ss.players = make(map[string] *Player)
    ss.games = make(map[*Player] *Game)
    ss.nextStranger = nil
	ss.Listen()
}