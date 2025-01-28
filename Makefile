.PHONY: build build-wasm exec run run-wasm cp-examples

build:
	@go build -o ./bin/pod-interpreter ./cmd/pod-interpreter/
	@echo "Native Build"

build-wasm:
	@GOOS=js GOARCH=wasm go build -o ./bin/web/main.wasm ./cmd/pod-interpreter
	@cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" ./bin/web/
	@cp ./web/* ./bin/web/
	@rm ./bin/web/*.go
	@echo "Wasm Build"

exec:
	@echo "Running..."
	./bin/pod-interpreter $(ARGS)

cp-examples:
	@echo "Copying Examples..."
	@python3 ./examples/example-index.py
	@cp -r ./examples ./bin/web
	@rm ./bin/web/examples/*.py
	@mv ./bin/example_index.json ./bin/web/examples

run: build exec
run-wasm: build-wasm cp-examples
	@go build -o ./bin/pod-interpreter ./web/main.go
	@$(MAKE) exec
