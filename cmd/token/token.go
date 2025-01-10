package token

import "fmt"

type TokenType int

const (
	None TokenType = iota

	L_BRACKET
	R_BRACKET
	L_BRACE
	R_BRACE

	SEMICOLON

	STAR
	COMMA
	PLUS
	MINUS
	SLASH
	CARET
	PERCENT

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

	NEWLINE
	WHITESPACE
	COMMENT
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

func (t Token)IsOperator() bool{
	return t.TokenType >= STAR && t.TokenType <= PIPE_PIPE
}

func (t Token)IsOperand() bool {
	return t.TokenType >= STRING && t.TokenType <= TRUE
}
