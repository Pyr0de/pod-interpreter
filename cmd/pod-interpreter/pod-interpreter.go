package main

import (
	"fmt"
	"os"

	"github.com/Pyr0de/pod-interpreter/cmd/parser"
	"github.com/Pyr0de/pod-interpreter/cmd/scanner"
)

func Interpreter(t string, f string)int {

	switch t {
		case "tokenize": {
			t, err := scanner.Tokenize(string(f))
			for _,v := range t {
				v.Display()
			}
			fmt.Println("EOF  null")
			if err != nil {
				fmt.Fprintln(os.Stderr, "Tokenize Error")
				return 65
			}
			return 0
		}
		case "parse": {
			t, err := scanner.Tokenize(string(f))
			if err != nil {
				fmt.Fprintln(os.Stderr, "Tokenize Error")
				return 65
			}

			exp, err := parser.Parse(t)
			for _, v := range exp {
				fmt.Println(v)
			}

			if err != nil {
				fmt.Fprintln(os.Stderr, "Parser Error")
				return 65
			}
			return 0
		}
		default:
			return 1
	}
}
