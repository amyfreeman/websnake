package main

import (
	"fmt"
	"os"

	"github.com/cdalizadeh/websnake/snake"
)

var port = ":8069"

func main() {
	Listen()
}

func testGame() {
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
