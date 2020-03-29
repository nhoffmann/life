package grid

type Cell struct {
	X, Y int
}

func NewCell(x, y int) Cell {
	return Cell{x, y}
}

func (c *Cell) MooreNeighborhood() []Cell {
	return []Cell{
		NewCell(c.X-1, c.Y-1), // NW
		NewCell(c.X, c.Y-1),   // N
		NewCell(c.X+1, c.Y-1), // NE
		NewCell(c.X-1, c.Y),   // W
		NewCell(c.X+1, c.Y),   // E
		NewCell(c.X-1, c.Y+1), // SW
		NewCell(c.X, c.Y+1),   // S
		NewCell(c.X+1, c.Y+1), // SE
	}
}
