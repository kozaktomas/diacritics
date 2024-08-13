#!/usr/bin/make -f
SRC_DIR	:= $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

all: deps lint test

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

test:
	go test -race ./...

.PHONY: deps
deps:
	go mod tidy && go mod verify
