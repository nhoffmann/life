package simulator

import (
	"bytes"
	"fmt"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/nhoffmann/life/grid"
)

type Simulator struct {
	Rows    int
	Columns int

	GenerationCount int

	grid *grid.Grid
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

	s.GenerationCount++
	s.grid = newGrid
}

// LoadPattern populates the grid with a given pattern
func (s *Simulator) LoadPattern(pattern [][]int) {
	patternWidth := len(pattern[0])
	patternHeigth := len(pattern)

	startY := (s.grid.Height() + patternHeigth) / 2
	startX := (s.grid.Width() + patternWidth) / 2

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

func (s *Simulator) evoluteCell(grid *grid.Grid, rowIndex, columnIndex int) {
	if s.cellLives(rowIndex, columnIndex) {
		grid.Populate(rowIndex, columnIndex)
	}
}

func (s *Simulator) cellLives(rowIndex, columnIndex int) bool {
	switch s.grid.AliveNeigborCount(rowIndex, columnIndex) {
	case 2:
		if s.grid.IsPopulated(rowIndex, columnIndex) {
			return true
		}
		return false
	case 3:
		return true
	default:
		return false
	}
}

func (s *Simulator) String() string {
	var out bytes.Buffer

	out.WriteString(s.grid.String())
	fmt.Fprintf(&out, "Generation: %d", s.GenerationCount)

	return out.String()
}

func (s *Simulator) Render(gc *draw2dimg.GraphicContext) bool {
	return s.grid.Render(gc)
}
