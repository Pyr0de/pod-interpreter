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

func (g *Group)AddGroup(child *Group) {
	if g.Operand1 == nil {
		g.Operand1 = child
	}else {
		g.Operand2 = child
	}
}

func (g *Group)CanAdd(t token.Token) bool{
	if t.IsOperand() {

		if g.Operand1 == nil {
			return true
		}else if g.Operator.TokenType != token.BANG && g.Operand2 == nil {
			return true
		}
	}
	if t.IsOperator() && g.Operator.TokenType == token.None{
		return true
	}

	
	return false
}

func (g *Group)AddToken(t token.Token) {

	if t.IsOperand() {

		if g.Operand1 == nil {
			g.Operand1 = t
		}else if g.Operator.TokenType != token.BANG && g.Operand2 == nil {
			g.Operand2 = t
		}
	}
	if t.IsOperator() && g.Operator.TokenType == token.None{
		g.Operator = t
	}
}

func (g Group)String() string {
	out := ""
	if g.Operator.TokenType == token.None {
		out += "group "
	}else {
		out += g.Operator.Raw + " "
	}

	val1, ok := g.Operand1.(*Group)
	if ok {
		out += val1.String()
	}else {
		if v, ok := g.Operand1.(token.Token); ok {
			out += v.Raw + " "
		}else {
			out += "nil "
		}
	}

	val2, ok := g.Operand2.(*Group)
	if ok {
		out += val2.String()
	}else {
		if v, ok := g.Operand2.(token.Token); ok {
			out += v.Raw + " "
		}else {
			out += "nil"
		}
	}


	return "(" + out + ") "
}

