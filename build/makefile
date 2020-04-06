SHELL                            :=/bin/bash
BINARY_NAME                      :=acp
SRC_FILES                        :=$(shell find . -type f -name '*.go' -print)
FAKE_FILES                       :=$(subst .go,.fake,$(SRC_FILES))
FAKE_TEST_FILES                  :=$(subst .go,.fake_test,$(SRC_FILES))

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

build: $(FAKE_FILES)
$(FAKE_FILES):%.fake:%.go
	gofmt -w $<
	go build $<

test: $(FAKE_TEST_FILES)
$(FAKE_TEST_FILES):%.fake_test:%.go
	go test -v -coverprofile $$$$.cover.out -covermode=atomic $<
