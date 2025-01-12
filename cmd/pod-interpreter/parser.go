package main

import (
	"fmt"

	"github.com/Pyr0de/pod-interpreter/cmd/group"
	"github.com/Pyr0de/pod-interpreter/cmd/token"
)


func Parse(tokens []token.Token) []group.Group {
	exp := []group.Group{}

	stack := []token.Token{}
	result := []token.Token{}

	for _, k := range tokens {
		if k.TokenType == token.SEMICOLON {
			for len(stack) > 0 {
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			exp = append(exp, togroup(result))
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
	if t.TokenType == token.CARET {
		return 6
	}else if t.TokenType == token.BANG {
		return 5
	}else if t.TokenType >= token.STAR && t.TokenType <= token.PERCENT {
		return 4
	}else if t.TokenType >= token.PLUS && t.TokenType <= token.MINUS {
		return 3
	}else if t.TokenType >= token.EQUAL_EQUAL && t.TokenType <= token.PIPE_PIPE {
		return 2
	}else if t.TokenType == token.EQUAL {
		return 1
	}
	return 0
}

func togroup(postfix []token.Token) group.Group {
	var start_group *group.Group
	var curr_group *group.Group

	for i := len(postfix)-1; i >= 0; i-- {
		if postfix[i].IsOperator() {
			if start_group != nil {
				g := &group.Group{Parent: curr_group, Operator: postfix[i]}
				if curr_group.Operand2 == nil && curr_group.Operator.TokenType != token.BANG{
					//add new group with operator2 with operator
					curr_group.Operand2 = g
				}else if curr_group.Operand1 == nil{
					//add new group with operator1 with operator
					curr_group.Operand1 = g
				}else {
					panic(fmt.Sprintf("Error parsing postfix: %s", curr_group))
				}
				curr_group = g
			}else {
				start_group = &group.Group{Operator: postfix[i]}
				curr_group = start_group
			}
		}else if postfix[i].IsOperand() {
				if curr_group.Operand2 == nil && curr_group.Operator.TokenType != token.BANG{
					//add new group with operator2 with operator
					curr_group.Operand2 = postfix[i]
				}else if curr_group.Operand1 == nil{
					//add new group with operator1 with operator
					curr_group.Operand1 = postfix[i]
					for curr_group.Operand1 != nil &&
						(curr_group.Operand2 != nil || curr_group.Operator.TokenType == token.BANG) &&
						curr_group.Parent != nil{
						curr_group = curr_group.Parent
					}
				}else {
					panic(fmt.Sprintf("Error parsing postfix: %s", curr_group))
				}
		}else {
			panic("Found something other than operator or operand")
		}
	}
	return *start_group
}
