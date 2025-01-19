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
	PRINT

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

func (t Token)String() string {
	if t.TokenType == None {
		return t.TokenType.String()
	}else if t.Value != nil {
		return fmt.Sprint(t.Value)
	}else {
		return t.Raw
	}
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

func (t Token)IsBool() bool {
	return t.TokenType == TRUE || t.TokenType == FALSE
}

func (t TokenType)String() string{
	switch t {
	case None:
		return "None"
	case L_BRACKET:
		return "LEFT_PAREN"
	case R_BRACKET:
		return "RIGHT_PAREN"
	case L_BRACE:
		return "LEFT_BRACE"
	case R_BRACE:
		return "RIGHT_BRACE"
	case SEMICOLON:
		return "SEMICOLON"
	case COMMA:
		return "COMMA"
	case EQUAL:
		return "EQUAL"
	case EQUAL_EQUAL:
		return "EQUAL_EQUAL"
	case BANG:
		return "BANG"
	case BANG_EQUAL:
		return "BANG_EQUAL"
	case LESS:
		return "LESS"
	case LESS_EQUAL:
		return "LESS_EQUAL"
	case GREATER:
		return "GREATER"
	case GREATER_EQUAL:
		return "GREATER_EQUAL"
	case AND_AND:
		return "AND_AND"
	case PIPE_PIPE:
		return "PIPE_PIPE"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case NEG:
		return "NEG"
	case STAR:
		return "STAR"
	case SLASH:
		return "SLASH"
	case PERCENT:
		return "PERCENT"
	case CARET:
		return "CARET"
	case STRING:
		return "STRING"
	case NUMBER:
		return "NUMBER"
	case IDENTIFIER:
		return "IDENTIFIER"
	case FALSE:
		return "FALSE"
	case TRUE:
		return "TRUE"
	case IF:
		return "IF"
	case ELSE:
		return "ELSE"
	case FOR:
		return "FOR"
	case WHILE:
		return "WHILE"
	case INIT:
		return "INIT"
	case FUNC:
		return "FUNC"
	case RETURN:
		return "RETURN"
	case PRINT:
		return "PRINT"
	}
	return ""
}

