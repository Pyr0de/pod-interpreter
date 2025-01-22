# Pod Interpreter
Pod is a dynamically typed language, similar to Python and Javascript
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

## Execute
Native Executable can be found in `./bin`
```sh
make run
```
Web interpreter can be found in `./bin/wasm`
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
- [ ] For loop
- [ ] Functions
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
