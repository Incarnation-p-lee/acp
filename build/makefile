SHELL                :=/bin/bash
BINARY_NAME          :=acp
PACKAGES             :=
SRC_FILES            :=$(shell find . -type f -name '*.go' -print)
PACKAGES_STRIP       =$(strip $(PACKAGES))

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

build:
	@gofmt -w $(SRC_FILES)
	go build $(PACKAGES_STRIP)

test: 
	@gofmt -w $(SRC_FILES)
	go test -v -coverprofile $$$$.cover.out -covermode=atomic $(PACKAGES_STRIP)
	@cat *.cover.out > cover.out && rm *.cover.out
