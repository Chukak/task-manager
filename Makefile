-include .env

PROJECT_NAME=$(shell basename "$(PWD)")
# GO Variables
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor:$(GOBASE)/cmd/$(PROJECT_NAME):$(GOBASE)/pkg:$(GOBASE)/internal
GOBIN=$(GOBASE)/bin
GOFILES=$(shell find cmd/ -name "*.go")
GO_SOURCES=$(shell find internal/ -name "*.go")

TIMERS_PACKAGE_FILES=$(shell find pkg/timers -name "*.go")
CHECKS_PACKAGE_FILES=$(shell find pkg/checks -name "*.go")

TIMERS_MODULE_PATH=github.com/chukak/task-manager/pkg/timers
CHECKS_MODULE_PATH=github.com/chukak/task-manager/pkg/checks
UTILITY_MODULE_PATH=github.com/chukak/task-manager/pkg/utility

# other
GIN_MODULE_URL=github.com/gin-gonic/gin
GIN_STATIC_MODULE_URL=github.com/gin-contrib/static

GO_DEP_DIRECTORY=vendor/src/*
GO_DEPENDENCIES=$(TIMERS_MODULE_PATH) $(TEST_MODULE_PATH) $(UTILITY_MODULE_PATH) \
	$(GIN_STATIC_MODULE_URL) $(GIN_MODULE_URL)

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

test-src:
	@echo Running tests...
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test $(GO_SOURCES)

test-package:
	@echo Running package tests...
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test $(TIMERS_PACKAGE_FILES)
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test $(CHECKS_PACKAGE_FILES)

test-all: | test-src test-package

init-modules:
	@echo Update dependepcies... 
	@for dep in $(GO_DEPENDENCIES); do GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -u $${dep}; done

remove-modules:
	@echo Removing all dependencies from this project...
	rm -rf $(GO_DEP_DIRECTORY)

update-modules: | remove-modules init-modules

run: | build react-run
	cd $(GOBIN) && ./$(PROJECT_NAME) \

clean: | react-clean
	@echo Clean bin/$(PROJECT_NAME)
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean; \
	rm -f bin/$(PROJECT_NAME) \
