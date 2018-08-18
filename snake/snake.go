package snake

import (
	"fmt"
	"math/rand"
)

var GAME_WIDTH = 10
var GAME_HEIGHT = 10
var NUM_SNAKES = 2
var NUM_FOODS = 1

type Snake struct {
	width       int
	height      int
	bodies      []*Body
	foods       []*Food
	gameover    bool
	isDead      []bool
	currentTurn int
}

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

	sn.currentTurn = (sn.currentTurn + 1) % NUM_SNAKES
	for sn.isDead[sn.currentTurn] {
		sn.currentTurn = (sn.currentTurn + 1) % NUM_SNAKES
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

func (sn *Snake) Gameover() bool {
	return sn.gameover
}

func (sn *Snake) GetStateString() string {
	var str string = ""
	for i := 0; i < GAME_WIDTH; i++ {
		for j := 0; j < GAME_HEIGHT; j++ {
			str += sn.getObjectAt(Cell{i, j})
		}
	}
	return str
}

func (sn *Snake) GetInvertedStateString() string {
	var str string = ""
	for i := GAME_WIDTH - 1; i >= 0; i-- {
		for j := GAME_HEIGHT - 1; j >= 0; j-- {
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

func (sn *Snake) PrintState() {
	for j := GAME_HEIGHT - 1; j >= 0; j-- {
		for i := 0; i < GAME_WIDTH; i++ {
			fmt.Print(sn.getObjectAt(Cell{i, j}) + " ")
		}
		fmt.Println()
	}
}

func CreateSnake() *Snake {
	sn := Snake{
		width:       GAME_WIDTH,
		height:      GAME_HEIGHT,
		bodies:      make([]*Body, NUM_SNAKES, NUM_SNAKES),
		foods:       make([]*Food, 0, NUM_FOODS),
		isDead:      make([]bool, NUM_SNAKES, NUM_SNAKES),
		currentTurn: 0,
	}
	if NUM_SNAKES == 2 {
		sn.bodies[0] = createBody(0, 0, 0, GAME_WIDTH*GAME_HEIGHT)
		sn.bodies[1] = createBody(GAME_WIDTH-1, GAME_HEIGHT-1, 2, GAME_WIDTH*GAME_HEIGHT)
	} else if NUM_SNAKES == 4 {
		sn.bodies[0] = createBody(0, 0, 0, GAME_WIDTH*GAME_HEIGHT)
		sn.bodies[1] = createBody(0, GAME_HEIGHT-1, 3, GAME_WIDTH*GAME_HEIGHT)
		sn.bodies[2] = createBody(GAME_WIDTH-1, GAME_HEIGHT-1, 2, GAME_WIDTH*GAME_HEIGHT)
		sn.bodies[3] = createBody(GAME_WIDTH-1, 0, 1, GAME_WIDTH*GAME_HEIGHT)
	}
	for i := 0; i < NUM_FOODS; i++ {
		sn.foods = append(sn.foods, createFood(sn.getUnoccupiedCell()))
	}
	return &sn
}
