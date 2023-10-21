
.PHONY: wasm
wasm:
	GOOS=js GOARCH=wasm go build -o public/main.wasm

.PHONY: install-tools
install-tools:
	go install github.com/cosmtrek/air@latest

.PHONY: devserver
devserver:
	go run $(CURDIR)/devserver & air -c .air.toml