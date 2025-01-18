package env

import "github.com/Pyr0de/pod-interpreter/cmd/token"


var store map[string]token.Token = make(map[string]token.Token)

func InitVar(variable string, val token.Token) {
	store[variable] = val
}

func SetVar(variable string, val token.Token) bool {
	if _, ok := store[variable]; ok {
		store[variable] = val
		return false
	}
	return true
}

func DestructVar(variable string) {
	if _, ok := store[variable]; ok {
		delete(store, variable)
	}
}

func GetVar(variable string) (token.Token, bool) {
	if v, ok := store[variable]; ok {
		return v, false
	}
	return token.Token{}, true
}
