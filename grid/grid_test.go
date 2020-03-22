package grid

import (
	"testing"
)

func TestGrid(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		grid := New(9, 9)

		if grid.Width() != 9 {
			t.Errorf("Grid width incorrect. Expected %d, got %d", 9, grid.Width())
		}

		if grid.Height() != 9 {
			t.Errorf("Grid height incorrect. Expected %d, got %d", 9, grid.Height())
		}
	})

	t.Run("IsPopulated(x, y)", func(t *testing.T) {
		grid := New(9, 9)

		if grid.IsPopulated(4, 5) {
			t.Errorf("Expected all cells to be unpopulatd.")
		}
	})

	t.Run("Populate(x, y)", func(t *testing.T) {
		grid := New(9, 9)

		grid.Populate(4, 5)

		if !grid.IsPopulated(4, 5) {
			t.Errorf("Expected cell 4/5 to be populatd.")
		}
	})
}
