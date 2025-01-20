package parser

import (
	"fmt"
	"os"

	"github.com/Pyr0de/pod-interpreter/cmd/stmt"
	"github.com/Pyr0de/pod-interpreter/cmd/token"
)

func Parse(tokens []token.Token) ([]stmt.Stmt, bool){
	brace_count := 0
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
			if exp[0].Operator.TokenType != token.EQUAL && exp[0].Operator.TokenType != token.None{
				fmt.Fprintf(os.Stderr, "[line %d] Error: Malformed init", tokens[j].Line)
				return code, true
			}
			if exp[0].Operand2 == nil {
				exp[0].Operand2 = token.Token{}
			}
			code = append(code, stmt.Stmt{
				Stype: token.INIT, Statement: stmt.StmtAssign{Expression: exp[0], Init: true},
			})
			i = j;
		}
		case token.L_BRACE:
			brace_count++
			code = append(code, stmt.Stmt{
				Stype: token.L_BRACE, Statement: stmt.StmtScope{Open: true},
			})
		case token.R_BRACE:
			code = append(code, stmt.Stmt{
				Stype: token.R_BRACE, Statement: stmt.StmtScope{Open: false},
			})

			brace_count--
			if brace_count < 0 {
				fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected brace found \"%s\"\n", tokens[i].Line, tokens[i].Raw)
				return code, true
			}
		default:
			j := i+1
			for ;j < len(tokens);j++ {
				if tokens[j].TokenType == token.SEMICOLON {
					break
				}
			}
			exp, err := ParseExpression(tokens[i:j])
			if err != nil {
				fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected token found \"%s\"\n", tokens[i].Line, tokens[i].Raw)
				return code, true
			}
			if exp[0].Operator.TokenType == token.EQUAL {
				code = append(code, stmt.Stmt{
					Stype: token.INIT, Statement: stmt.StmtAssign{Expression: exp[0], Init: false},
				})
			}
			i = j;
		}

	}
	
	if brace_count > 0 {
		fmt.Fprintf(os.Stderr, "[line %d] Error: Expected brace \"}\"\n",)
		return code, true
	}

	return code, false
}
