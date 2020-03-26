package rle

import (
	"reflect"
	"testing"
)

func TestRLE(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		tests := []struct {
			input           string
			expectedPattern [][]int
			expectedComment string
			expectedWidth   int
			expectedHeight  int
			expectedRule    string
		}{
			{
				input: `#C This is a glider.
				x = 3, y = 3
				bo$2bo$3o!`,
				expectedPattern: [][]int{
					{0, 1, 0},
					{0, 0, 1},
					{1, 1, 1},
				},
				expectedWidth:   3,
				expectedHeight:  3,
				expectedComment: "This is a glider.",
				expectedRule:    "",
			},
			{
				input: `#N Gosper glider gun
				#C This was the first gun discovered.
				#C As its name suggests, it was discovered by Bill Gosper.
				x = 36, y = 9, rule = B3/S23
				24bo$22bobo$12b2o6b2o12b2o$11bo3bo4b2o12b2o$2o8bo5bo3b2o$2o8bo3bob2o4b
				obo$10bo5bo7bo$11bo3bo$12b2o!`,
				expectedPattern: [][]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
					{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
				expectedWidth:  36,
				expectedHeight: 9,
				expectedComment: `This was the first gun discovered.
				As its name suggests, it was discovered by Bill Gosper.`,
				expectedRule: "B3/S23",
			},
		}

		for _, test := range tests {
			rle, err := Parse(test.input)

			if err != nil {
				t.Error(err)
			}

			if rle.Comment != test.expectedComment {
				t.Errorf("Comment does not match")
			}

			if rle.Width != test.expectedWidth {
				t.Errorf("Width dos not match")
			}

			if rle.Height != test.expectedHeight {
				t.Errorf("Height does not match")
			}

			if rle.Rule != test.expectedRule {
				t.Errorf("Rule does not match")
			}

			if !reflect.DeepEqual(rle.Pattern, test.expectedPattern) {
				t.Errorf(
					"Patterns do not match.\nExpected: %v\nGot: %v",
					test.expectedPattern,
					rle.Pattern,
				)
			}
		}
	})

}
