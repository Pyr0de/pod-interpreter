package group

import "github.com/Pyr0de/pod-interpreter/cmd/token"

type Group struct {
	operator token.Token
	operand1 any
	operand2 any
	parent *Group
}

func (g Group)String() string {
	return ""
}

