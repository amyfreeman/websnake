package main

import (
    "log"
    "net/http"
	//"fmt"
    "github.com/desertbit/glue"
)

func main() {
    server := glue.NewServer(glue.Options{
        HTTPListenAddress: ":8039",
    })
    defer server.Release()
    server.OnNewSocket(onNewSocket)
    go server.Run()
    http.HandleFunc("/", handler)
	http.HandleFunc("/stranger", handler2)
	http.HandleFunc("/friend", handler3)
	http.HandleFunc("/js/glue.js", handler4)
    http.ListenAndServe(":8039", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "public/index.html")
}

func handler2(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "public/stranger.html")
}

func handler3(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "public/friend.html")
}

func handler4(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "public/js/glue.js")
}

func onNewSocket(s *glue.Socket) {
    // Set a function which is triggered as soon as the socket is closed.
    s.OnClose(func() {
        log.Printf("socket closed with remote address: %s", s.RemoteAddr())
    })

    // Set a function which is triggered during each received message.
    s.OnRead(func(data string) {
        // Echo the received data back to the client.
        s.Write(data)
    })

    // Send a welcome string to the client.
	log.Printf("socket open with remote address: %s", s.RemoteAddr())
    s.Write("Hello Client")
}