package main

import (
	"github.com/desertbit/glue"
	"fmt"
)

type SocketServer struct {
	port string
    gameIds map[string]string
    games map[string]Game
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
    s.OnRead(func(data string) {
        fmt.Println("socket read")
        if len(data) > 8 && data[:8] == "[create]" {
            ch := make(chan string)
            ss.gameIds[s.ID()] = data[8:]
            ss.games[data[8:]] = createGame(data[8:], ch, s)
        } else {
            fmt.Println("sending to channel")
            ss.games[ss.gameIds[s.ID()]].ch <- data
        }
    })
    s.OnClose(func() {
        fmt.Println("socket closed with remote address:", s.RemoteAddr())
    })
	fmt.Println("socket open with remote address:", s.RemoteAddr())
}

func createSocketServer(port string) {
	ss := SocketServer{}
	ss.port = port
    ss.gameIds = make(map[string]string)
    ss.games = make(map[string]Game)
	ss.Listen()
}