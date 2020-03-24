package simulator

import (
	"strconv"
	"strings"
)

type Rule struct {
	BornCounts    []int
	SurviveCounts []int
}

func ParseRule(ruleString string) Rule {
	rule := Rule{}

	ruleSplit := strings.Split(ruleString, "/")

	rule.BornCounts = convertStringArrayToIntArray(strings.Split(ruleSplit[0], "")[1:])
	rule.SurviveCounts = convertStringArrayToIntArray(strings.Split(ruleSplit[1], "")[1:])

	return rule
}

func convertStringArrayToIntArray(stringArray []string) []int {
	intArray := make([]int, len(stringArray))

	for index, element := range stringArray {
		intArray[index], _ = strconv.Atoi(element)
	}

	return intArray
}
