package token

import "fmt"

type TokenType int

const (
	L_BRACKET TokenType = iota //0
	R_BRACKET
	L_BRACE
	R_BRACE

	STAR // 4
	COMMA
	SEMICOLON
	PLUS
	MINUS
	SLASH
	CARET

	EQUAL // 11
	EQUAL_EQUAL
	BANG
	BANG_EQUAL
	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL

	AND_AND // 19
	PIPE_PIPE

	STRING // 21
	NUMBER

	NEWLINE
	WHITESPACE
	COMMENT

	IDENTIFIER // 26

	FALSE //27
	TRUE

	IF // 29
	ELSE
	FOR
	WHILE

	INIT // 33
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
