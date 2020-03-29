package grid

import "testing"

func TestCell(t *testing.T) {
	t.Run("MooreNeighborhood", func(t *testing.T) {
		c := NewCell(0, 0)

		neighborhood := c.MooreNeighborhood()

		if len(neighborhood) != 8 {
			t.Errorf("Incorrect number of neighbors. Expected: %d, got %d", 8, len(neighborhood))
		}

		tests := []struct {
			expectedX int
			expectedY int
		}{
			{-1, -1},
			{0, -1},
			{1, -1},
			{-1, 0},
			{1, 0},
			{-1, 1},
			{0, 1},
			{1, 1},
		}

		for i, test := range tests {
			cellUnderTest := neighborhood[i]

			if cellUnderTest.X != test.expectedX {
				t.Errorf("X do not match for index: %d", i)
			}

			if cellUnderTest.Y != test.expectedY {
				t.Errorf("Y do not match for index: %d", i)
			}
		}
	})
}
