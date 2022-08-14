.PHONY: build test

build:
	go generate ./...
	go build -o vm_translator *.go

test:
	sh ./run_test.sh