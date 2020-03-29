package simulator

import (
	"reflect"
	"testing"
)

func TestSimulator(t *testing.T) {
	t.Run("Evolute()", func(t *testing.T) {
		tests := []struct {
			pattern         [][]bool
			expectedPattern [][]bool
		}{
			{
				[][]bool{{false, false, false}, {false, false, false}, {false, false, false}},
				[][]bool{{false, false, false}, {false, false, false}, {false, false, false}},
			},
			{
				[][]bool{{false, false, false}, {false, true, false}, {false, false, false}},
				[][]bool{{false, false, false}, {false, false, false}, {false, false, false}},
			},
			{
				[][]bool{
					{false, false, false, false, false},
					{false, false, true, false, false},
					{false, false, true, false, false},
					{false, false, true, false, false},
					{false, false, false, false, false},
				},
				[][]bool{
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, true, true, true, false},
					{false, false, false, false, false},
					{false, false, false, false, false},
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
