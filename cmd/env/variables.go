package env

import (
	"fmt"
	"os"

	"github.com/Pyr0de/pod-interpreter/cmd/token"
)

var curr_env *environment = &environment{
	store: make(map[string]token.Token),
	functions: make(map[string]token.Token),
	global: nil,
}

type environment struct {
	store map[string]token.Token
	functions map[string]token.Token
	global *environment
}


func InitVar(env *environment, variable string, val token.Token) bool {
	if findVar(env, variable) == nil{
		env.store[variable] = val
		return false
	}
	fmt.Fprintf(os.Stderr, "Error: Variable \"%s\" already exists cannot reinitialize\n", variable)
	return true
}

func SetVar(env *environment, variable string, val token.Token) bool {
	e := findVar(env, variable)
	if e != nil {
		e.store[variable] = val
		return false
	}
	return true
}

func DestructVar(env *environment, variable string) {
	e := findVar(env, variable)
	if e != nil {
		delete(e.store, variable)
	}
}

func GetVar(env *environment, variable string) (token.Token, bool) {
	e := findVar(env, variable)
	if e != nil {
		t, _ := e.store[variable]
		return t, false
	}
	return token.Token{}, true
}

func findVar(env *environment, variable string) *environment {
	if env == nil {
		env = curr_env
	}
	curr := env
	for curr != nil {
		if _, ok := curr.store[variable]; ok {
			return curr
		}
		curr = curr.global
	}
	return nil
}

func NextScope(env *environment) *environment {
	if env == nil {
		env = curr_env
	}
	return &environment{
		store: make(map[string]token.Token),
		functions: make(map[string]token.Token),
		global: env,
	}
}

func PrevScope(env *environment) *environment {
	if env == nil {
		env = curr_env
	}
	return env.global
}
