// filename: calcLex.go
package simplelex

import (
    "text/scanner"
    "log"
    "strconv"
    "strings"
)

var LexPrintToken = false

type Token struct {
    Type int
    Str string
}

type simpleLex struct {
    scanner.Scanner
    value float64
}

// 这个接口必须实现，是词法分析的入口
func (s *simpleLex) Lex(lval *yySymType) int {
    r, lit := s.Scan(), s.TokenText()
    var token Token
    token.Str = lit

    switch r {
    case scanner.EOF:
        return 0
    case scanner.Int:
        i, _ := strconv.Atoi(lit)
        lval.num = float64(i)
        token.Type = scanner.Float
    case scanner.Float:
        lval.num, _ = strconv.ParseFloat(lit, 64)
        token.Type = scanner.Float
    default:
        token.Type = int(r)
    }

    // if LexPrintToken {
    //     fmt.Printf("<<token: %s, %s>>\r\n",scanner.TokenString(token.type), token.Str)
    // }
    if token.Type == scanner.Float {
        return NUMBER
    } else {
        return token.Type
    }
    // return token.Type
}

// 词法分析异常处理 该接口必须实现
func (s *simpleLex) Error(s1 string) {
    log.Printf("parse error: %s", s1)
}

// 计算入口
func Parse(code string) float64{
    s := new(simpleLex)
    s.Init(strings.NewReader(code))
    yyParse(s)
    return s.value
}