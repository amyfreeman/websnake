package main

import (
	"github.com/desertbit/glue"
    "github.com/nu7hatch/gouuid"
	"fmt"
)

type SocketServer struct {
	port string
    gameIds map[string]string
    games map[string] *Game
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
    s.Channel("create").OnRead(func(data string){
        ch := make(chan string)
        u4, err := uuid.NewV4()
        if err != nil {
            fmt.Println("error:", err)
        }
        gameId := u4.String()
        ss.gameIds[s.ID()] = gameId
        fmt.Println(gameId)
        game := createGame(gameId, ch, s)
        ss.games[gameId] = game
    })
    s.Channel("join").OnRead(func(gameId string){
        fmt.Println("Player joining")
        game := ss.games[gameId]
        (*game).addPlayer(s)
        (*game).ch <- "new player"
    })
    s.OnClose(func(){
        fmt.Println(s.ID())
    })
	fmt.Println("socket open with remote address:", s.RemoteAddr())
}

func createSocketServer(port string) {
	ss := SocketServer{}
	ss.port = port
    ss.gameIds = make(map[string]string)
    ss.games = make(map[string] *Game)
	ss.Listen()
}