package env

import (
	"fmt"
	"os"

	"github.com/Pyr0de/pod-interpreter/cmd/token"
)

var curr_env *Environment = NewEnv()

type Environment struct {
	store map[string]token.Token
	Functions map[string]token.Token
	global *Environment
}


func InitVar(variable string, val token.Token) bool {
	if findVar(variable) == nil{
		curr_env.store[variable] = val
		return false
	}
	fmt.Fprintf(os.Stderr, "Error: Variable \"%s\" already exists cannot reinitialize\n", variable)
	return true
}

func SetVar(variable string, val token.Token) bool {
	e := findVar(variable)
	if e != nil {
		e.store[variable] = val
		return false
	}
	return true
}

func DestructVar(variable string) {
	e := findVar(variable)
	if e != nil {
		delete(e.store, variable)
	}
}

func GetVar(variable string) (token.Token, bool) {
	e := findVar(variable)
	if e != nil {
		t, _ := e.store[variable]
		return t, false
	}
	return token.Token{}, true
}

func findVar(variable string) *Environment {
	curr := curr_env
	for curr != nil {
		if _, ok := curr.store[variable]; ok {
			return curr
		}
		curr = curr.global
	}
	return nil
}

func NextScope() {
	c := NewEnv()
	c.global = curr_env
	curr_env = c
}

func PrevScope() {
	curr_env = curr_env.global
}

func SwapEnv(swap *Environment) *Environment {
	e := curr_env
	curr_env = swap
	return e
}

func NewEnv() *Environment {
	return &Environment{
		store: make(map[string]token.Token),
		Functions: make(map[string]token.Token),
		global: nil,
	}
}
