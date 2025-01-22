package parser

import (
	"fmt"
	"os"

	"github.com/Pyr0de/pod-interpreter/cmd/group"
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
		case token.INIT:
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
				fmt.Fprintf(os.Stderr, "[line %d] Error: Malformed init\n", tokens[j].Line)
				return code, true
			}
			if exp[0].Operand2 == nil {
				exp[0].Operand2 = token.Token{}
			}
			code = append(code, stmt.Stmt{
				Stype: token.INIT, Statement: stmt.StmtAssign{Expression: exp[0], Init: true},
			})
			i = j
		case token.IF:
			j := i+1
			for ;j < len(tokens);j++ {
				if tokens[j].TokenType == token.L_BRACE {
					break
				}
			}
			
			if tokens[j].TokenType != token.L_BRACE {
				fmt.Fprintf(os.Stderr, "[line %d] Error: Expected expression, found \"%s\"",
					tokens[j].Line, tokens[j].Raw)
				return code, true
			}
			exp, err := ParseExpression(tokens[i+1:j])
			if err != nil || len(exp) != 1 {
				
				fmt.Fprintf(os.Stderr, "[line %d] Error: Expected expression\n", tokens[j].Line)
				return code, true
			}
			i = j
			j++
			count := 1
			for ;j < len(tokens);j++ {
				if tokens[j].TokenType == token.R_BRACE {
					count--
					if count <= 0 {
						break
					}
				}
				if tokens[j].TokenType == token.L_BRACE {
					count++
				}
			}
			if tokens[j].TokenType != token.R_BRACE || count > 0 {
				fmt.Fprintf(os.Stderr, "[line %d] Error: Expected token \"}\"\n",
					tokens[i].Line, tokens[i].Raw)
				return code, true
			}
			s, e := Parse(tokens[i+1:j])
			if e {
				return code, true
			}
			code = append(code, stmt.Stmt{
				Stype: token.IF, Statement: stmt.StmtIf{Expression: exp[0], Block: stmt.StmtBlock{Block: s}},
			})
			i = j
		case token.ELSE:
			if iffy, ok := code[len(code) - 1].Statement.(stmt.StmtIf); ok {
				j := i+1
				count := 0
				for ;j < len(tokens); j++ {
					if tokens[j].TokenType == token.R_BRACE {
						count--
						if count <= 0 {
							break
						}
					}
					if tokens[j].TokenType == token.L_BRACE {
						count++
					}
				}
				if tokens[j].TokenType != token.R_BRACE || count > 0{
					fmt.Fprintf(os.Stderr, "[line %d] Error: Expected token \"}\"\n",
					tokens[i].Line)
					return code, true
				}
				if count < 0 {
					fmt.Fprintf(os.Stderr, "[line %d] Error: Found \"}\"\n",
					tokens[i].Line)
					return code, true
				}
				s, err := Parse(tokens[i+1:j+1])
				if err {
					return code, true
				}
				curr := &iffy
				for curr.Else != nil {
					curr = curr.Else
				}
				if if_stmt, ok := s[0].Statement.(stmt.StmtIf); ok {
					curr.Else = &if_stmt
				}else if else_stmt, ok := s[0].Statement.(stmt.StmtBlock); ok {
					exp := group.Group{Operand1: token.Token{TokenType: token.TRUE}}
					curr.Else = &stmt.StmtIf{Expression: exp, Block: else_stmt}
				}else {
					fmt.Fprintln(os.Stderr, "found something other than if/else: ", s[0])
					return code, true
				}
				code[len(code)-1].Statement = iffy
				
				i = j
			}else {
				fmt.Fprintf(os.Stderr, "[line %d] Error: Expected \"if\" before \"else\"\n",
					tokens[i].Line, tokens[i].Raw)
				return code, true
			}
		case token.L_BRACE:
			count := 1
			j := i+1
			for ;j < len(tokens);j++ {
				if tokens[j].TokenType == token.R_BRACE {
					count--
					if count <= 0 {
						break
					}
				}
				if tokens[j].TokenType == token.L_BRACE {
					count++
				}
			}
			if tokens[j].TokenType != token.R_BRACE || count > 0{
				fmt.Fprintf(os.Stderr, "[line %d] Error: Expected token \"}\"\n",
					tokens[i].Line, tokens[i].Raw)
				return code, true
			}
			s, err := Parse(tokens[i+1:j])
			if err {
				return code, true
			}
			code = append(code, stmt.Stmt{
				Stype: token.L_BRACE, Statement: stmt.StmtBlock{Block: s},
			})
			i = j
		default:
			j := i+1
			for ;j < len(tokens);j++ {
				if tokens[j].TokenType == token.SEMICOLON {
					break
				}
			}
			exp, err := ParseExpression(tokens[i:j])
			if err != nil {
				fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected token found \"%s\"\n",
					tokens[i].Line, tokens[i].Raw)
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
