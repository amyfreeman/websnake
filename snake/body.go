package snake

import (
	"math"
)

// Body represents a single snake's body.
type Body struct {
	cells   []Cell
	head    Cell
	dir     int
	nextDir int
	addCell bool
	isDead  bool
}

func (b *Body) step() {
	b.dir = b.nextDir
	if b.addCell {
		b.addCell = false
	} else {
		b.cells = b.cells[1:]
	}
	if b.dir == 0 {
		b.head = Cell{b.head.x + 1, b.head.y}
	} else if b.dir == 1 {
		b.head = Cell{b.head.x, b.head.y + 1}
	} else if b.dir == 2 {
		b.head = Cell{b.head.x - 1, b.head.y}
	} else if b.dir == 3 {
		b.head = Cell{b.head.x, b.head.y - 1}
	}
	b.cells = append(b.cells, b.head)
}

func (b *Body) setDir(dir int) {
	if b.dir != dir && math.Abs(float64(b.dir-dir)) != 2 {
		b.nextDir = dir
	}
}

func (b *Body) grow() {
	b.addCell = true
}

func (b *Body) contains(cell Cell) bool {
	for i := range b.cells {
		if cell.x == b.cells[i].x && cell.y == b.cells[i].y {
			return true
		}
	}
	return false
}

func newBody(x int, y int, dir int, maxlength int) *Body {
	b := Body{
		cells:   make([]Cell, 1, maxlength),
		dir:     dir,
		nextDir: dir,
		addCell: false,
		isDead:  false,
	}
	b.cells[0] = Cell{x, y}
	b.head = b.cells[0]
	return &b
}
