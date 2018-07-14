package main

import (
    "net/http"
	"fmt"
	"os"
	"github.com/cdalizadeh/websnake/snake"
)

var PORT1 = ":8039"
var PORT2 = ":8040"

func main() {
	go createSocketServer(PORT2)

	fmt.Println("Now Serving.")
	http.ListenAndServe(PORT1, http.FileServer(http.Dir("./public/dist")))
}

func testGame(){
	fmt.Println(string(3))
	sn := snake.CreateSnake()
	sn.Step()
	sn.Step()
	sn.Move(1, 0)
	sn.Step()
	sn.Step()
	sn.Move(1, 3)
	sn.Step()
	sn.Move(1, 2)
	sn.Step()
	sn.Step()
	sn.Step()
	sn.Step()
	sn.PrintState()
	os.Exit(3)
}
