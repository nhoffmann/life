package rle

import (
	"reflect"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "bo$2bo$3o!"

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{DEAD_CELL, "b"},
		{ALIVE_CELL, "o"},
		{EOL, "$"},
		{RUN_COUNT, "2"},
		{DEAD_CELL, "b"},
		{ALIVE_CELL, "o"},
		{EOL, "$"},
		{RUN_COUNT, "3"},
		{ALIVE_CELL, "o"},
		{EOP, "!"},
	}

	l := NewLexer(input)

	for _, test := range tests {
		token := l.NextToken()

		if token.Type != test.expectedType {
			t.Errorf("Token typ not correct")
		}

		if token.Literal != test.expectedLiteral {
			t.Errorf("Literal not correct")
		}
	}
}

func TestParsePattern(t *testing.T) {
	tests := []struct {
		input    string
		expected [][]int
		width    int
		height   int
	}{
		{
			input: "bo$2bo$3o!",
			expected: [][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 1, 1},
			},
			width:  3,
			height: 3,
		},
		{
			input: `24bo$22bobo$12b2o6b2o12b2o$11bo3bo4b2o12b2o$2o8bo5bo3b2o$2o8bo3bob2o4b
							obo$10bo5bo7bo$11bo3bo$12b2o!`,
			expected: [][]int{
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
			width:  36,
			height: 9,
		},
	}

	for _, test := range tests {
		l := NewLexer(test.input)
		pp := NewParser(l)
		result := pp.ParsePattern(test.width, test.height)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf(
				"Patterns do not match.\nExpected: %v\nGot: %v",
				test.expected,
				result,
			)
		}
	}
}
