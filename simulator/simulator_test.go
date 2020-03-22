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
				[][]int{{0, 1, 0}, {0, 1, 0}, {0, 1, 0}},
				[][]int{{0, 0, 0}, {1, 1, 1}, {0, 0, 0}},
			},
		}

		for _, test := range tests {
			s := New(3, 3)
			s.LoadPattern(test.pattern, 0, 0)
			s.Evolute()

			expected := New(3, 3)
			expected.LoadPattern(test.expectedPattern, 0, 0)

			if !reflect.DeepEqual(s.grid, expected.grid) {
				t.Errorf("Evolute output incorrect.\nExpected: %q\n Got:%q", expected.grid, s.grid)
			}
		}
	})
}
