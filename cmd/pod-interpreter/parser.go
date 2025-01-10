package main

import (
	"fmt"

	"github.com/Pyr0de/pod-interpreter/cmd/group"
	"github.com/Pyr0de/pod-interpreter/cmd/token"
)


func Parse(tokens []token.Token) []group.Group {
	exp := []group.Group{}

	//var start_group *group.Group = nil
	//var curr_group *group.Group = nil

	stack := []token.Token{}
	result := []token.Token{}
	
	for _, k := range tokens {
		if k.TokenType == token.SEMICOLON {
			for len(stack) > 0 {
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			fmt.Println(result)
			result = []token.Token {}
		}else if k.IsOperand() {
			result = append(result, k)
		}else if k.TokenType == token.L_BRACKET {
			stack = append(stack, k)
		}else if k.TokenType == token.R_BRACKET {
			for len(stack) > 0 && stack[len(stack)-1].TokenType != token.L_BRACKET {
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}else {
				panic("Error in parsing expression")
			}
		}else {
			for len(stack) > 0 && precedence(k) <= precedence(stack[len(stack)-1]) {
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			
			stack = append(stack, k)
		}
	}

	return exp
}

func precedence(t token.Token) uint{
	if t.TokenType == token.BANG {
		return 5
	}else if t.TokenType == token.CARET {
		return 4
	}else if t.TokenType >= token.STAR && t.TokenType <= token.PERCENT {
		return 3
	}else if t.TokenType >= token.PLUS && t.TokenType <= token.MINUS {
		return 2
	}else if t.TokenType >= token.EQUAL && t.TokenType <= token.PIPE_PIPE {
		return 1
	}
	return 0
}
