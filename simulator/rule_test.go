package simulator

import (
	"reflect"
	"testing"
)

type ruleTest struct {
	ruleString            string
	expectedBornCounts    []int
	expectedSurviveCounts []int
}

func TestRule(t *testing.T) {
	t.Run("Conway Life (B3/S23)", func(t *testing.T) {
		tests := []ruleTest{
			// Life
			{
				ruleString:            "B3/S23",
				expectedBornCounts:    []int{3},
				expectedSurviveCounts: []int{2, 3},
			},
			// Highlife
			{
				ruleString:            "B36/S23",
				expectedBornCounts:    []int{3, 6},
				expectedSurviveCounts: []int{2, 3},
			},
			// Day & Night
			{
				ruleString:            "B3678/S34678",
				expectedBornCounts:    []int{3, 6, 7, 8},
				expectedSurviveCounts: []int{3, 4, 6, 7, 8},
			},
			// Seeds
			{
				ruleString:            "B2/S",
				expectedBornCounts:    []int{2},
				expectedSurviveCounts: []int{},
			},
		}

		for _, test := range tests {
			rule := ParseRule(test.ruleString)

			if !reflect.DeepEqual(rule.BornCounts, test.expectedBornCounts) {
				t.Errorf(
					"Born counts do not match. Expected %v, got %v",
					test.expectedBornCounts,
					rule.BornCounts,
				)
			}

			if !reflect.DeepEqual(rule.SurviveCounts, test.expectedSurviveCounts) {
				t.Errorf(
					"Survive counts do not match. Expected %v, got %v",
					test.expectedSurviveCounts,
					rule.SurviveCounts,
				)
			}
		}
	})
}
