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
	PERCENT

	EQUAL // 12
	EQUAL_EQUAL
	BANG
	BANG_EQUAL
	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL

	AND_AND // 20
	PIPE_PIPE

	STRING // 22
	NUMBER

	NEWLINE
	WHITESPACE
	COMMENT

	IDENTIFIER // 27

	FALSE //28
	TRUE

	IF // 30
	ELSE
	FOR
	WHILE

	INIT // 34
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
