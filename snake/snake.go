package snake

import (
	"math/rand"
	"fmt"
)

var GAME_WIDTH = 10
var GAME_HEIGHT = 10

type Snake struct {
	width int
	height int
	bodies []*Body
	foods []*Food
	gameover bool
}

func (sn *Snake) Step(){
	for i:= 0; i < len(sn.bodies); i++ {
		if sn.bodies[i].step(){
			sn.gameover = true
		}
	}
	//implement eating
}

func (sn *Snake) Move(player int, dir int){
	sn.bodies[player].setDir(dir)
}

func (sn *Snake) getUnoccupiedCell() Cell{
	x := rand.Intn(sn.width)
	y := rand.Intn(sn.height)
	cell := Cell{x, y}
	for sn.isOccupied(cell) {
		x = rand.Intn(sn.width)
		y = rand.Intn(sn.height)
		cell = Cell{x, y}
	}
	return cell
}

func (sn *Snake) isOccupied(cell Cell) bool{
	if sn.getObjectAt(cell) == "0"{
		return false
	}
	return true
}

func (sn *Snake) getObjectAt(cell Cell) string{
	for i := 0; i < len(sn.bodies); i++ {
		if sn.bodies[i].contains(cell){
			return ("X")
			//return string(i + 1)
		}
	}
	for i := 0; i < len(sn.foods); i++ {
		if sn.foods[i].contains(cell){
			return "F"
		}
	}
	return "0"
}

func (sn *Snake) PrintState(){
	for j := GAME_HEIGHT - 1; j >= 0; j--{
		for i := 0; i < GAME_WIDTH; i++{
			fmt.Print(sn.getObjectAt(Cell{i, j}) + " ")
		}
		fmt.Println()
	}
}

func CreateSnake() *Snake{
	sn := Snake{}
	sn.width = GAME_WIDTH
	sn.height = GAME_HEIGHT
	sn.bodies = make([]*Body, 2, 4)
	sn.bodies[0] = createBody(0, 0, 0, GAME_WIDTH * GAME_HEIGHT)
	sn.bodies[1] = createBody(GAME_WIDTH - 1, GAME_HEIGHT - 1, 2, GAME_WIDTH * GAME_HEIGHT)
	sn.foods = make([]*Food, 0, 4)
	sn.foods = append(sn.foods, createFood(sn.getUnoccupiedCell()))
	return &sn
}