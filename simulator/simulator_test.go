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
				[][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
				[][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			{
				[][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
				[][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
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
			s := NewSimulatorWithPattern(test.pattern, "B3/S23")

			s.Evolute()

			expected := NewSimulatorWithPattern(test.expectedPattern, "B3/S23")

			if !reflect.DeepEqual(s.Generation(), expected.Generation()) {
				t.Errorf("Evolute output incorrect.\nExpected: %v\n Got:%v", expected.Generation(), s.Generation())
			}
		}
	})
}
