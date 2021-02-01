GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOTOOL=$(GOCMD) tool
GOMOD=$(GOCMD) mod
GOVEND=$(GOMOD) vendor

CMD_PATH=cmd/*.go
BINARY_NAME=poc

start:
				$(GORUN) $(CMD_PATH)
build:
				$(GOBUILD) -o ./$(BINARY_NAME) $(CMD_PATH)

mocks:
				$(GOGET) github.com/vektra/mockery/v2/.../
				$(GOVEND)
				$(GORUN) ./vendor/github.com/vektra/mockery/v2/main.go --all --dir ./internal --output ./tests/unit/mocks/ --keeptree --case underscore

unit-tests:mocks
				$(GOTEST) ./tests/unit

cover:
				$(GOTEST) -v -coverpkg ./test -coverprofile=cover.out
				$(GOTOOL) cover -func=cover.out

cover-browser:
				$(GOTEST) -v -coverpkg ./internal/action,./internal/datasource,./internal/dispatcher,./internal/health,./internal/metric,./internal/metricsgroup,./internal/metricsgroupaction,./internal/moove,./internal/plugin ./internal/tests/ -coverprofile=cover.out
				$(GOTOOL) cover -html=cover.out -o cover.html
				open cover.html