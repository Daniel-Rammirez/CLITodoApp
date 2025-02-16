.DEFAULT_GOAL := install

.PHONY: fmt vet build install
fmt: 
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build
install: build
	go install