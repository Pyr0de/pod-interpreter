package parser

import (
	"fmt"
	"os"

	"github.com/Pyr0de/pod-interpreter/cmd/stmt"
	"github.com/Pyr0de/pod-interpreter/cmd/token"
)

func Parse(tokens []token.Token) ([]stmt.Stmt, bool){
	code := []stmt.Stmt{}
	for i, k := range tokens {
		switch k.TokenType {
		case token.PRINT:
			j := i+1
			for ;j < len(tokens);j++ {
				if tokens[j].TokenType == token.SEMICOLON {
					break
				}
			}
			exp, err := ParseExpression(tokens[i+1:j])
			if err != nil || len(exp) != 1{
				fmt.Fprintf(os.Stderr, "[line %d] Error: Expected an expression", tokens[j].Line)
			}

			code = append(code, stmt.Stmt{
				Stype: token.PRINT, Statement: stmt.StmtPrint{Expression: exp[0]},
			})
			

		}
		
	}
	
	return code, false
}
