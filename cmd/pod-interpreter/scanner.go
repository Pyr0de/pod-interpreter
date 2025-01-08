package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Pyr0de/pod-interpreter/cmd/token"
)

func Tokenize(input string) []token.Token {
	tokens := []token.Token {}
	//err := false

	var line uint = 1
	for i := 0; i < len(input); i++ {
		start := i
		c := input[i]

		if c >= '0' && c <= '9' {
			i++
			for (input[i] >= '0' && input[i] <= '9') || input[i] == '.' {
				i++
			}
			i--
			val, err := strconv.ParseFloat(input[start:i+1], 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[line %d] Improper Number: %s\n", line, input[start:i+1])
			}else {
				tokens = append(tokens, token.Token{TokenType: token.NUMBER, Raw: input[start:i+1], Value: val, Line: line})
			}
			continue
		}
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '_' {
			i++
			for (input[i]>='A' && input[i]<='Z') || (input[i]>='a' && input[i]<='z') || 
				(input[i]>='0' && input[i]<='9') || input[i]=='_' {
				i++
			}
			i--
			token_type := reserved(input[start:i+1])
			if token_type == token.IDENTIFIER {
				tokens = append(tokens, token.Token{TokenType: token.IDENTIFIER, Raw: input[start:i+1], Value: input[start:i+1], Line: line})
			}else {
				tokens = append(tokens, token.Token{TokenType: token_type, Raw: input[start:i+1], Line: line})
			}
			continue
		}

		switch c {
		case '(': 
			tokens = append(tokens, token.Token{TokenType: token.L_BRACKET, Raw: input[start:i+1], Line: line})
		case ')':
			tokens = append(tokens, token.Token{TokenType: token.R_BRACKET, Raw: input[start:i+1], Line: line})
		case '{': 
			tokens = append(tokens, token.Token{TokenType: token.L_BRACE, Raw: input[start:i+1], Line: line})
		case '}':
			tokens = append(tokens, token.Token{TokenType: token.R_BRACE, Raw: input[start:i+1], Line: line})
		case '+':
			tokens = append(tokens, token.Token{TokenType: token.PLUS, Raw: input[start:i+1], Line: line})
		case '-':
			tokens = append(tokens, token.Token{TokenType: token.MINUS, Raw: input[start:i+1], Line: line})
		case '*':
			tokens = append(tokens, token.Token{TokenType: token.STAR, Raw: input[start:i+1], Line: line})
		case '^':
			tokens = append(tokens, token.Token{TokenType: token.CARET, Raw: input[start:i+1], Line: line})
		case ',':
			tokens = append(tokens, token.Token{TokenType: token.COMMA, Raw: input[start:i+1], Line: line})
		case ';':
			tokens = append(tokens, token.Token{TokenType: token.SEMICOLON, Raw: input[start:i+1], Line: line})
			//case (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '_':
			//	for (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '_' {

			//	}
			//	tokens = append(tokens, token.Token{TokenType: token.IDENTIFIER, Raw: input[start:i], Line: line})
		case '=':
			switch input[i+1] {
			case '=':
				i++
				tokens = append(tokens, token.Token{TokenType: token.EQUAL_EQUAL, Raw: input[start:i+1], Line: line})
			default:
				tokens = append(tokens, token.Token{TokenType: token.EQUAL, Raw: input[start:i+1], Line: line})
			}
		case '!':
			switch input[i+1] {
			case '=':
				i++
				tokens = append(tokens, token.Token{TokenType: token.BANG_EQUAL, Raw: input[start:i+1], Line: line})
			default:
				tokens = append(tokens, token.Token{TokenType: token.BANG, Raw: input[start:i+1], Line: line})
			}
		case '>':
			switch input[i+1] {
			case '=':
				i++
				tokens = append(tokens, token.Token{TokenType: token.GREATER_EQUAL, Raw: input[start:i+1], Line: line})
			default:
				tokens = append(tokens, token.Token{TokenType: token.GREATER, Raw: input[start:i+1], Line: line})
			}
		case '<':
			switch input[i+1] {
			case '=':
				i++
				tokens = append(tokens, token.Token{TokenType: token.GREATER_EQUAL, Raw: input[start:i+1], Line: line})
			default:
				tokens = append(tokens, token.Token{TokenType: token.GREATER, Raw: input[start:i+1], Line: line})
			}
		case '/':
			switch input[i+1] {
			case '/':
				for input[i] != '\n' && i < len(input) {
					i++
				}
				i--
			default:
				tokens = append(tokens, token.Token{TokenType: token.SLASH, Raw: input[start:i+1], Line: line})
			}
		case '"', '\'': {
			i++
			for input[i] != c {
				if input[i] == '\n' {
					i--
					fmt.Fprintf(os.Stderr, "[line %d] Unterminated String: %s\n", line, input[start:i+1])
					break
				}
				i++
			}
			tokens = append(tokens, token.Token{TokenType: token.STRING, Raw: input[start:i+1], Value: input[start+1:i], Line: line})
		}
		case ' ', '\t':
		case '\n':
			line++
		default: {
			fmt.Fprintf(os.Stderr, "[line %d] Unexpected character: %s\n", line, input[start:i+1])
			//err = true
			}
		}
	}

	return tokens
}

func reserved(identifier string) token.TokenType {
	switch identifier {
	case "true":
		return token.TRUE
	case "false":
		return token.FALSE
	case "init":
		return token.INIT
	case "if":
		return token.IF
	case "else":
		return token.ELSE
	case "while":
		return token.WHILE
	case "for":
		return token.FOR
	case "func":
		return token.FUNC
	case "return":
		return token.RETURN
	}
	return token.IDENTIFIER
}
