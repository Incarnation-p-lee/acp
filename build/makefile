SHELL                            :=/bin/bash
BINARY_NAME                      :=acp
SRC_FILES                        :=$(shell find . -type f -name '*.go' -print)

.PHONY: help test build

help:
	@echo Usage
	@echo
	@echo "	make [command]"
	@echo
	@echo " The commands are:"
	@echo
	@echo "	help:    print this help message."
	@echo "	build:   build this repository."
	@echo "	test:    run test code."
	@echo

build: $(SRC_FILES)
	gofmt -w $^
	go build $^

test: $(SRC_FILES)
	go test -v -coverprofile $$$$.cover.out -covermode=atomic $^
