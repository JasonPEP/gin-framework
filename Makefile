## load .env
include $(PWD)/Makefile.env
GO ?= go
TESTFOLDER := $(shell $(GO) list ./... | grep -E 'gin$$|binding$$|render$$' | grep -v examples)

## install dep
.PHONY: install
install:
	echo 'instaling'
######################################
## defualt build for current system
.PHONY: build
build: install
	go build -o $(BINARY_NAME) -v
######################################
## run test
.PHONY: test
test: install
	$(GO) test ./... -v -race -coverprofile=profile.json -covermode=atomic
######################################
## run app
.PHONY: run
run: build
	@./$(BINARY_NAME)
######################################
## build for linux/amd64
.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) -v
######################################
## build for windows/amd64
.PHONY: build-windows
build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME) -v
