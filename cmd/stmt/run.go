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

func (s StmtExpression)Run() bool {
	val, err := eval.Evaluate(s.Expression)
	fmt.Printf("[line %d] >> %s\n", val.Line, val.String())
	return err
}

func (s StmtBlock)Run() bool {
	env.NextScope()
	for _, v := range s.Block {
		e := v.Statement.Run()
		if e {
			fmt.Fprintf(os.Stderr, "Error\n")
			return true
		}
	}
	env.PrevScope()
	return false
}

func (s StmtIf)Run() bool {
	v, err := eval.Evaluate(s.Expression)
	b, ok := v.Value.(bool)
	if err || !v.IsBool() || !ok {
		fmt.Fprintf(os.Stderr, "[line %d] Error: Expected boolean expression\n", )
		return true
	}
	if b {
		return s.Block.Run()
	}else {
		if s.Else != nil {
			return s.Else.Run()
		}
	}
	return false
}

func (s StmtWhile)Run() bool {
	for {
		v, err := eval.Evaluate(s.Expression)
		b, ok := v.Value.(bool)
		if err || !v.IsBool() || !ok {
			fmt.Fprintf(os.Stderr, "[line %d] Error: Expected boolean expression\n", )
			return true
		}
		if !b {
			break
		}
		e := s.Block.Run()
		if e {
			return err
		}
	}

	return false
}

func (_ StmtEmpty)Run() bool {
	return false
}
