GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOTOOL=$(GOCMD) tool
GOMOD=$(GOCMD) mod

CMD_PATH=cmd/*.go
BINARY_NAME=poc

start:
				$(GORUN) $(CMD_PATH)

build:
				$(GOBUILD) -o ./$(BINARY_NAME) $(CMD_PATH)

mocks:
				$(GOMOD) vendor
				$(GORUN) ./vendor/github.com/vektra/mockery/v2/main.go --all --dir ./internal --output ./tests/unit/mocks/ --keeptree --case underscore

unit-tests:mocks
				$(GOTEST) ./tests/unit

unit-tests-with-cover:mocks
				$(GOTEST) ./tests/unit -coverpkg ./internal/... -coverprofile=cover.out
				$(GOTOOL) cover -func=cover.out

unit-tests-with-cover-html:unit-tests-with-cover
				$(GOTOOL) cover -html=cover.out -o cover.html