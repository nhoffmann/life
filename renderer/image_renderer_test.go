package renderer

import (
	"testing"

	"github.com/nhoffmann/life/grid"
)

func TestImageRenderer(t *testing.T) {
	t.Run("coordinates", func(t *testing.T) {
		ir := NewImageRenderer(50, 50)

		cell := grid.NewCell(0, 0)

		x, y := ir.coordinates(cell)

		if x != 20 {
			t.Errorf("X coordinate is not correct. Got %f, want %f", x, 20.0)
		}

		if y != 20 {
			t.Errorf("Y coordinate is not correct. Got %f, want %f", y, 20.0)
		}
	})
}
