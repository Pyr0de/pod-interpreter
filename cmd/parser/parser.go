package parser

import (
	"fmt"
	"os"

	"github.com/Pyr0de/pod-interpreter/cmd/stmt"
	"github.com/Pyr0de/pod-interpreter/cmd/token"
)

func Parse(tokens []token.Token) ([]stmt.Stmt, bool){
	code := []stmt.Stmt{}
	for i := 0; i < len(tokens); i++ {
		k := tokens[i]
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
				return code, true
			}

			code = append(code, stmt.Stmt{
				Stype: token.PRINT, Statement: stmt.StmtPrint{Expression: exp[0]},
			})
			i = j;
		case token.INIT: {
			j := i+1
			for ;j < len(tokens);j++ {
				if tokens[j].TokenType == token.SEMICOLON {
					break
				}
			}
			exp, err := ParseExpression(tokens[i+1:j])
			if err != nil || len(exp) != 1{
				fmt.Fprintf(os.Stderr, "[line %d] Error: Expected an expression", tokens[j].Line)
				return code, true
			}
			if exp[0].Operator.TokenType != token.EQUAL {
				fmt.Fprintf(os.Stderr, "[line %d] Error: Malformed init", tokens[j].Line)
				return code, true
			}
			code = append(code, stmt.Stmt{
				Stype: token.INIT, Statement: stmt.StmtInit{Expression: exp[0]},
			})
			i = j;
		}
		default:
			fmt.Fprintf(os.Stderr, "[line %d] Error: Unknown token found \"%s\"\n", tokens[i].Line, tokens[i].Raw)
			break
		}
		
	}
	
	return code, false
}
