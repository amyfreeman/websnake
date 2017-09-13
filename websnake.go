package main

import (
    "net/http"
	"fmt"
)

var port1 string = ":8039"
var port2 string = ":8040"

func main() {
	go createSocketServer(port2)
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/stranger", handleStranger)
	http.HandleFunc("/friend", handleFriend)
	http.HandleFunc("/js/glue.js", handleGlue)
    http.ListenAndServe(port1, nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "public/index.html")
	default:
		// Give an error message.
	}
}

func handleStranger(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "public/stranger.html")
	case "POST":
		// Return game page
	default:
		// Give an error message.
	}
}

func handleFriend(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "public/friend.html")
	case "POST":
		fmt.Println("recieve")
	default:
		// Give an error message.
	}
}

func handleGlue(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "public/js/glue.js")
}