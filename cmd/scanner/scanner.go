package scanner

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/Pyr0de/pod-interpreter/cmd/token"
)

func Tokenize(input string) ([]token.Token, error) {
	input += "\n"
	tokens := []token.Token{}
	err := false

	var line uint = 1
	for i := 0; i < len(input); i++ {
		start := i
		c := input[i]
		is_int := true

		if c >= '0' && c <= '9' {
			i++
			for (input[i] >= '0' && input[i] <= '9') || input[i] == '.' {
				if input[i] == '.' {
					is_int = false
				}
				i++
			}
			i--
			if is_int {
				int, err := strconv.ParseInt(input[start:i+1], 10, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "[line %d] Improper Int: %s\n", line, input[start:i+1])
					continue
				}
				tokens = append(tokens, token.Token{TokenType: token.INT, Raw: input[start : i+1], Value: int, Line: line})
			} else {
				float, err := strconv.ParseFloat(input[start:i+1], 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "[line %d] Improper Float: %s\n", line, input[start:i+1])
					continue
				}
				tokens = append(tokens, token.Token{TokenType: token.FLOAT, Raw: input[start : i+1], Value: float, Line: line})
			}
			continue
		}
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '_' {
			i++
			for (input[i] >= 'A' && input[i] <= 'Z') || (input[i] >= 'a' && input[i] <= 'z') ||
				(input[i] >= '0' && input[i] <= '9') || input[i] == '_' {
				i++
			}
			i--
			tokens = append(tokens, token.Token{TokenType: reserved(input[start : i+1]), Raw: input[start : i+1], Line: line})
			continue
		}

		switch c {
		case '(':
			tokens = append(tokens, token.Token{TokenType: token.L_BRACKET, Raw: input[start : i+1], Line: line})
		case ')':
			tokens = append(tokens, token.Token{TokenType: token.R_BRACKET, Raw: input[start : i+1], Line: line})
		case '{':
			tokens = append(tokens, token.Token{TokenType: token.L_BRACE, Raw: input[start : i+1], Line: line})
		case '}':
			tokens = append(tokens, token.Token{TokenType: token.R_BRACE, Raw: input[start : i+1], Line: line})
		case '+':
			tokens = append(tokens, token.Token{TokenType: token.PLUS, Raw: input[start : i+1], Line: line})
		case '-':
			{
				if len(tokens) == 0 || !tokens[len(tokens)-1].IsOperand() {
					tokens = append(tokens, token.Token{TokenType: token.NEG, Raw: input[start : i+1], Line: line})
				} else {
					tokens = append(tokens, token.Token{TokenType: token.MINUS, Raw: input[start : i+1], Line: line})
				}
			}
		case '*':
			tokens = append(tokens, token.Token{TokenType: token.STAR, Raw: input[start : i+1], Line: line})
		case '^':
			tokens = append(tokens, token.Token{TokenType: token.CARET, Raw: input[start : i+1], Line: line})
		case '%':
			tokens = append(tokens, token.Token{TokenType: token.PERCENT, Raw: input[start : i+1], Line: line})
		case ',':
			tokens = append(tokens, token.Token{TokenType: token.COMMA, Raw: input[start : i+1], Line: line})
		case ';':
			tokens = append(tokens, token.Token{TokenType: token.SEMICOLON, Raw: input[start : i+1], Line: line})
		case '=':
			switch input[i+1] {
			case '=':
				i++
				tokens = append(tokens, token.Token{TokenType: token.EQUAL_EQUAL, Raw: input[start : i+1], Line: line})
			default:
				tokens = append(tokens, token.Token{TokenType: token.EQUAL, Raw: input[start : i+1], Line: line})
			}
		case '!':
			switch input[i+1] {
			case '=':
				i++
				tokens = append(tokens, token.Token{TokenType: token.BANG_EQUAL, Raw: input[start : i+1], Line: line})
			default:
				tokens = append(tokens, token.Token{TokenType: token.BANG, Raw: input[start : i+1], Line: line})
			}
		case '>':
			switch input[i+1] {
			case '=':
				i++
				tokens = append(tokens, token.Token{TokenType: token.GREATER_EQUAL, Raw: input[start : i+1], Line: line})
			default:
				tokens = append(tokens, token.Token{TokenType: token.GREATER, Raw: input[start : i+1], Line: line})
			}
		case '<':
			switch input[i+1] {
			case '=':
				i++
				tokens = append(tokens, token.Token{TokenType: token.LESS_EQUAL, Raw: input[start : i+1], Line: line})
			default:
				tokens = append(tokens, token.Token{TokenType: token.LESS, Raw: input[start : i+1], Line: line})
			}
		case '/':
			switch input[i+1] {
			case '/':
				for input[i] != '\n' && i < len(input) {
					i++
				}
				i--
			default:
				tokens = append(tokens, token.Token{TokenType: token.SLASH, Raw: input[start : i+1], Line: line})
			}
		case '|', '&':
			{
				if c == input[i+1] {
					i++
					if c == '&' {
						tokens = append(tokens, token.Token{TokenType: token.AND_AND, Raw: input[start : i+1], Line: line})
					} else {
						tokens = append(tokens, token.Token{TokenType: token.PIPE_PIPE, Raw: input[start : i+1], Line: line})
					}
				} else {
					fmt.Fprintf(os.Stderr, "[line %d] Unexpected character: %s\n", line, input[start:i+1])
					err = true
				}
			}
		case '"', '\'':
			{
				i++
				for input[i] != c {
					if input[i] == '\n' {
						i--
						fmt.Fprintf(os.Stderr, "[line %d] Unterminated String: %s\n", line, input[start:i+1])
						err = true
						break
					}
					i++
				}
				tokens = append(tokens, token.Token{TokenType: token.STRING, Raw: input[start : i+1], Value: input[start+1 : i], Line: line})
			}
		case ' ', '\t':
		case '\n':
			line++
		default:
			{
				fmt.Fprintf(os.Stderr, "[line %d] Unexpected character: %s\n", line, input[start:i+1])
				err = true
			}
		}
	}

	if err {
		return tokens, errors.New("Error Tokenizer")
	}
	return tokens, nil
}

func reserved(identifier string) token.TokenType {
	switch identifier {
	case "true":
		return token.TRUE
	case "false":
		return token.FALSE
	case "let":
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
	case "print":
		return token.PRINT
	}
	return token.IDENTIFIER
}
