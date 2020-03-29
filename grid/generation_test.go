package grid

import "testing"

func TestGeneration(t *testing.T) {
	t.Run("Populate", func(t *testing.T) {
		g := NewGeneration()

		g.Populate(NewCell(-2, 4))

		if len(g.CellMap) != 1 {
			t.Errorf("Number of cells is not correct.")
		}

		g.Populate(NewCell(-2, 4))

		if len(g.CellMap) != 1 {
			t.Errorf("Expected cell with same coordinates as existing one to be added only once.")
		}
	})

	t.Run("IsPopulated", func(t *testing.T) {
		g := NewGeneration()

		g.Populate(NewCell(-2, 4))

		if !g.IsPopulated(NewCell(-2, 4)) {
			t.Errorf("Expected cell {-2, 4} to be populated")
		}

		if g.IsPopulated(NewCell(3, 1)) {
			t.Errorf("Expected cell {3, 1} to not be populated")
		}
	})

	t.Run("LoadPattern", func(t *testing.T) {
		g := NewGeneration()

		g.LoadPattern([][]bool{
			{false, false, false},
			{true, true, true},
			{false, false, false},
		})

		expectedCells := []Cell{
			NewCell(-1, 0),
			NewCell(0, 0),
			NewCell(1, 0),
		}

		for _, cell := range expectedCells {
			if !g.IsPopulated(cell) {
				t.Errorf("Expected cell to be populated: %v. Generation: %v", cell, g)
			}
		}
	})
}
