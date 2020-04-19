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

# other
GIN_MODULE_URL=github.com/gin-gonic/gin
GIN_STATIC_MODULE_URL=github.com/gin-contrib/static

GODEP_INTERNAL=$(TIMERS_MODULE_PATH) $(TEST_MODULE_PATH) $(UTILITY_MODULE_PATH)
GODEP_EXTERNAL=$(GIN_STATIC_MODULE_URL) $(GIN_MODULE_URL)

MAKEFLAGS=-silent

REACT_BIN_DIRECTORY=$(shell pwd)/bin/web/
REACT_LOG=$(REACT_BIN_DIRECTORY)/react.log

export CURRENT_SOURCE_PATH=$(shell pwd)/

react-prepare:
	@echo Preparing react directories...
	cd $(REACT_BIN_DIRECTORY)/; \
	rm -rf ./public ./src; \
	cp -r $(CURRENT_SOURCE_PATH)web/* ./; \
	mv ./templates ./public; \

react-build: | react-prepare
	@echo Installing react modules...
	cd $(REACT_BIN_DIRECTORY)/ && npm install && npm run build

react-start:
	@echo Start client on http://localhost:3000
	cd $(REACT_BIN_DIRECTORY)/ && npm start

react-clean:
	@echo Clean 'bin/web'
	cd $(REACT_BIN_DIRECTORY)/; \
	rm -rf build public src node_modules package.json *.log; \

react-run:
	@echo Running react...
	cd $(REACT_BIN_DIRECTORY)/; \
	npm start > react.log 2>&1 & \

# todo
react-stop: 
	@echo Stopping react...
	cd $(REACT_BIN_DIRECTORY)/; \
	npm stop | true \

build: | react-build
	@echo Building golang modules...
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECT_NAME) $(GOFILES)

test:
	@echo Running tests...
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test $(GOFILES)

test-package:
	@echo Running package tests...
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test $(TIMERS_PACKAGE_FILES)

test-all: | test test-package

init-modules:
	@echo Update dependepcies... 
	$(foreach dep, $(GODEP_INTERNAL), $(shell GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get $(dep))); \
	$(foreach dep, $(GODEP_EXTERNAL), $(shell GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -u $(dep))) 

run: | build react-run
	cd $(GOBIN) && ./$(PROJECT_NAME) \

clean: | react-clean
	@echo Clean bin/$(PROJECT_NAME)
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean; \
	rm -f bin/$(PROJECT_NAME) \
