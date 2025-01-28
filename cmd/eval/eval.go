package eval

import (
	"fmt"
	"os"

	"github.com/Pyr0de/pod-interpreter/cmd/env"
	"github.com/Pyr0de/pod-interpreter/cmd/group"
	"github.com/Pyr0de/pod-interpreter/cmd/token"
)

func Evaluate(in_group group.Group) (token.Token, bool){
	if group1, ok := in_group.Operand1.(*group.Group); ok {
		g, err :=  Evaluate(*group1)
		if err {
			return token.Token{}, true
		}
		in_group.Operand1 = g
	}else if _, ok := in_group.Operand1.(token.Token); !ok {
		panic("Error evaluating operand 1")
	}
	
	if !in_group.Operator.IsUnary() && in_group.Operator.TokenType != token.None{

		if group2, ok := in_group.Operand2.(*group.Group); ok {
			g, err :=  Evaluate(*group2)
			if err {
				return token.Token{}, true
			}
			in_group.Operand2 = g
		}else if _, ok := in_group.Operand2.(token.Token); !ok {
			panic("Error evaluating operand 2")
		}
	}
	
	t1, ok1 := in_group.Operand1.(token.Token)
	if t1.TokenType == token.IDENTIFIER {
		op, err := env.GetVar(nil, t1.Raw)
		if err {
			fmt.Fprintf(os.Stderr, "[line %d] Error: Uninitialized variable \"%s\"\n",
				t1.Line, t1.Raw)
			return token.Token{}, true
		}
		t1 = op
	}else if t1.IsBool() {
		t1.Value = t1.TokenType == token.TRUE
	}
	if in_group.Operator.TokenType == token.None {
		return t1, false
	}

	t2, ok2 := in_group.Operand2.(token.Token)
	if t2.TokenType == token.IDENTIFIER {
		op, err := env.GetVar(nil, t2.Raw)
		if err {
			fmt.Fprintf(os.Stderr, "[line %d] Error: Uninitialized variable \"%s\"\n",
				t2.Line, t2.Raw)
			return token.Token{}, true
		}
		t2 = op
	}else if t2.IsBool() {
		t2.Value = t2.TokenType == token.TRUE
	}
	if !ok1 || (!ok2 && in_group.Operand2 != nil) {
		panic(fmt.Sprintf("operands should be tokens: operand1=%s, operand2=%s", in_group.Operand1, in_group.Operand2))
	}
	t := eval_token(in_group.Operator, t1, t2)
	if t.TokenType == token.None {
		return t, true
	}
	return t, false
}

