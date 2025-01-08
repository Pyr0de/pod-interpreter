package token

import "fmt"

type TokenType int

const (
	L_BRACKET TokenType = iota //0
	R_BRACKET
	L_BRACE
	R_BRACE

	STAR // 4
	DOT
	COMMA
	PLUS
	MINUS
	SLASH

	EQUAL // 10
	EQUAL_EQUAL
	BANG
	BANG_EQUAL
	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL

	AND_AND // 18
	PIPE_PIPE

	STRING // 20
	NUMBER

	NEWLINE // 22
	WHITESPACE
	COMMENT

	IDENTIFIER // 25

	FALSE // 26
	TRUE

	IF // 28
	ELSE
	FOR
	WHILE

	INIT // 32
	FUNC
	RETURN // 34
)

type Token struct {
	TokenType TokenType
	Raw string
	Value interface{}
}

func (t Token)Display() {
	fmt.Println(t.TokenType, t.Raw, t.Value)
}
