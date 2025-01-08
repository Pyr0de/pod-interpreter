package token

import "fmt"

type TokenType int

const (
	L_BRACKET TokenType = iota //0
	R_BRACKET
	L_BRACE
	R_BRACE

	STAR
	COMMA
	SEMICOLON
	PLUS
	MINUS
	SLASH
	CARET

	EQUAL
	EQUAL_EQUAL
	BANG
	BANG_EQUAL
	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL

	AND_AND
	PIPE_PIPE

	STRING
	NUMBER

	NEWLINE
	WHITESPACE
	COMMENT

	IDENTIFIER

	FALSE
	TRUE

	IF
	ELSE
	FOR
	WHILE

	INIT
	FUNC
	RETURN
)

type Token struct {
	TokenType TokenType
	Raw string
	Value interface{}
	Line uint
}

func (t Token)Display() {
	fmt.Println(t.TokenType, t.Raw, t.Value, t.Line)
}
