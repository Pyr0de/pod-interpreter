package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		switch os.Args[1] {
			case "tokenize": {
				f, err := os.ReadFile(os.Args[2])
				if err != nil {
					fmt.Fprintln(os.Stderr, "Could not read file:", os.Args[2])
					os.Exit(65)
				}
				t, err := Tokenize(string(f))
				for _,v := range t {
					v.Display()
				}
				fmt.Println("EOF  null")
				if err != nil {
					os.Exit(65)
				}
				return
			}
			case "parse": {
				f, err := os.ReadFile(os.Args[2])
				if err != nil {
					fmt.Fprintln(os.Stderr, "Could not read file:", os.Args[2])
					os.Exit(65)
				}
				t, err := Tokenize(string(f))
				if err != nil {
					os.Exit(65)
				}

				for _, v := range Parse(t) {
					fmt.Println(v)
				}
				return
			}
		}

	}
	fmt.Println("Usage:", os.Args[0], "{tokenize/parse} {filename}")
}
