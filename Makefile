-include .env

PROJECT_NAME=$(shell basename "$(PWD)")
# GO Variables
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor:$(GOBASE)/cmd/$(PROJECT_NAME):$(GOBASE)/pkg:$(GOBASE)/internal
GOBIN=$(GOBASE)/bin
GOFILES=$(shell find cmd/ -name "*.go")
TIMERS_PACKAGE_FILES=$(shell find pkg/timers -name "*.go")

TIMERS_MODULE_PATH=github.com/chukak/task-manager/pkg/timers
TEST_MODULE_PATH=github.com/chukak/task-manager/pkg/test
UTILITY_MODULE_PATH=github.com/chukak/task-manager/pkg/utility

MAKEFLAGS=-silent

build:
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECT_NAME) $(GOFILES)

test:
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test $(GOFILES)

test-package:
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test $(TIMERS_PACKAGE_FILES)

test-all: | test test-package

init-modules:
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get $(TIMERS_MODULE_PATH)
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get $(TEST_MODULE_PATH)
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get $(UTILITY_MODULE_PATH)

run: | build
	cd $(GOBIN) && ./$(PROJECT_NAME)

clean:
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean
