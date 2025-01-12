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


	for i, k := range tokens {
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
			result = append(result, k)
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
			result = append(result, k)
		}else {
			for len(stack) > 0 && precedence(k) <= precedence(stack[len(stack)-1]) &&
				(i > 0 && !k.IsUnary()) {
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			stack = append(stack, k)
		}
	}
	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	exp = append(exp, togroup(result))

	return exp
}

func precedence(t token.Token) uint{
	if t.TokenType == token.CARET {
		return 6
	}else if t.IsUnary() {
		return 5
	}else if t.TokenType >= token.STAR && t.TokenType <= token.SLASH {
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
	var start_group *group.Group = &group.Group{}
	var curr_group *group.Group = start_group
	fmt.Println(postfix)
	if len(postfix) == 1 {
		return group.Group{Operand1: postfix[0]}
	}
	for i := len(postfix)-1; i >= 0; i-- {
		if postfix[i].IsOperator() {
			if start_group != nil {
				g := &group.Group{Parent: curr_group, Operator: postfix[i]}
				if curr_group.Operand1 == nil {
					curr_group.Operand1 = g
				}else if curr_group.Operand2 == nil && !curr_group.Operator.IsUnary(){
					curr_group.Operand2 = curr_group.Operand1
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
				if curr_group.Operand1 == nil {
					curr_group.Operand1 = postfix[i]
					for curr_group.Operand1 != nil &&
						(curr_group.Operand2 != nil || curr_group.Operator.IsUnary()) &&
						curr_group.Parent != nil{
						curr_group = curr_group.Parent
					}
				}else if curr_group.Operand2 == nil && !curr_group.Operator.IsUnary() {
					curr_group.Operand2 = curr_group.Operand1
					curr_group.Operand1 = postfix[i]
					for curr_group.Operand1 != nil &&
						(curr_group.Operand2 != nil || curr_group.Operator.IsUnary()) &&
						curr_group.Parent != nil{
						curr_group = curr_group.Parent
					}
				}else {
					panic(fmt.Sprintf("Error parsing postfix: %s", curr_group))
				}

		}else if postfix[i].TokenType == token.R_BRACKET {
			

			if start_group != nil {
				g := &group.Group{Parent: curr_group}
				if curr_group.Operand1 == nil {
					curr_group.Operand1 = g
				}else if curr_group.Operand2 == nil && !curr_group.Operator.IsUnary(){
					curr_group.Operand2 = curr_group.Operand1
					curr_group.Operand1 = g
				}else {
					panic(fmt.Sprintf("Error parsing postfix: %s", curr_group))
				}
				curr_group = g
			}else {
				start_group = &group.Group{}
				curr_group = start_group
			}

		}else if postfix[i].TokenType == token.L_BRACKET {
			if curr_group.Parent != nil {
				curr_group = curr_group.Parent
			}
			for curr_group.Operand1 != nil &&
				(curr_group.Operand2 != nil || curr_group.Operator.IsUnary()) &&
				curr_group.Parent != nil{
				curr_group = curr_group.Parent
			}
		}else {
			panic("Found something other than operator or operand")
		}
	}
	if v, ok := start_group.Operand1.(*group.Group); ok {
		return *v
	}
	return *start_group
}

