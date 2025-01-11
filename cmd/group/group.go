package group

import (
	"github.com/Pyr0de/pod-interpreter/cmd/token"
)

type Group struct {
	Operator token.Token
	Operand1 any
	Operand2 any
	Parent *Group
}


func (g Group)String() string {
	out := ""
	if g.Operator.TokenType == token.None {
		out += "group"
	}else {
		out += g.Operator.Raw
	}

	val1, ok := g.Operand1.(*Group)
	if ok {
		out += " " + val1.String()
	}else {
		if v, ok := g.Operand1.(token.Token); ok {
			out += " " + v.Raw
		}else {
			out += " nil"
		}
	}

	val2, ok := g.Operand2.(*Group)
	if ok {
		out += " " + val2.String()
	}else {
		if v, ok := g.Operand2.(token.Token); ok {
			out += " " + v.Raw
		}else {
			out += " nil"
		}
	}


	return "(" + out + ")"
}

