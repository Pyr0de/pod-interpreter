package main

import (
	"fmt"

	"github.com/Pyr0de/pod-interpreter/cmd/group"
	"github.com/Pyr0de/pod-interpreter/cmd/token"
)


func Parse(tokens []token.Token) []group.Group {
	exp := []group.Group{}
	
	var start_group *group.Group = nil
	var curr_group *group.Group = nil

	for _,k := range tokens {
		switch k.TokenType {
		case token.L_BRACKET: {
			if curr_group != nil {
				g := &group.Group{Parent: curr_group}
				curr_group.AddGroup(g)
				curr_group = g
			}else {
				start_group = &group.Group{}
				curr_group = start_group
			}
		}
		case token.R_BRACKET: {
			if curr_group != nil {
				if curr_group.Parent != nil {
					curr_group = curr_group.Parent
				}else {
					exp = append(exp, *start_group)
					curr_group = nil
					start_group = nil
				}
			}else {
				panic(fmt.Sprintf("[line %d] Invalid bracket", k.Line))
			}
		}
		default: {
			if curr_group != nil {
				if curr_group.CanAdd(k) {
					curr_group.AddToken(k)
				}
			}
		}
		}
	}
	return exp
}