func eval_token(operator token.Token, operand1 token.Token, operand2 token.Token) token.Token {

	if (operand1.TokenType != operand2.TokenType) && (!operand1.IsBool() || !operand2.IsBool()) &&
		operand2.TokenType != token.None {
		fmt.Fprintf(os.Stderr, "Cannot operate on %s and %s\n", operand1.TokenType, operand2.TokenType)
		return token.Token{}
	}

	res_type := token.None

	switch operator.TokenType {
	case token.EQUAL_EQUAL:{
		operand1.Value = operand1.Value == operand2.Value
		if operand1.Value.(bool) {
			res_type = token.TRUE
		}else {
			res_type = token.FALSE
		}
	}
	case token.BANG:{
		switch operand1.TokenType {
		case token.TRUE, token.FALSE: {
			operand1.Value = !operand1.Value.(bool)
			if operand1.Value.(bool) {
				res_type = token.TRUE
			}else {
				res_type = token.FALSE
			}
		}
		default:
			fmt.Fprintf(os.Stderr, "Cannot use operator \"%s\" on type %s\n",
				operator.Raw, operand1.TokenType)
			return token.Token{}
		}
	}
	case token.BANG_EQUAL:{
		operand1.Value = operand1.Value != operand2.Value
		if operand1.Value.(bool) {
			res_type = token.TRUE
		}else {
			res_type = token.FALSE
		}
	}
	case token.LESS:{
		switch operand1.TokenType {
		case token.NUMBER:
			operand1.Value = operand1.Value.(float64) < operand2.Value.(float64)
		default:
			fmt.Fprintf(os.Stderr, "Cannot use operator \"%s\" on type %s\n",
				operator.Raw, operand1.TokenType)
			return token.Token{}
		}
		if operand1.Value.(bool) {
			res_type = token.TRUE
		}else {
			res_type = token.FALSE
		}
	}
	case token.LESS_EQUAL:{
		switch operand1.TokenType {
		case token.NUMBER:
			operand1.Value = operand1.Value.(float64) <= operand2.Value.(float64)
		default:
			fmt.Fprintf(os.Stderr, "Cannot use operator \"%s\" on type %s\n",
				operator.Raw, operand1.TokenType)
			return token.Token{}
		}
		if operand1.Value.(bool) {
			res_type = token.TRUE
		}else {
			res_type = token.FALSE
		}
	}
	case token.GREATER:{
		switch operand1.TokenType {
		case token.NUMBER:
			operand1.Value = operand1.Value.(float64) > operand2.Value.(float64)
		default:
			fmt.Fprintf(os.Stderr, "Cannot use operator \"%s\" on type %s\n",
				operator.Raw, operand1.TokenType)
			return token.Token{}
		}
		if operand1.Value.(bool) {
			res_type = token.TRUE
		}else {
			res_type = token.FALSE
		}
	}
	case token.GREATER_EQUAL:{
		switch operand1.TokenType {
		case token.NUMBER:
			operand1.Value = operand1.Value.(float64) >= operand2.Value.(float64)
		default:
			fmt.Fprintf(os.Stderr, "Cannot use operator \"%s\" on type %s\n",
				operator.Raw, operand1.TokenType)
			return token.Token{}
		}
		if operand1.Value.(bool) {
			res_type = token.TRUE
		}else {
			res_type = token.FALSE
		}
	}
	case token.AND_AND:{
		switch operand1.TokenType {
		case token.TRUE, token.FALSE:
			operand1.Value = operand1.Value.(bool) && operand2.Value.(bool)
		default:
			fmt.Fprintf(os.Stderr, "Cannot use operator \"%s\" on type %s\n",
				operator.Raw, operand1.TokenType)
			return token.Token{}
		}
		if operand1.Value.(bool) {
			res_type = token.TRUE
		}else {
			res_type = token.FALSE
		}
	}
	case token.PIPE_PIPE:{
		switch operand1.TokenType {
		case token.TRUE, token.FALSE:
			operand1.Value = operand1.Value.(bool) || operand2.Value.(bool)
		default:
			fmt.Fprintf(os.Stderr, "Cannot use operator \"%s\" on type %s\n",
				operator.Raw, operand1.TokenType)
			return token.Token{}
		}
		if operand1.Value.(bool) {
			res_type = token.TRUE
		}else {
			res_type = token.FALSE
		}
	}
	case token.PLUS:{
		switch operand1.TokenType {
		case token.NUMBER:
			operand1.Value = operand1.Value.(float64) + operand2.Value.(float64)
			res_type = token.NUMBER
		case token.STRING:
			operand1.Value = operand1.Value.(string) + operand2.Value.(string)
			res_type = token.STRING
		default:
			fmt.Fprintf(os.Stderr, "Cannot use operator \"%s\" on type %s\n",
				operator.Raw, operand1.TokenType)
			return token.Token{}
		}
	}
	case token.MINUS:{
		switch operand1.TokenType {
		case token.NUMBER:
			operand1.Value = operand1.Value.(float64) - operand2.Value.(float64)
			res_type = token.NUMBER
		default:
			fmt.Fprintf(os.Stderr, "Cannot use operator \"%s\" on type %s\n",
				operator.Raw, operand1.TokenType)
			return token.Token{}
		}
	}
	case token.NEG:{
		switch operand1.TokenType {
		case token.NUMBER:
			operand1.Value = -1 * operand1.Value.(float64)
			res_type = token.NUMBER
		default:
			fmt.Fprintf(os.Stderr, "Cannot use operator \"%s\" on type %s\n",
				operator.Raw, operand1.TokenType)
			return token.Token{}
		}
	}
	case token.STAR:{
		switch operand1.TokenType {
		case token.NUMBER:
			operand1.Value = operand1.Value.(float64) * operand2.Value.(float64)
			res_type = token.NUMBER
		default:
			fmt.Fprintf(os.Stderr, "Cannot use operator \"%s\" on type %s\n",
				operator.Raw, operand1.TokenType)
			return token.Token{}
		}
	}
	case token.SLASH:{
		switch operand1.TokenType {
		case token.NUMBER:
			if operand2.Value.(float64) == 0.0 {
				fmt.Fprintf(os.Stderr, "Error: Cannot divide by zero\n")
				return token.Token{}
			}
			operand1.Value = operand1.Value.(float64) / operand2.Value.(float64)
			res_type = token.NUMBER
		default:
			fmt.Fprintf(os.Stderr, "Cannot use operator \"%s\" on type %s\n",
				operator.Raw, operand1.TokenType)
			return token.Token{}
		}
	}
	case token.PERCENT:{
		switch operand1.TokenType {
		case token.NUMBER:
			panic("not implemented")
			//operand1.Value = operand1.Value.(float64) % operand2.Value.(float64)
			//res_type = token.NUMBER
		default:
			fmt.Fprintf(os.Stderr, "Cannot use operator \"%s\" on type %s\n",
				operator.Raw, operand1.TokenType)
			return token.Token{}
		}
	}
	case token.CARET:{
		switch operand1.TokenType {
		case token.NUMBER:
			panic("not implemented")
			//v, ok := int(operand2.Value.(float64))
			//if !ok {
			//	//fmt.Fprintf(os.Stderr, "Error: Cannot divide by zero\n")
			//	return token.Token{}
			//}
			//operand1.Value = operand1.Value.(float64)  operand2.Value.(float64)
			//res_type = token.NUMBER
		default:
			fmt.Fprintf(os.Stderr, "Cannot use operator \"%s\" on type %s\n",
				operator.Raw, operand1.TokenType)
			return token.Token{}
		}
	}

	}
	if res_type == token.None {
		fmt.Fprintf(os.Stderr, "Result of operation %s %s %s is invalid\n", operand1, operator, operand2)
		return token.Token{}
	}
	operand1.TokenType = res_type
	if operator.IsUnary() {
		operand1.Raw = "(" + operator.Raw + operand1.Raw + ")"
	}else {
		operand1.Raw = "(" + operand1.Raw + operator.Raw + operand2.Raw + ")"
	}
	return operand1
}

