package main

import (
	"github.com/Pyr0de/pod-interpreter/cmd/group"
	"github.com/Pyr0de/pod-interpreter/cmd/token"
)


func Parse(tokens []token.Token) []group.Group {
	exp := []group.Group{}
	
	var start_group group.Group = nil
	var curr_group *group.Group = nil

	for _,k := range tokens {

	}
	return exp
}
