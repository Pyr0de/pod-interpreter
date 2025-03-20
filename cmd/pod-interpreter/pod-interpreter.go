package main

import (
	"fmt"
	"os"

	"github.com/Pyr0de/pod-interpreter/cmd/eval"
	"github.com/Pyr0de/pod-interpreter/cmd/parser"
	"github.com/Pyr0de/pod-interpreter/cmd/scanner"
)

func Interpreter(t string, f string) int {

	switch t {
	case "tokenize":
		{
			t, err := scanner.Tokenize(string(f))
			for _, v := range t {
				v.Display()
			}
			fmt.Println("EOF  null")
			if err != nil {
				fmt.Fprintln(os.Stderr, "Tokenize Error")
				return 65
			}
			return 0
		}
	case "parse":
		{
			t, err := scanner.Tokenize(string(f))
			if err != nil {
				fmt.Fprintln(os.Stderr, "Tokenize Error")
				return 65
			}

			exp, err := parser.ParseExpression(t)
			for _, v := range exp {
				fmt.Println(v)
			}

			if err != nil {
				fmt.Fprintln(os.Stderr, "Parser Error")
				return 65
			}
			return 0
		}
	case "evaluate":
		{
			t, err := scanner.Tokenize(string(f))
			if err != nil {
				fmt.Fprintln(os.Stderr, "Tokenize Error")
				return 65
			}

			exp, err := parser.ParseExpression(t)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Parser Error")
				return 65
			}
			e := false
			for _, v := range exp {
				out, err := eval.Evaluate(v)
				if err {
					e = err
				}
				fmt.Println(out)
			}
			if e {
				return 66
			}
			return 0
		}
	case "run":
		t, err := scanner.Tokenize(string(f))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Tokenize Error")
			return 65
		}
		s, e := parser.Parse(t)
		if e {
			fmt.Fprintln(os.Stderr, "Parser Error")
			return 65
		}
		for _, v := range s {
			e := v.Statement.Run()
			if e {
				fmt.Fprintf(os.Stderr, "Error\n")
				return 66
			}
		}
		return 0
	default:
		return 1
	}
}
