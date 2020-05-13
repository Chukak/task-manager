-include .env

PROJECT_NAME=$(shell basename "$(PWD)")
# GO Variables
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor:$(GOBASE)/cmd/$(PROJECT_NAME):$(GOBASE)/pkg:$(GOBASE)/internal
GOBIN=$(GOBASE)/bin
GOFILES=$(shell find cmd/ -name "*.go")
GO_SOURCES=$(shell find internal/ -name "*.go")
GO_SOURCE_DIR=$(GOBASE)/internal

TIMERS_PACKAGE_FILES=$(shell find pkg/timers -name "*.go")
CHECKS_PACKAGE_FILES=$(shell find pkg/checks -name "*.go")

TIMERS_MODULE_PATH=github.com/chukak/task-manager/pkg/timers
CHECKS_MODULE_PATH=github.com/chukak/task-manager/pkg/checks
UTILITY_MODULE_PATH=github.com/chukak/task-manager/pkg/utility

# other
GIN_MODULE_URL=github.com/gin-gonic/gin
GIN_STATIC_MODULE_URL=github.com/gin-contrib/static
PGX_MODULE_URL=github.com/jackc/pgx
PGXPUDDLE_MODULE_URL=github.com/jackc/puddle

GO_DEP_DIRECTORY=vendor/src/*
GO_DEPENDENCIES=$(TIMERS_MODULE_PATH) $(TEST_MODULE_PATH) $(UTILITY_MODULE_PATH) \
	$(GIN_STATIC_MODULE_URL) $(GIN_MODULE_URL) $(PGX_MODULE_URL) $(PGXPUDDLE_MODULE_URL) 

REACT_BIN_DIRECTORY=$(shell pwd)/bin/web/
REACT_LOG=$(REACT_BIN_DIRECTORY)/react.log

SQL_FILES_DIRECTORY=$(shell pwd)/db

export CURRENT_SOURCE_PATH=$(shell pwd)/
# exported variables
export DB_HOST=${TEST_PGHOST}
export DB_PORT=${TEST_PGPORT}
export DB_NAME=${TEST_DATABASE}
export DB_USER=${TEST_PGUSER}
export DB_PASSWORD=${TEST_PGPASSWORD}

react-prepare: | react-clean
	@echo Preparing react directories...
	cd $(REACT_BIN_DIRECTORY); \
	cp -r $(CURRENT_SOURCE_PATH)web/* ./; \
	cp $(CURRENT_SOURCE_PATH)web/.babelrc ./; \
	mv ./templates ./public; \

react-build: | react-prepare
	@echo Installing react modules...
	cd $(REACT_BIN_DIRECTORY) && yarn install && yarn run build

react-start:
	@echo Start client on http://localhost:3000
	cd $(REACT_BIN_DIRECTORY) && npm start

react-clean:
	@echo Clean 'bin/web'
	cd $(REACT_BIN_DIRECTORY); \
	rm -rf build public src node_modules dist *.json *.log *.lock *.js; \

react-run:
	@echo Running react...
	cd $(REACT_BIN_DIRECTORY); \
	npm start > react.log 2>&1 & \

# todo
react-stop: 
	@echo Stopping react...
	cd $(REACT_BIN_DIRECTORY); \
	npm stop | true \

build: | react-build
	@echo Building golang modules...
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -v $(LDFLAGS) -o $(GOBIN)/$(PROJECT_NAME) $(GOFILES)

test-src: | remove-db prepare-db
	@echo Running tests...
	cd $(GO_SOURCE_DIR); \
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -v ./...

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

prepare-db: 
	@echo Preparing test database...
	sudo -H -u postgres psql < $(SQL_FILES_DIRECTORY)/roles.sql
	PGPASSWORD=$(DB_PASSWORD) psql -U $(DB_USER) -d $(DB_NAME) < $(SQL_FILES_DIRECTORY)/schema.sql

remove-db: 
	@echo Removing test database...
	sudo -H -u postgres psql < $(SQL_FILES_DIRECTORY)/clear.sql

clean: | react-clean remove-test-db
	@echo Clean bin/$(PROJECT_NAME)
	GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean; \
	rm -f bin/$(PROJECT_NAME) \
