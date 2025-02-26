# [Pod Interpreter](https://pyr0de.github.io/pod-interpreter/)
Written in Go!
### Example:
```
let a = 10;
let b = 12;
if a == 10 {
    print a;
}else if b == 12 {
    print a+b;
}else {
    print 0;
}

```

## Prerequsites
- go (>=1.22.2)
## Build
```sh
make build
```
```sh
# Wasm
make build-wasm
```

## Usage
Native Executable can be found in `./bin`
```sh
./bin/podinterpreter {arguments} {.pod file}
```
Arguments
- `tokenize`
- `parse` (parse expression)
- `evaluate` (evaluate expression)
- `run` (run program)

Web server starts at localhost:8080. Web interpreter files can be found in `./bin/wasm`
```sh
# Wasm (server at localhost:8080)
make run-wasm
```

## To-Do
- [X] Tokenizer
- [X] Parser
- [X] Expression Evaluator
- [X] Variable
- [X] Print
- [X] If statement
- [X] While loop
- [X] For loop
- [X] Functions
- [X] Function Arguments
- [ ] Return
- [ ] File read/write
- [ ] Stdin
- [ ] Stdlib

## References
- [Crafting Interpreters](https://craftinginterpreters.com/)

### Operators
`+`, `-`, `*`, `/`, `^`, `%`, `==`, `!=`, `>`, `>=`, `<`, `<=`, `&&`, `||`
### Function
`func`, `return`
### Initialize
`let`
### If and loops
`if`, `else`, `for`, `while`
### Data types
`string`, `number`, `true/false`
### Other
`;`, `(`, `)`, `{`, `}`
