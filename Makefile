.PHONY: build run

build:
	@go build -o ./bin/pod-interpreter ./cmd/pod-interpreter/
	@echo "Native Build"

build-wasm:
	@GOOS=js GOARCH=wasm go build -o ./bin/web/main.wasm ./cmd/pod-interpreter
	@cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./bin/web/
	@cp ./web/* ./bin/web/
	@rm ./bin/web/*.go
	@echo "Wasm Build"
	@go build -o ./bin/pod-interpreter ./web/main.go
	@echo "Build Server"

exec:
	@echo "Running..."
	./bin/pod-interpreter $(ARGS)

run: build exec
run-wasm: build-wasm exec
