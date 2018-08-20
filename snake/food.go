package snake

// Food repesents a single piece of food.
type Food struct {
	cell Cell
}

func (f *Food) equals(cell Cell) bool {
	if cell.x == f.cell.x && cell.y == f.cell.y {
		return true
	}
	return false
}

func newFood(cell Cell) *Food {
	f := Food{cell}
	return &f
}
