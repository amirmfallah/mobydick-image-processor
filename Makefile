SHELL := /bin/bash
.DEFAULT_GOAL := build

.PHONY: build clean tidy run

tidy:
	# To update and prune the dependencies
	cd src/lib; go mod tidy
	cd src/mux/muxlib; go mod tidy
	cd src/mux; go mod tidy

build:
	export GO111MODULE=on
	cd src/mux; env GOOS=linux go build -ldflags="-s -w" -o ../../bin/mux main.go

test:
	cd src/mux; go build -o ../../bin/mux-local main.go;
	cd ./bin; ./mux-local;
