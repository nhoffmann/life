package simulator

import (
	"reflect"
	"testing"
)

func TestSimulator(t *testing.T) {
	t.Run("Evolute()", func(t *testing.T) {
		tests := []struct {
			pattern         [][]int
			expectedPattern [][]int
		}{
			{
				[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
				[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			},
			{
				[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
				[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			},
			{
				[][]int{
					{0, 0, 0, 0, 0},
					{0, 0, 1, 0, 0},
					{0, 0, 1, 0, 0},
					{0, 0, 1, 0, 0},
					{0, 0, 0, 0, 0},
				},
				[][]int{
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 1, 1, 1, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0},
				},
			},
		}

		for _, test := range tests {
			width := len(test.pattern[0])
			height := len(test.pattern)

			s := New(width, height, "B3/S23")
			err := s.LoadPatternAt(test.pattern, 0, 0)
			if err != nil {
				t.Fatalf("Error loading pattern: %s", err)
			}

			s.Evolute()

			expected := New(width, height, "B3/S23")
			err = expected.LoadPatternAt(test.expectedPattern, 0, 0)
			if err != nil {
				t.Fatalf("Error loading pattern: %s", err)
			}

			if !reflect.DeepEqual(s.grid, expected.grid) {
				t.Errorf("Evolute output incorrect.\nExpected: %q\n Got:%q", expected.grid, s.grid)
			}
		}
	})
}
