.PHONY: lint
lint:
	staticcheck -f stylish ./...

.PHONY: gen
gen:
	go run ./cmd/codegen gen -m $(module) -t $(target)