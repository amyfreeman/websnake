package snake

import (
	"fmt"
	"math/rand"
)

var gameWidth = 10
var gameHeight = 10
var numSnakes = 2
var numFoods = 1

// Snake holds the internal state of the snake game. It holds the game width, game height, list of snake bodies, list of food,
// gameover flag, and index of the snake which has the current turn.
type Snake struct {
	width       int
	height      int
	bodies      []*Body
	foods       []*Food
	gameover    bool
	isDead      []bool
	currentTurn int
}

// Step is the function which is called to step the game forward one move. A single snake moves one
// square forward on a step call.
func (sn *Snake) Step() {
	sn.bodies[sn.currentTurn].step()
	sn.isDead[sn.currentTurn] = sn.legalCheck(sn.bodies[sn.currentTurn])

	if sn.isDead[sn.currentTurn] {
		sn.gameover = true
	} else {
		for i, food := range sn.foods {
			if food.contains(sn.bodies[sn.currentTurn].head) {
				sn.bodies[sn.currentTurn].grow()
				sn.foods[i] = createFood(sn.getUnoccupiedCell())
			}
		}
	}

	sn.currentTurn = (sn.currentTurn + 1) % numSnakes
	for sn.isDead[sn.currentTurn] {
		sn.currentTurn = (sn.currentTurn + 1) % numSnakes
	}
}

func (sn *Snake) legalCheck(b *Body) bool {
	if b.legalCheck() {
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
	if !sn.isDead[player] {
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
		if food.contains(cell) {
			return "F"
		}
	}
	return "."
}

// Gameover returns the gameover state
func (sn *Snake) Gameover() bool {
	return sn.gameover
}

// GetStateString returns the current display state of the game encoded as a string, from player 1's perspective
func (sn *Snake) GetStateString() string {
	var str string
	for i := 0; i < gameWidth; i++ {
		for j := 0; j < gameHeight; j++ {
			str += sn.getObjectAt(Cell{i, j})
		}
	}
	return str
}

// GetInvertedStateString returns the current display state of the game encoded as a string, from player 2's perspective
func (sn *Snake) GetInvertedStateString() string {
	var str string
	for i := gameWidth - 1; i >= 0; i-- {
		for j := gameHeight - 1; j >= 0; j-- {
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
	for j := gameHeight - 1; j >= 0; j-- {
		for i := 0; i < gameWidth; i++ {
			fmt.Print(sn.getObjectAt(Cell{i, j}) + " ")
		}
		fmt.Println()
	}
}

// New returns a new game of snake.
func New() *Snake {
	sn := Snake{
		width:       gameWidth,
		height:      gameHeight,
		bodies:      make([]*Body, numSnakes, numSnakes),
		foods:       make([]*Food, 0, numFoods),
		isDead:      make([]bool, numSnakes, numSnakes),
		currentTurn: 0,
	}
	if numSnakes == 2 {
		sn.bodies[0] = createBody(0, 0, 0, gameWidth*gameHeight)
		sn.bodies[1] = createBody(gameWidth-1, gameHeight-1, 2, gameWidth*gameHeight)
	} else if numSnakes == 4 {
		sn.bodies[0] = createBody(0, 0, 0, gameWidth*gameHeight)
		sn.bodies[1] = createBody(0, gameHeight-1, 3, gameWidth*gameHeight)
		sn.bodies[2] = createBody(gameWidth-1, gameHeight-1, 2, gameWidth*gameHeight)
		sn.bodies[3] = createBody(gameWidth-1, 0, 1, gameWidth*gameHeight)
	}
	for i := 0; i < numFoods; i++ {
		sn.foods = append(sn.foods, createFood(sn.getUnoccupiedCell()))
	}
	return &sn
}
