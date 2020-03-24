package simulator

import (
	"bytes"
	"fmt"
	"log"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/nhoffmann/life/grid"
)

type Simulator struct {
	Rows            int
	Columns         int
	GenerationCount int
	rule            Rule
	grid            *grid.Grid
}

func New(rows, columns int, ruleString string) *Simulator {
	return &Simulator{
		Rows:    rows,
		Columns: columns,
		rule:    ParseRule(ruleString),
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
func (s *Simulator) LoadPattern(pattern [][]int) error {
	patternWidth := len(pattern[0])
	patternHeigth := len(pattern)

	startX := (s.grid.Width() - patternWidth) / 2
	startY := (s.grid.Height() - patternHeigth) / 2

	return s.LoadPatternAt(pattern, startX, startY)
}

func (s *Simulator) LoadPatternAt(pattern [][]int, startX, startY int) error {
	patternWidth := len(pattern[0])
	patternHeight := len(pattern)

	if patternWidth > s.grid.Width() || patternHeight > s.grid.Height() {
		log.Print(patternWidth >= s.grid.Width())
		return fmt.Errorf("pattern too big for current grid")
	}

	for rowIndex, patternRow := range pattern {
		for colIndex := range patternRow {
			if pattern[rowIndex][colIndex] == 1 {
				s.Populate(colIndex+startX, rowIndex+startY)
			}
		}
	}

	return nil
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
	count := s.grid.AliveNeigborCount(rowIndex, columnIndex)

	if s.grid.IsPopulated(rowIndex, columnIndex) {
		// apply survive rules
		for _, surviveCount := range s.rule.SurviveCounts {
			if surviveCount == count {
				return true
			}
		}
	} else {
		// apply born rules
		for _, bornCount := range s.rule.BornCounts {
			if bornCount == count {
				return true
			}
		}

		return false
	}

	return false
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
