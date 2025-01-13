//go:build !wasm
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Building Native")
	e := 1
	if len(os.Args) > 2 {
		f, err := os.ReadFile(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not read file:", os.Args[2])
			os.Exit(1)
		}
		e = Interpreter(os.Args[1],string(f))
	}
	if e != 0 {
		fmt.Println("Usage:", os.Args[0], "{tokenize/parse} {filename}")
	}
	os.Exit(e)
}
