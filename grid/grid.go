package grid

import (
	"bytes"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

// Grid is a datastructure to contain an evolution of a game of life
type Grid struct {
	matrix [][]int

	cellWidth  float64
	cellHeight float64
}

// New creates a new instance of Grid
func New(width, height int) *Grid {
	grid := Grid{cellWidth: 10.0, cellHeight: 10.0}

	grid.matrix = make([][]int, width)
	for i := 0; i < width; i++ {
		grid.matrix[i] = make([]int, height)
	}

	return &grid
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
	return g.matrix[x][y] == 1
}

// ValueAt returns the value at a given grid coordinate. Wraps around at the edges.
func (g *Grid) ValueAt(x, y int) int {
	switch {
	case x < 0:
		return g.ValueAt(g.Width()-1, y)
	case x >= g.Width():
		return g.ValueAt(0, y)
	case y < 0:
		return g.ValueAt(x, g.Height()-1)
	case y >= g.Height():
		return g.ValueAt(x, 0)
	}

	return g.matrix[x][y]
}

// AliveNeigborCount returns the number of alive cells in the Moore Neighborhood
// of a given cell coordinate
func (g *Grid) AliveNeigborCount(x, y int) int {
	return g.ValueAt(x-1, y-1) + // NW
		g.ValueAt(x, y-1) + // N
		g.ValueAt(x+1, y-1) + // NE
		g.ValueAt(x-1, y) + // W
		g.ValueAt(x+1, y) + // E
		g.ValueAt(x-1, y+1) + // SW
		g.ValueAt(x, y+1) + // S
		g.ValueAt(x+1, y+1) // SE
}

// Populate populates the cell at the given coordinates
func (g *Grid) Populate(x, y int) {
	g.matrix[x][y] = 1
}

// Render paints the grid on a given draw2dimg.GraphicContext
func (g *Grid) Render(gc *draw2dimg.GraphicContext) bool {
	black := color.RGBA{0x00, 0x00, 0x00, 0xFF}
	white := color.RGBA{0xff, 0xff, 0xff, 0xff}

	gc.SetFillColor(white)
	gc.Clear()

	gc.SetLineWidth(1)

	gc.SetStrokeColor(white)
	gc.SetFillColor(black)
	gc.BeginPath()
	for rowIndex, row := range g.matrix {
		for columnIndex, cell := range row {
			if cell == 1 {
				x := float64(rowIndex) * g.cellWidth
				y := float64(columnIndex) * g.cellHeight

				draw2dkit.Rectangle(gc, x, y, x+g.cellWidth, y+g.cellHeight)
			}
		}
	}
	gc.FillStroke()
	gc.Close()

	return true
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
