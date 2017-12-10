package snake

type Food struct {
    cell Cell
}

func (f *Food) contains(cell Cell) bool{
	if cell.x == f.cell.x && cell.y == f.cell.y{
		return true
	}
	return false
}

func createFood(cell Cell) *Food{
	f := Food{cell}
	return &f
}