package eval

import (
	"fmt"
	"os"

	"github.com/Pyr0de/pod-interpreter/cmd/env"
	"github.com/Pyr0de/pod-interpreter/cmd/group"
	"github.com/Pyr0de/pod-interpreter/cmd/token"
)

func Evaluate(in_group group.Group) (token.Token, bool) {
	if group1, ok := in_group.Operand1.(*group.Group); ok {
		g, err := Evaluate(*group1)
		if err {
			return token.Token{}, true
		}
		in_group.Operand1 = g
	} else if _, ok := in_group.Operand1.(token.Token); !ok {
		panic("Error evaluating operand 1")
	}

	if !in_group.Operator.IsUnary() && in_group.Operator.TokenType != token.None {

		if group2, ok := in_group.Operand2.(*group.Group); ok {
			g, err := Evaluate(*group2)
			if err {
				return token.Token{}, true
			}
			in_group.Operand2 = g
		} else if _, ok := in_group.Operand2.(token.Token); !ok {
			panic("Error evaluating operand 2")
		}
	}

	t1, ok1 := in_group.Operand1.(token.Token)
	if t1.TokenType == token.IDENTIFIER {
		op, err := env.GetVar(t1.Raw)
		if err {
			fmt.Fprintf(os.Stderr, "[line %d] Error: Uninitialized variable \"%s\"\n",
				t1.Line, t1.Raw)
			return token.Token{}, true
		}
		t1 = op
	} else if t1.IsBool() {
		t1.Value = t1.TokenType == token.TRUE
	}
	if in_group.Operator.TokenType == token.None {
		return t1, false
	}

	t2, ok2 := in_group.Operand2.(token.Token)
	if t2.TokenType == token.IDENTIFIER {
		op, err := env.GetVar(t2.Raw)
		if err {
			fmt.Fprintf(os.Stderr, "[line %d] Error: Uninitialized variable \"%s\"\n",
				t2.Line, t2.Raw)
			return token.Token{}, true
		}
		t2 = op
	} else if t2.IsBool() {
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
