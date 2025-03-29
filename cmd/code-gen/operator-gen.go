package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const OPERATOR_FILE = "cmd/eval/generated-operators.go"

func OperatorGen() (string, error) {
	operators := []string{
		case_operator("token.LESS", case_number("<")),
		case_operator("token.LESS_EQUAL", case_number("<=")),
		case_operator("token.GREATER", case_number(">")),
		case_operator("token.GREATER_EQUAL", case_number(">=")),
		case_operator("token.AND_AND", case_bool("&&")),
		case_operator("token.PIPE_PIPE", case_bool("||")),
		case_operator("token.PLUS", case_number("+")+case_string("+")),
		case_operator("token.MINUS", case_number("-")),
		case_operator("token.STAR", case_number("*")),
		pre_case_operator("token.SLASH", case_number("/"), divide_zero()),
		pre_case_operator("token.PERCENT", case_int("%"), divide_zero()),
	}

	template := "cmd/code-gen/operator.go.code"
	template_code, err := os.ReadFile(template)
	if err != nil {
		return "", errors.New("Failed to open: " + template)
	}

	gen_code := strings.Join(operators, "")

	return strings.Replace(string(template_code), "//[[code-gen]]", gen_code, 1), nil
}

func default_case() string {
	return `default: {
					fmt.Fprintf(os.Stderr, "Cannot operator on \"%s %s %s\"\n",
					operand1.Raw, operator.Raw, operand2.Raw)
					return token.Token{}
				}
`
}
func pre_case_operator(token string, cases string, pre string) string {
	return fmt.Sprintf(
	`case %s: {%s
		switch operand1.TokenType {
		%s
		%s
		}
	}
`,token, pre, cases, default_case())
}

func case_operator(token string, cases string) string {
	return pre_case_operator(token, cases, "")
}

func case_number(op string) string {
	return fmt.Sprintf(
			`case token.FLOAT:
				switch operand2.TokenType {
				case token.FLOAT:
					operand1.Value = operand1.Value.(float64) %s operand2.Value.(float64)
				case token.INT:
					operand1.Value = operand1.Value.(float64) %s float64(operand2.Value.(int64))
				%s
				}
			case token.INT:
				switch operand2.TokenType {
				case token.FLOAT:
					operand1.Value = float64(operand1.Value.(int64)) %s operand2.Value.(float64)
				case token.INT:
					operand1.Value = operand1.Value.(int64) %s operand2.Value.(int64)
				%s
				}
`, op, op, default_case(), op, op, default_case())
}

func case_bool(op string) string {
	return fmt.Sprintf(
`		case token.TRUE, token.FALSE:
			operand1.Value = operand1.Value.(bool) %s operand2.Value.(bool)
`, op)
}

func case_string(op string) string {
	return fmt.Sprintf(
`			case token.STRING:
				operand1.Value = operand1.Value.(string) %s operand2.Value.(string)
`, op)
}

func case_int(op string) string {
	return fmt.Sprintf(
`			case token.INT:
				switch operand2.TokenType {
					case token.INT:
						operand1.Value = operand1.Value.(int64) %s operand2.Value.(int64)
					%s
				}
`, op, default_case())
}

func divide_zero() string {
	return `
		switch operand2.TokenType {
		case token.INT:
			if operand2.Value.(int64) == 0 {
				fmt.Fprintf(os.Stderr, "[line %d] Error: Cannot divide by zero\n",
				operand2.Line)
				return token.Token{}
			}
		case token.FLOAT:
			if operand2.Value.(float64) == 0 {
				fmt.Fprintf(os.Stderr, "[line %d] Error: Cannot divide by zero\n",
				operand2.Line)
				return token.Token{}
			}
		}
	`
}
