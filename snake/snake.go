package snake

import (
	"fmt"
	"math/rand"
)

const (
	defaultGameWidth  = 10
	defaultGameHeight = 10
	defaultNumBodies  = 2
	defaultNumFoods   = 1
)

// Snake holds the internal state of the snake game. It holds the game width, game height, list of snake bodies, list of food,
// gameover flag, and index of the snake which has the current turn.
type Snake struct {
	width       int
	height      int
	bodies      []*Body
	numBodies   int
	foods       []*Food
	Gameover    bool
	currentTurn int
}

// Step is the function which is called to step the game forward one move. A single snake moves one
// square forward on a step call.
func (sn *Snake) Step() {
	currentBody := sn.bodies[sn.currentTurn]
	currentBody.step()
	currentBody.isDead = sn.legalCheck(currentBody)

	if currentBody.isDead {
		sn.Gameover = true
	} else {
		for i, food := range sn.foods {
			if food.equals(currentBody.head) {
				currentBody.grow()
				sn.foods[i] = newFood(sn.getUnoccupiedCell())
			}
		}
	}

	sn.currentTurn = (sn.currentTurn + 1) % sn.numBodies
	for currentBody.isDead {
		sn.currentTurn = (sn.currentTurn + 1) % sn.numBodies
		currentBody = sn.bodies[sn.currentTurn]
	}
}

func (sn *Snake) legalCheck(b *Body) bool {
	for i := 0; i < len(b.cells)-1; i++ {
		if b.head.x == b.cells[i].x && b.head.y == b.cells[i].y {
			return true
		}
	}
	if b.head.x < 0 {
		return true
	}
	if b.head.y < 0 {
		return true
	}
	if b.head.x > sn.width-1 {
		return true
	}
	if b.head.y > sn.height-1 {
		return true
	}

	for _, body := range sn.bodies {
		if body != b {
			for _, cell := range body.cells {
				if b.head.x == cell.x && b.head.y == cell.y {
					return true
				}
			}
		}
	}
	return false
}

// Move sets the next direction of the snake specified.
func (sn *Snake) Move(player int, dir int) {
	if !sn.bodies[player].isDead {
		sn.bodies[player].setDir(dir)
	}
}

func (sn *Snake) getUnoccupiedCell() Cell {
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

func (sn *Snake) isOccupied(cell Cell) bool {
	if sn.getObjectAt(cell) == "." {
		return false
	}
	return true
}

func (sn *Snake) getObjectAt(cell Cell) string {
	for i := range sn.bodies {
		if sn.bodies[i].contains(cell) {
			if i == 0 {
				return "0"
			} else if i == 1 {
				return "1"
			} else if i == 2 {
				return "2"
			} else if i == 3 {
				return "3"
			}
		}
	}
	for _, food := range sn.foods {
		if food.equals(cell) {
			return "F"
		}
	}
	return "."
}

// GetStateString returns the current display state of the game encoded as a string, from player 1's perspective
func (sn *Snake) GetStateString() string {
	var str string
	for i := 0; i < sn.width; i++ {
		for j := 0; j < sn.height; j++ {
			str += sn.getObjectAt(Cell{i, j})
		}
	}
	return str
}

// GetInvertedStateString returns the current display state of the game encoded as a string, from player 2's perspective
func (sn *Snake) GetInvertedStateString() string {
	var str string
	for i := sn.width - 1; i >= 0; i-- {
		for j := sn.height - 1; j >= 0; j-- {
			var o = sn.getObjectAt(Cell{i, j})

			if o == "0" {
				str += "1"
			} else if o == "1" {
				str += "0"
			} else {
				str += o
			}
		}
	}
	return str
}

// PrintState prints the current state of the game to the console in human-readable format
func (sn *Snake) PrintState() {
	for j := sn.height - 1; j >= 0; j-- {
		for i := 0; i < sn.width; i++ {
			fmt.Print(sn.getObjectAt(Cell{i, j}) + " ")
		}
		fmt.Println()
	}
}

// New returns a new game of snake.
func New() *Snake {
	sn := Snake{
		width:       defaultGameWidth,
		height:      defaultGameHeight,
		bodies:      make([]*Body, defaultNumBodies, defaultNumBodies),
		numBodies:   defaultNumBodies,
		foods:       make([]*Food, 0, defaultNumFoods),
		currentTurn: 0,
	}
	if defaultNumBodies == 2 {
		sn.bodies[0] = newBody(0, 0, 0, sn.width*sn.height)
		sn.bodies[1] = newBody(sn.width-1, sn.height-1, 2, sn.width*sn.height)
	} else if defaultNumBodies == 4 {
		sn.bodies[0] = newBody(0, 0, 0, sn.width*sn.height)
		sn.bodies[1] = newBody(0, sn.height-1, 3, sn.width*sn.height)
		sn.bodies[2] = newBody(sn.width-1, sn.height-1, 2, sn.width*sn.height)
		sn.bodies[3] = newBody(sn.width-1, 0, 1, sn.width*sn.height)
	}
	for i := 0; i < defaultNumFoods; i++ {
		sn.foods = append(sn.foods, newFood(sn.getUnoccupiedCell()))
	}
	return &sn
}
