-include .env

PROJECT_NAME=$(shell basename "$(PWD)")
# GO Variables
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor:$(GOBASE)/cmd/$(PROJECT_NAME):$(GOBASE)/pkg:$(GOBASE)/internal
GOBIN=$(GOBASE)/bin
GOFILES=$(shell find cmd/ -name "*.go")
PACKAGES_FILES=$(shell find pkg/ -name "*.go")

MAKEFLAGS=-silent

build:
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECT_NAME) $(GOFILES)

test:
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test $(GOFILES)

test-package:
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test $(PACKAGES_FILES)

test-all: | test test-package

run: | build
	cd $(GOBIN) && ./$(PROJECT_NAME)

clean:
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean


