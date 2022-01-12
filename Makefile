SHELL := /bin/bash
.DEFAULT_GOAL := build
BUILDTIME=$(shell date +"%s")

.EXPORT_ALL_VARIABLES:
JWT_SECRET ?= cf40b7a444b3f2a53cafacc29ca5f672275274379eea842dfd
S3_SECRET_KEY ?= 54789f14abe263dc9c668a2e4c88f21d49cb6638b33e08c1904213f471ed6188
S3_ACCESS_KEY ?= 4a73fb9f-647e-4e88-b273-ae1a60dbb881
S3_ENDPOINT ?= https://s3.ir-thr-at1.arvanstorage.com
S3_BUCKET_NAME ?= mobydick-beta

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


deploy:
	docker build -t amirmfallah/mobydick-app-frontend:image-processor.0.0.${BUILDTIME} .
	docker push amirmfallah/mobydick-app-frontend:image-processor.0.0.${BUILDTIME}
	~/arvan paas set image deployment/mobydick-image-processor mobydick-image-processor=amirmfallah/mobydick-app-frontend:image-processor.0.0.${BUILDTIME}