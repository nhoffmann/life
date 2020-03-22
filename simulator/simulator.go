package simulator

import "github.com/nhoffmann/life/grid"

type Simulator struct {
	Rows    int
	Columns int

	grid grid.Grid
}

func New(rows, columns int) *Simulator {
	return &Simulator{
		Rows:    rows,
		Columns: columns,
		grid:    grid.New(rows, columns),
	}
}

// Evolute creates the next state of the grid and then replaces the current one
func (s *Simulator) Evolute() {
	newGrid := grid.New(s.Rows, s.Columns)

	for rowIndex := 0; rowIndex < s.grid.Width(); rowIndex++ {
		for columnIndex := 0; columnIndex < s.grid.Height(); columnIndex++ {
			s.evoluteCell(newGrid, rowIndex, columnIndex)
		}
	}

	s.grid = newGrid
}

// LoadPattern populates the grid with a given pattern
func (s *Simulator) LoadPattern(pattern [][]int, startX, startY int) {
	for rowIndex, patternRow := range pattern {
		for colIndex := range patternRow {
			if pattern[rowIndex][colIndex] == 1 {
				s.Populate(rowIndex+startX, colIndex+startY)
			}
		}
	}
}

// Populate populates the cell at the given coordinates
func (s *Simulator) Populate(x, y int) {
	s.grid.Populate(x, y)
}

func (s *Simulator) evoluteCell(grid grid.Grid, rowIndex, columnIndex int) {
	if s.cellLives(rowIndex, columnIndex) {
		grid.Populate(rowIndex, columnIndex)
	}
}

func (s *Simulator) cellLives(rowIndex, columnIndex int) bool {
	cellPopulated := s.grid.IsPopulated(rowIndex, columnIndex)

	neighbors := []bool{
		s.grid.IsPopulated(rowIndex-1, columnIndex-1), // topLeft
		s.grid.IsPopulated(rowIndex-1, columnIndex),   // topCenter
		s.grid.IsPopulated(rowIndex-1, columnIndex+1), // topRight
		s.grid.IsPopulated(rowIndex, columnIndex-1),   // left
		s.grid.IsPopulated(rowIndex, columnIndex+1),   // right
		s.grid.IsPopulated(rowIndex+1, columnIndex-1), // bottomLeft
		s.grid.IsPopulated(rowIndex+1, columnIndex),   // bottomCenter
		s.grid.IsPopulated(rowIndex+1, columnIndex+1), // bottomRight
	}

	var aliveNeighborCount int

	for _, neighborPopulated := range neighbors {
		if neighborPopulated {
			aliveNeighborCount++
		}
	}

	if cellPopulated {
		// Any live cell with two or three neighbors survives.
		if aliveNeighborCount == 2 || aliveNeighborCount == 3 {
			return true
		}
	} else {
		// Any dead cell with three live neighbors becomes a live cell.
		if aliveNeighborCount == 3 {
			return true
		}
	}

	// All other cells die in the next generation.
	return false
}

func (s *Simulator) String() string {
	return s.grid.String()
}
