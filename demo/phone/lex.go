package demo

import (
	"bytes"
	"errors"
	"unicode"
)

type Result int

func Parse(input []byte) (Result, error) {
	l := newLex(input)
	_ = yyParse(l)
	return l.result, l.err
}


type lex struct {
	input  *bytes.Buffer
	result Result
	err    error
}

func newLex(input []byte) *lex {
	return &lex{
		input: bytes.NewBuffer(input),
	}
}

func (l *lex) Lex(lval *yySymType) int {
	b := l.nextb()
	if unicode.IsDigit(rune(b)) {
		return D
	}
	return int(b)
}

func (l *lex) nextb() byte {
	b, err := l.input.ReadByte()
	if err != nil {
		return 0
	}
	return b
}

// Error satisfies yyLexer.
func (l *lex) Error(s string) {
	l.err = errors.New(s)
}
