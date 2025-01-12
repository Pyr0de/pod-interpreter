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
	if g.Operand2 == nil && g.Operator.TokenType == token.None && g.Operand1 != nil && g.Parent == nil{
		if v, ok := g.Operand1.(token.Token); ok {
			return v.String()
		}
	}

	out := ""
	if g.Operator.TokenType == token.None {
		out += "group"
	}else {
		out += g.Operator.Raw
	}

	if v, ok := g.Operand1.(*Group); ok {
		out += " " + v.String()
	}else if v, ok := g.Operand1.(token.Token); ok {
		out += " " + v.String()
	}else {
		out += " null"
	}

	if v, ok := g.Operand2.(*Group); ok {
		out += " " + v.String()
	}else if v, ok := g.Operand2.(token.Token); ok {
		out += " " + v.String()
	}else {
		out += " null"
	}

	return "(" + out + ")"
}

func (g Group)Empty() bool{
	return g.Operand1 == nil && g.Operand2 == nil && g.Operator.TokenType == 0
}
