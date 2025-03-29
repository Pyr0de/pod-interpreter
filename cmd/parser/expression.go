package parser

import (
	"errors"
	"fmt"
	"os"

	"github.com/Pyr0de/pod-interpreter/cmd/group"
	"github.com/Pyr0de/pod-interpreter/cmd/token"
)

func ParseExpression(tokens []token.Token) ([]group.Group, error) {
	exp := []group.Group{}

	stack := []token.Token{}
	result := []token.Token{}
	err := false

	for i, k := range tokens {
		if k.TokenType == token.SEMICOLON {
			for len(stack) > 0 {
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(result) > 0 {
				a, e := togroup(result)
				if e {
					err = true
				} else if !a.Empty() {
					exp = append(exp, a)
				}
			} else {
				exp = append(exp, group.Group{})
			}
			result = []token.Token{}
		} else if k.IsOperand() {
			result = append(result, k)
		} else if k.TokenType == token.L_BRACKET {
			stack = append(stack, k)
			result = append(result, k)
		} else if k.TokenType == token.R_BRACKET {
			for len(stack) > 0 && stack[len(stack)-1].TokenType != token.L_BRACKET {
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				panic("Error in parsing expression")
			}
			result = append(result, k)
		} else {
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
	if len(result) > 0 {
		a, e := togroup(result)
		if e {
			err = true
		} else if !a.Empty() {
			exp = append(exp, a)
		}
	}
	if err {
		return exp, errors.New("Error Parser")
	}
	return exp, nil
}

func precedence(t token.Token) uint {
	if t.IsUnary() {
		return 7
	} else if t.TokenType == token.CARET {
		return 6
	} else if t.TokenType >= token.STAR && t.TokenType <= token.PERCENT{
		return 5
	} else if t.TokenType >= token.PLUS && t.TokenType <= token.MINUS {
		return 4
	} else if t.TokenType >= token.EQUAL_EQUAL && t.TokenType <= token.GREATER_EQUAL {
		return 3
	} else if t.TokenType == token.EQUAL {
		return 2
	}else if t.TokenType >= token.AND_AND && t.TokenType <= token.PIPE_PIPE {
		return 1
	}
	return 0
}

func togroup(postfix []token.Token) (group.Group, bool) {
	if len(postfix) == 1 && postfix[0].IsOperand() {
		return group.Group{Operand1: postfix[0]}, false
	}

	var start *group.Group = &group.Group{}
	groups := grouper{Start_group: start, Curr_group: start}

	i := len(postfix) - 1
	for ; i >= 0; i-- {
		if postfix[i].IsOperator() {
			g := &group.Group{Parent: groups.Curr_group, Operator: postfix[i]}
			groups.add_to_curr(g)
			groups.Curr_group = g
		} else if postfix[i].IsOperand() {
			groups.add_to_curr(postfix[i])
			groups.back()
		} else if postfix[i].TokenType == token.R_BRACKET {
			g := &group.Group{Parent: groups.Curr_group}
			groups.add_to_curr(g)
			groups.Curr_group = g
		} else if postfix[i].TokenType == token.L_BRACKET {
			if !groups.valid() {
				i--
				break
			}
			if groups.Curr_group.Parent != nil {
				groups.Curr_group = groups.Curr_group.Parent
			}
			groups.back()
		} else {
			panic(fmt.Sprintln("Found something other than operator or operand:", postfix[i]))
		}
	}
	if !groups.valid() {
		fmt.Fprintf(os.Stderr, "[line %d] Error at '%s': Expected expression\n", postfix[i+1].Line, postfix[i+1].Raw)
	}
	if v, ok := groups.Start_group.Operand1.(*group.Group); ok {
		return *v, !groups.valid()
	}
	return *groups.Start_group, !groups.valid()
}

type grouper struct {
	Start_group *group.Group
	Curr_group  *group.Group
}

func (grouper *grouper) back() {
	for grouper.Curr_group.Operand1 != nil &&
		(grouper.Curr_group.Operand2 != nil || grouper.Curr_group.Operator.IsUnary()) &&
		grouper.Curr_group.Parent != nil {
		grouper.Curr_group = grouper.Curr_group.Parent
	}
}

func (grouper *grouper) add_to_curr(g any) {
	if grouper.Curr_group.Operand1 == nil {
		grouper.Curr_group.Operand1 = g
	} else if grouper.Curr_group.Operand2 == nil && !grouper.Curr_group.Operator.IsUnary() {
		grouper.Curr_group.Operand2 = grouper.Curr_group.Operand1
		grouper.Curr_group.Operand1 = g
	} else {
		panic(fmt.Sprintf("Error parsing postfix: %s", grouper.Curr_group))
	}
}

func (grouper *grouper) valid() bool {
	if grouper.Curr_group.Operator.TokenType == token.None {
		return grouper.Curr_group.Operand1 != nil && grouper.Curr_group.Operand2 == nil
	}
	return grouper.Curr_group.Operand1 != nil && (grouper.Curr_group.Operand2 != nil || grouper.Curr_group.Operator.IsUnary())
}
