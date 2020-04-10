-include .env

PROJECT_NAME=$(shell basename "$(PWD)")
# GO Variables
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor:$(GOBASE)/cmd/$(PROJECT_NAME):$(GOBASE)/pkg
GOBIN=$(GOBASE)/bin
GOFILES=$(shell find . -name "*.go")

MAKEFLAGS=-silent

build:
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECT_NAME) $(GOFILES)

test:
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test

run: | build
	cd $(GOBIN) && ./$(PROJECT_NAME)

clean:
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean


