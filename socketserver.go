package main

import (
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
	server := glue.NewServer(glue.Options{
        HTTPListenAddress: ss.port,
    })
    defer server.Release()
    server.OnNewSocket(ss.OnNewSocket)
	server.Run()
}

func (ss *SocketServer) OnNewSocket(s *glue.Socket) {
    s.OnRead(func(data string){
        p := ss.players[s.ID()]
        if (data == "stranger"){
            fmt.Println("stranger")
            if (ss.nextStranger == nil){
                ss.nextStranger = p
                s.Write("confirmed. waiting for stranger")
            } else{
                var p2 *Player = ss.nextStranger
                ss.nextStranger = nil
                ss.makeGame(p, p2)
            }
        } else if (data == "friend"){
            fmt.Println("friend")
        } else if (data == "left"){
            ss.games[p].leftPress(s.ID())
        } else if (data == "right"){
            ss.games[p].rightPress(s.ID())
        } else if (data == "up"){
            ss.games[p].upPress(s.ID())
        } else if (data == "down"){
            ss.games[p].downPress(s.ID())
        }
    })
    
    s.OnClose(func(){
        fmt.Println("socket closed with remote address:", s.RemoteAddr())
    })

    p := createPlayer(s)
    ss.players[s.ID()] = p

	fmt.Println("socket open with remote address:", s.RemoteAddr())
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