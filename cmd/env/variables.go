package env

import "github.com/Pyr0de/pod-interpreter/cmd/token"


var env *environment = &environment{store: make(map[string]token.Token), global: nil}

func InitVar(variable string, val token.Token) {
	env.store[variable] = val
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

type environment struct {
	store map[string]token.Token
	global *environment
}

func findVar(variable string) *environment {
	curr := env

	for curr != nil {
		
		if _, ok := curr.store[variable]; ok {
			return curr
		}

		curr = curr.global
	}
	return nil
}
