package stmt

import (
	"fmt"

	"github.com/Pyr0de/pod-interpreter/cmd/env"
	"github.com/Pyr0de/pod-interpreter/cmd/eval"
	"github.com/Pyr0de/pod-interpreter/cmd/group"
	"github.com/Pyr0de/pod-interpreter/cmd/token"
)

func (s StmtPrint)Run() bool {
	v, err := eval.Evaluate(s.Expression)
	if err {
		return true
	}
	fmt.Println(v)
	return false
}

func (s StmtInit)Run() bool {
	val, ok := s.Expression.Operand2.(token.Token)
	if !ok {
		v, ok := s.Expression.Operand2.(*group.Group)
		if !ok {
			panic("InitRun: val is not token/group")
		}
		v_token, err := eval.Evaluate(*v)
		if err {
			return true
		}
		val = v_token
	}

	initVar(s.Expression.Operand1, val)

	return false
}

func initVar(in_group any, val token.Token) {
	if t, ok := in_group.(token.Token); ok {
		env.InitVar(t.Raw, val)
	}else {
		g, ok := in_group.(*group.Group)
		if !ok {
			panic("in_group is not token/group")
		}
		initVar(g.Operand1, val)
		initVar(g.Operand2, val)
	}
}
