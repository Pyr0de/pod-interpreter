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
				isElse := false
				if if_stmt, ok := s[0].Statement.(stmt.StmtIf); ok {
					curr.Else = &if_stmt
				}else if else_stmt, ok := s[0].Statement.(stmt.StmtBlock); ok {
					exp := group.Group{Operand1: token.Token{TokenType: token.TRUE}}
					curr.Else = &stmt.StmtIf{Expression: exp, Block: else_stmt}
					isElse = true
				}else {
					fmt.Fprintln(os.Stderr, "found something other than if/else: ", s[0])
					return code, true
				}
				code[len(code)-1].Statement = iffy
				if isElse {
					code = append(code, stmt.Stmt{Stype: token.None, Statement: stmt.StmtEmpty{}})
				}
				i = j
			}else {
				fmt.Fprintf(os.Stderr, "[line %d] Error: Expected \"if\" before \"else\"\n",
					tokens[i].Line)
				return code, true
			}
		case token.WHILE:
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
				Stype: token.WHILE, Statement: stmt.StmtWhile{Expression: exp[0], Block: stmt.StmtBlock{Block: s}},
			})
			i = j
		case token.FOR:
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
			statements, err := Parse(tokens[i+1:j])
			if len(statements) == 2 {
				statements = append(statements, stmt.Stmt{Stype: token.SEMICOLON, Statement: stmt.StmtExpression{}})
			}
			if err || len(statements) != 3 {
				fmt.Fprintf(os.Stderr, "[line %d] Error: Expected 3 expressions\n", tokens[j].Line)
				return code, true
			}
			condition, ok := statements[1].Statement.(stmt.StmtExpression)
			if !ok {
				fmt.Fprintf(os.Stderr, "[line %d] Error: Expected 3 expressions\n", tokens[j].Line)
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
				Stype: token.FOR, Statement: stmt.StmtFor{
					Initialization: statements[0],
					Condition: condition.Expression,
					Block: stmt.StmtBlock{Block: s},
					Step: statements[2],
				},
			})
			i = j
		case token.FUNC:
			i++
			if tokens[i].TokenType != token.IDENTIFIER {
				fmt.Fprintf(os.Stderr, "[line %d] Error: Expected identifier after \"func\"\n",
				tokens[i].Line)
				return code, true
			}
			func_name := tokens[i]
			i++
			if tokens[i].TokenType != token.L_BRACKET {
				fmt.Fprintf(os.Stderr, "[line %d] Error: Expected L_BRACKET(\"(\") after \"IDENTIFIER\"\n",
				tokens[i].Line)
				return code, true
			}
			parameters := []token.Token{}
			i++
			for i < len(tokens) && tokens[i].TokenType != token.R_BRACKET {
				if tokens[i].TokenType == token.IDENTIFIER {
					if tokens[i+1].TokenType == token.COMMA && tokens[i+2].TokenType == token.IDENTIFIER{
						parameters = append(parameters, tokens[i])
						i += 2
					}else if tokens[i+1].TokenType == token.R_BRACKET{
						parameters = append(parameters, tokens[i])
						i += 1
						break
					}else {
						fmt.Fprintf(os.Stderr, "[line %d] Error: Malformed function declaration\n",
						tokens[i].Line)
						return code, true
						//error malformed function declaration
					}
				}else {
					fmt.Fprintf(os.Stderr, "[line %d] Error: Expected parameters, found \"%s\"\n",
					tokens[i].Line, tokens[i].Raw)
					return code, true
				}
			}
			if tokens[i].TokenType != token.R_BRACKET || tokens[i+1].TokenType != token.L_BRACE{
				fmt.Fprintf(os.Stderr, "[line %d] Error: Malformed function declaration\n",
				tokens[i].Line)
				return code, true
				// error malformed function declaration
			}
			i += 1
			j := i+1
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
			fmt.Println(func_name, parameters, s)
			//code = append(code, stmt.Stmt{
			//	Stype: token.WHILE, Statement: stmt.StmtWhile{Expression: exp[0], Block: stmt.StmtBlock{Block: s}},
			//})
			i = j
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
			j := i
			for ;j < len(tokens);j++ {
				if tokens[j].TokenType == token.SEMICOLON {
					break
				}
			}
			exp, err := ParseExpression(tokens[i:j])
			if j < len(tokens) {
				exp, err = ParseExpression(tokens[i:j+1])
			}
			if err != nil || len(exp) <= 0{
				fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected token found \"%s\"\n",
					tokens[i].Line, tokens[i].Raw)
				return code, true
			}
			if exp[0].Operator.TokenType == token.EQUAL {
				code = append(code, stmt.Stmt{
					Stype: token.INIT, Statement: stmt.StmtAssign{Expression: exp[0], Init: false},
				})
			}else {
				code = append(code, stmt.Stmt{
					Stype: tokens[i].TokenType, Statement: stmt.StmtExpression{Expression: exp[0]},
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
