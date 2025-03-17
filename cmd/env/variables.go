package env

import (
	"fmt"
	"os"

	"github.com/Pyr0de/pod-interpreter/cmd/token"
)

var curr_env *Environment = NewEnv()

type Environment struct {
	store map[string]token.Token
	Functions map[string]FunctionEntry
	Global *Environment
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
		curr = curr.Global
	}
	return nil
}

func NextScope() {
	c := NewEnv()
	c.Global = curr_env
	curr_env = c
}

func PrevScope() {
	curr_env = curr_env.Global
}

func SwapEnv(swap *Environment) *Environment {
	e := curr_env
	curr_env = swap
	return e
}

func NewEnv() *Environment {
	return &Environment{
		store: make(map[string]token.Token),
		Functions: make(map[string]FunctionEntry),
		Global: nil,
	}
}

func CurrEnv() *Environment {
	return curr_env
}
