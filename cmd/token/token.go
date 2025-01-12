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
	COMMA

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

	PLUS
	MINUS
	NEG

	STAR
	SLASH
	PERCENT

	CARET

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
	return t.TokenType >= EQUAL && t.TokenType <= CARET
}

func (t Token)IsOperand() bool {
	return t.TokenType >= STRING && t.TokenType <= TRUE
}

func (t Token)IsUnary() bool {
	return t.TokenType == BANG || t.TokenType == NEG 
}
