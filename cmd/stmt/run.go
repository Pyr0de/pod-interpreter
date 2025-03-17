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
			panic(fmt.Sprintln("Expected operand2 to be group/token, found: ", s.Expression.Operand2))
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

// in_group is all the '=' together in a group
// example: a = b = c = d = 10;
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
		fmt.Fprintf(os.Stderr, "[line %d] Error: Expected boolean expression\n", v.Line)
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
			fmt.Fprintf(os.Stderr, "[line %d] Error: Expected boolean expression\n", v.Line)
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

func (s StmtFor)Run() bool {
	e := false
	if assign, ok := s.Initialization.Statement.(StmtAssign); ok {
		e = assign.Run()
	}
	if e {
		return e
	}
	condition := s.Condition
	if condition.Empty() {
		condition = group.Group{Operand1: token.Token{TokenType: token.TRUE, Value: true}}
	}
	step, step_ok := s.Step.Statement.(StmtAssign)

	for {
		v, err := eval.Evaluate(condition)
		b, ok := v.Value.(bool)
		if err || !v.IsBool() || !ok {
			fmt.Fprintf(os.Stderr, "[line %d] Error: Expected boolean expression\n", v.Line)
			return true
		}
		if !b {
			break
		}
		e := s.Block.Run()
		if e {
			return err
		}
		if step_ok {
			e = step.Run()
		}
		if e {
			return err
		}
	}
	return false
}

func (s StmtFunc)Run() bool {
	return env.InitFunc(s.Name.Raw, token.Token{
		TokenType: token.FUNC, Value: &s, Line: s.Name.Line,
	})

}

func (s StmtFuncCall)Run() bool {
	e := env.FindFunc(s.Name.Raw)
	if e == nil {
		// function not found
		fmt.Fprintf(os.Stderr, "[line %d] Error: Uninitialized function \"%s\"\n", s.Name.Line, s.Name.Raw)
		return true
	}

	func_entry := e.Functions[s.Name.Raw]
	f, ok := func_entry.TokenPointerToFunc.Value.(*StmtFunc)
	if !ok {
		panic("Found something other than stmtfunc")
	}
	if len(f.Parameters) != len(s.Parameters) {
		fmt.Printf(
			"[line %d] Error: Mismatched number of parameters "+
			"expected %d parameters, found %d parameters\n",
			s.Name.Line, len(f.Parameters), len(s.Parameters),
		)
		return true
	}
	args_vals := []token.Token{}
	for i := range len(s.Parameters) {
		val, err := eval.Evaluate(s.Parameters[i])
		if err {
			return true
		}
		args_vals = append(args_vals, val)
	}

	new_env := env.NewEnv()
	new_env.Global = func_entry.Env
	
	prev_env := env.SwapEnv(new_env)

	for i := range len(s.Parameters) {
		initVar(f.Parameters[i], args_vals[i], true)
	}

	f.Block.Run()
	
	env.SwapEnv(prev_env)
	return false
}

func (_ StmtEmpty)Run() bool {
	return false
}
