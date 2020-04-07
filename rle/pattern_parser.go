package rle

import (
	"strconv"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	RUN_COUNT  = "RUN_COUNT"
	DEAD_CELL  = "DEAD_CELL"
	ALIVE_CELL = "ALIVE_CELL"
	EOL        = "EOL"
	EOP        = "EOP"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.char {
	case '$':
		tok = newToken(EOL, l.char)
	case '!':
		tok = newToken(EOP, l.char)
	case 'b':
		tok = newToken(DEAD_CELL, l.char)
	case 'o':
		tok = newToken(ALIVE_CELL, l.char)
	default:
		if isDigit(l.char) {
			tok.Type = RUN_COUNT
			tok.Literal = l.readNumber()
			return tok
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType TokenType, char byte) Token {
	return Token{Type: tokenType, Literal: string(char)}
}

type PatternParser struct {
	lexer        *Lexer
	currentToken Token
	peekToken    Token
}

func NewParser(lexer *Lexer) *PatternParser {
	p := &PatternParser{
		lexer: lexer,
	}
	p.nextToken()
	p.nextToken()

	return p
}

func (pp *PatternParser) ParsePattern(width, height int) [][]int {
	result := make([][]int, height)

	row := make([]int, width)
	var rowIndex int
	var colIndex int
	for {
		switch pp.currentToken.Type {
		case RUN_COUNT:
			count, _ := strconv.Atoi(pp.currentToken.Literal)
			for i := 0; i < count; i++ {
				switch pp.peekToken.Type {
				case ALIVE_CELL:
					row[rowIndex+i] = 1
				case DEAD_CELL:
					row[rowIndex+i] = 0
				case EOL:
					result[colIndex] = row
					row = make([]int, width)
					rowIndex = -1
					colIndex++
				}
			}

			if pp.peekToken.Type != EOL {
				rowIndex += count - 1
			}
			pp.nextToken()
		case ALIVE_CELL:
			row[rowIndex] = 1
		case DEAD_CELL:
			row[rowIndex] = 0
		case EOL:
			result[colIndex] = row
			row = make([]int, width)
			rowIndex = -1
			colIndex++
		case EOP:
			result[colIndex] = row
			return result
		}
		rowIndex++
		pp.nextToken()
	}
}

func (pp *PatternParser) nextToken() {
	pp.currentToken = pp.peekToken
	pp.peekToken = pp.lexer.NextToken()
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.char) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}
