package main

import (
	"github.com/desertbit/glue"
    "github.com/nu7hatch/gouuid"
	"fmt"
)

type SocketServer struct {
	port string
    gameIds map[string]string //mapping from socket ID to game ID
    games map[string] *Game //mapping from game ID to game pointer
    nextStranger *glue.Socket
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
        if (data == "stranger"){
            fmt.Println("stranger")
            if (ss.nextStranger == nil){
                ss.nextStranger = s
                s.Write("confirmed. waiting for stranger")
            } else{
                var s2 *glue.Socket = ss.nextStranger
                ss.nextStranger = nil
                s2.Write("game beginning")
                s.Write("game beginning")
                ss.makeGame(s, s2)
            }
        } else if (data == "friend"){
            fmt.Println("friend")
        } else if (data == "left"){
            ss.games[ss.gameIds[s.ID()]].leftPress(s.ID())
        }
    })
    
    s.OnClose(func(){
        fmt.Println("socket closed with remote address:", s.RemoteAddr())
    })
	fmt.Println("socket open with remote address:", s.RemoteAddr())
}

func (ss *SocketServer) makeGame(s1 *glue.Socket, s2 *glue.Socket) {
    fmt.Println("Making Game")
    ch := make(chan string)
    u4, err := uuid.NewV4()
    if err != nil {
        fmt.Println("error:", err)
    }
    gameId := u4.String()
    ss.gameIds[s1.ID()] = gameId
    ss.gameIds[s2.ID()] = gameId
    game := createGame(gameId, ch, s1, s2)
    ss.games[gameId] = game
}

func createSocketServer(port string) {
	ss := SocketServer{}
	ss.port = port
    ss.gameIds = make(map[string]string)
    ss.games = make(map[string] *Game)
    ss.nextStranger = nil
	ss.Listen()
}