.PHONY: fmt example

fmt:
	go fmt ./...

example:
	go run ./example/main.go

