package grid

type Generation struct {
	CellMap map[Cell]bool
}

func NewGeneration() *Generation {
	return &Generation{
		CellMap: make(map[Cell]bool),
	}
}

func (b *Generation) LoadPattern(pattern [][]int) {
	startX := len(pattern[0]) / -2
	startY := len(pattern) / -2

	for rowIndex, patternRow := range pattern {
		for colIndex := range patternRow {
			if pattern[rowIndex][colIndex] > 0 {
				b.Populate(NewCell(colIndex+startX, rowIndex+startY))
			}
		}
	}
}

func (b *Generation) CellCount() int {
	return len(b.CellMap)
}

func (b *Generation) Populate(cell Cell) {
	b.CellMap[cell] = true
}

func (b *Generation) IsPopulated(cell Cell) bool {
	_, present := b.CellMap[cell]

	return present
}

func (b *Generation) PopulatedNeighborsCount(cell Cell) int {
	var count int
	for _, neighboringCell := range cell.MooreNeighborhood() {
		if b.IsPopulated(neighboringCell) {
			count++
		}
	}
	return count
}
