%{

package simplelex

%}

%union {
    num float64
}

%type <num> expression term factor
%token '+' '-' '*' '/' '(' ')'
%token <num> NUMBER

%%
top        : expression
                {
                    if l, ok := yylex.(*simpleLex); ok {
                        l.value = $1
                    }
                }
           ;
expression : expression '+' term
                { $$ = $1 + $3 }
           | expression '-' term
                { $$ = $1 - $3 }
           | term
                { $$ = $1 }
           ;
term       : term '*' factor
                { $$ = $1 * $3 }
           | term '/' factor
                { $$ = $1 / $3 }
           | factor
                { $$ = $1 }
           ;
factor     : NUMBER
                { $$ = $1 }
           | '(' expression ')'
                { $$ = $2 }
           ;
%%