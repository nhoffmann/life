package grid

import (
	"bytes"
)

// Grid is a datastructure to contain an evolution of a game of life
type Grid struct {
	matrix [][]int
}

// New creates a new instance of Grid
func New(width, height int) Grid {
	grid := Grid{}

	grid.matrix = make([][]int, width)
	for i := 0; i < width; i++ {
		grid.matrix[i] = make([]int, height)
	}

	return grid
}

// Width returns the width of the current grids matrix
func (g *Grid) Width() int {
	return len(g.matrix)
}

// Height returns the height of the current grids matrix
func (g *Grid) Height() int {
	return len(g.matrix[0])
}

// IsPopulated returns true when the cell at the given coordinates is populated,
// it returns false otherwise
func (g *Grid) IsPopulated(x, y int) bool {
	if x < 0 || x >= g.Width() {
		return false
	}

	if y < 0 || y >= g.Height() {
		return false
	}

	return g.matrix[x][y] == 1
}

// Populate populates the cell at the given coordinates
func (g *Grid) Populate(x, y int) {
	g.matrix[x][y] = 1
}

func (g *Grid) String() string {
	var out bytes.Buffer

	for _, row := range g.matrix {
		for _, cell := range row {
			if cell == 1 {
				out.WriteString("O")
			} else {
				out.WriteString("•")
			}
		}

		out.WriteString("\n")
	}

	return out.String()
}
