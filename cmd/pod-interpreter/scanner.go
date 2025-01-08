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
		start := i
		c := input[i]

		switch c {
		case '(': 
			tokens = append(tokens, token.Token{TokenType: token.L_BRACKET, Raw: input[start:i+1]})
		case ')':
			tokens = append(tokens, token.Token{TokenType: token.R_BRACKET, Raw: input[start:i+1]})
		case '+':
			tokens = append(tokens, token.Token{TokenType: token.PLUS, Raw: input[start:i+1]})
		case '-':
			tokens = append(tokens, token.Token{TokenType: token.MINUS, Raw: input[start:i+1]})
		case '*':
			tokens = append(tokens, token.Token{TokenType: token.STAR, Raw: input[start:i+1]})
		case ',':
			tokens = append(tokens, token.Token{TokenType: token.COMMA, Raw: input[start:i+1]})
		case ';':
			tokens = append(tokens, token.Token{TokenType: token.SEMICOLON, Raw: input[start:i+1]})
		//case (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '_':
		//	for (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '_' {

		//	}
		//	tokens = append(tokens, token.Token{TokenType: token.IDENTIFIER, Raw: input[start:i]})
		case '\n':
			line++
		default: {
			fmt.Fprintf(os.Stderr, "[line %d] Unexpected character: %s\n", line, input[start:i+1])
		}
		}
	}

	return tokens
}
