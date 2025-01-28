package env

import (
	"fmt"
	"os"

	"github.com/Pyr0de/pod-interpreter/cmd/token"
)

func InitFunc(env *environment, name string, token_pointer_to_func token.Token) bool {
	if token_pointer_to_func.TokenType != token.FUNC {
		panic(fmt.Sprintln("Expected token to be func with address to StmtFunc, got", token_pointer_to_func))
	}
	e := findFunc(env, name) 
	if e != nil {
		fmt.Fprintf(os.Stderr, "[line %d] Error: \"%s\" is already declared cannot redeclare\n")
		return true
	}else {
		e = curr_env
	}
	e.functions[name] = token_pointer_to_func

	return false
}

func findFunc(env *environment, name string) *environment {
	if env == nil {
		env = curr_env
	}
	curr := env
	for curr != nil {
		if _, ok := curr.functions[name]; ok {
			return curr
		}
		curr = curr.global
	}
	return nil
}
