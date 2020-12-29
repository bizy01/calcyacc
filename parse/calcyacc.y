%{
	package parse

	func setResult(l yyLexer, v Result) {
		if l, ok := l.(*lex); ok {
      l.result = v
    }
	}
%}

%union{
	result Result
}

%type <result> expression term factor
%token '+' '-' '*' '/' '(' ')'
%token <result> NUMBER

%start main

%%

main: expression
  {
    setResult(yylex, $1)
  }

expression : expression '+' term
                { $$ = $1 + $3 }
           | expression '-' term
                { $$ = $1 - $3 }
           | term
                { $$ = $1 }
term       : term '*' factor
                { $$ = $1 * $3 }
           | term '/' factor
                { $$ = $1 / $3 }
           | factor
                { $$ = $1 }
factor     : NUMBER
                { $$ = $1 }
           | '(' expression ')'
                { $$ = $2 }
%%


