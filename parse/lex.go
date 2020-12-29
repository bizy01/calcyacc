package parse

import (
	"text/scanner"
	"strings"
	"strconv"
	"errors"
)

func Parse(str string) (Result, error) {
	// parse
	l := newLex(str)

	yyParse(l)

	return l.result, l.err
}

type Result float64

type lex struct{
	scan scanner.Scanner
	result Result
	err    error
}

func newLex(str string) *lex {
	l := new(lex)

	l.scan.Init(strings.NewReader(str))

	return l
}

func (l *lex) Lex(lval *yySymType) int {
	tok, text := l.scan.Scan(),  l.scan.TokenText()

	switch tok {
		case scanner.EOF:
			return 0
		case scanner.Int:
			i, _ := strconv.Atoi(text)
        	lval.result = Result(i)
        	return	NUMBER
		case scanner.Float:
			val, _ := strconv.ParseFloat(text, 64)
			lval.result = Result(val)
			return	NUMBER
		default:
			return  int(tok)
	}
}

func (l *lex) Error(s string) {
	l.err = errors.New(s)
}