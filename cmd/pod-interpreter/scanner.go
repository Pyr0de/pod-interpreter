package main

import (
	"fmt"
	"os"

	"github.com/Pyr0de/pod-interpreter/cmd/token"
)

func Tokenize(input string) []token.Token {
	tokens := []token.Token {}
	line := 1
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '(': 
			tokens = append(tokens, token.Token{TokenType: token.L_BRACKET, Raw: "("})
		case ')':
			tokens = append(tokens, token.Token{TokenType: token.R_BRACKET, Raw: ")"})
		default: {
			fmt.Fprintf(os.Stderr, "[line %d] Unexpected character: %s\n", line, string(input[i]))
		}
		}
	}

	return tokens
}
