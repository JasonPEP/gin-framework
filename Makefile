## load .env
include $(PWD)/Makefile.env
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
	go test -v ./...
######################################
## run app
.PHONY: run
run: build
	@./$(BINARY_NAME)
######################################
## build for linux/amd64
.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) -v
######################################
## build for windows/amd64
.PHONY: build-windows
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME) -v
