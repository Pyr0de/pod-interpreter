//go:build wasm

package main

import (
	"syscall/js"
)

func run(this js.Value, p []js.Value) any {

	t := p[0].String()
	file := p[1].String()

	err := Interpreter(t, file)
	return js.ValueOf(err)
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("run", js.FuncOf(run))

	<-c
}
