package env

import (
	"fmt"
	"os"

	"github.com/Pyr0de/pod-interpreter/cmd/token"
)

func InitFunc(name string, token_pointer_to_func token.Token) bool {
	if token_pointer_to_func.TokenType != token.FUNC {
		panic(fmt.Sprintln("Expected token to be func with address to StmtFunc, got", token_pointer_to_func))
	}
	e := FindFunc(name) 
	if e != nil {
		fmt.Println(name, token_pointer_to_func, e)
		fmt.Fprintf(os.Stderr, "[line %d] Error: \"%s\" is already declared cannot redeclare\n", token_pointer_to_func.Line, name)
		return true
	}else {
		e = curr_env
	}
	e.Functions[name] = token_pointer_to_func

	return false
}

func FindFunc(name string) *Environment {
	curr := curr_env
	for curr != nil {
		if _, ok := curr.Functions[name]; ok {
			return curr
		}
		curr = curr.global
	}
	return nil
}
