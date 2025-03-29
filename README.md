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
- python3
- bash

## Build
```sh
./run build
```
```sh
# Wasm
./run build wasm
```

## Usage
Native Executable can be found in `./out`
```sh
RUN_TYPE={interpreter-arg} ./run {script-args} {.pod file} 
# OR
./out/pod-interpreter {interpreter-arg} {.pod file}
```
Interpreter Arguments
- `tokenize`
- `parse` (parse expression)
- `evaluate` (evaluate expression)
- `run` (run program)

Script Arguments (not set by default)
- `build`: builds for selected target
- `wasm`: sets target to WASM
- `test`: Runs tests
- `bless`: (only with `test`) saves state of stdout & stderr for future testing

Web server starts at localhost:8080. Web interpreter files can be found in `./out`
```sh
# Wasm (server at localhost:8080)
./run wasm
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
- [X] Testing

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
`string`, `int`, `float`, `true/false`
### Other
`;`, `(`, `)`, `{`, `}`
