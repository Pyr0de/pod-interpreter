package stmt

import (
	"fmt"
	"os"

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
	fmt.Println(v.String())
	return false
}

func (s StmtAssign)Run() bool {
	g, ok := s.Expression.Operand2.(*group.Group)
	if !ok {
		t, ok := s.Expression.Operand2.(token.Token)
		if !ok {
			panic("InitRun: val is not token/group")
		}
		g = &group.Group{Operand1: t}
	}

	val, err := eval.Evaluate(*g)
	if err {
		return true
	}

	initVar(s.Expression.Operand1, val, s.Init)

	return false
}

func initVar(in_group any, val token.Token, init bool) {
	if t, ok := in_group.(token.Token); ok {
		if init {
			env.InitVar(t.Raw, val)
		}else {
			err := env.SetVar(t.Raw, val)
			if err {
				fmt.Fprintf(os.Stderr, "[line %d] Error: Uninitialized variable \"%s\"\n", t.Line, t.Raw)
			}
		}
	}else {
		g, ok := in_group.(*group.Group)
		if !ok {
			panic("in_group is not token/group")
		}
		initVar(g.Operand1, val, init)
		initVar(g.Operand2, val, init)
	}
}
