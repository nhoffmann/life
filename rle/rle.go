package rle

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type RLE struct {
	Comment    string  // #C
	Name       string  // #N
	Originator string  // #O
	Rule       string  // #R or rule
	Position   string  // #P
	Width      int     // x
	Height     int     // y
	Pattern    [][]int // The actual pattern

	inputLines       []string
	headerLineIndex  int
	patternLineIndex int
}

func Parse(input string) (RLE, error) {
	rle := RLE{
		inputLines: strings.Split(input, "\n"),
	}

	rle.partitionFile()

	err := rle.parseComments()
	if err != nil {
		return RLE{}, err
	}
	err = rle.parseHeader()
	if err != nil {
		return RLE{}, err
	}
	err = rle.parsePattern()
	if err != nil {
		return RLE{}, err
	}

	return rle, nil
}

func (rle *RLE) partitionFile() error {
	for index, line := range rle.inputLines {
		cleanLine := removeWhitespace(line)
		if strings.HasPrefix(cleanLine, "x=") {
			rle.headerLineIndex = index
			rle.patternLineIndex = index + 1
			return nil
		}
	}

	return fmt.Errorf("Invlaid input: Header is missing")
}

func (rle *RLE) parseComments() error {
	return nil
}

func (rle *RLE) parseHeader() (err error) {
	headerLine := removeWhitespace(rle.inputLines[rle.headerLineIndex])

	headerElements := strings.SplitN(headerLine, ",", 3)

	rle.Width, err = strconv.Atoi(strings.TrimPrefix(headerElements[0], "x="))
	if err != nil {
		return err
	}
	rle.Height, err = strconv.Atoi(strings.TrimPrefix(headerElements[1], "y="))
	if err != nil {
		return err
	}

	rle.Pattern = make([][]int, rle.Width)

	// check wehter a rule is present, since it's optional
	if len(headerElements) == 3 {
		rle.Rule = strings.TrimPrefix(headerElements[2], "rule=")
	}

	return nil
}

func (rle *RLE) parsePattern() error {
	patternString := strings.Join(rle.inputLines[rle.patternLineIndex:], "")

	l := NewLexer(patternString)
	pp := NewParser(l)

	rle.Pattern = pp.ParsePattern(rle.Width, rle.Height)

	return nil
}

func removeWhitespace(input string) string {
	re := regexp.MustCompile(` *\t*\r*\n*`)
	return re.ReplaceAllString(input, "")
}
