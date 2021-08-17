SHELL := /bin/bash
.DEFAULT_GOAL := build

.EXPORT_ALL_VARIABLES:
JWT_SECRET ?= cf40b7a444b3f2a53cafacc29ca5f672275274379eea842dfd
S3_SECRET_KEY ?= e980d0963c1cefc7460563aae7e034f1bdec9daec0499842bfb96500e2a4daa9
S3_ACCESS_KEY ?= 26ac7d43-f24b-488e-b00e-41a4a1342477
S3_ENDPOINT ?= https://s3.ir-thr-at1.arvanstorage.com
S3_BUCKET_NAME ?= dokkaan-dev

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
