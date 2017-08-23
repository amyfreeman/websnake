package main

import (
	"github.com/desertbit/glue"
	"fmt"
)

type SocketServer struct {
	port string
}

func (ss *SocketServer) Listen() {
	server := glue.NewServer(glue.Options{
        HTTPListenAddress: ":" + ss.port,
    })
    defer server.Release()
    server.OnNewSocket(ss.OnNewSocket)
	server.Run()
}

func (ss *SocketServer) OnNewSocket(s *glue.Socket) {
    s.OnClose(func() {
        fmt.Println("socket closed with remote address:", s.RemoteAddr())
    })

    s.OnRead(func(data string) {
        s.Write(data)
    })

	fmt.Println("socket open with remote address:", s.RemoteAddr())
    s.Write("Hello Client")
}

func createSocketServer(port string) {
	ss := SocketServer{}
	ss.port = port
	ss.Listen()
}