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
				sh build-plugins.sh
				$(GORUN) $(CMD_PATH)
build:
				$(GOBUILD) -o $(DIST_PATH)/$(BINARY_NAME) $(CMD_PATH)

mocks:
				$(GOGET) github.com/vektra/mockery@v1.1.2
				$(GOVEND)
				$(GORUN) $(GOPATH)/mod/github.com/vektra/mockery@v1.1.2/cmd/mockery/mockery.go --all --dir ./internal/user --output ./tests/unit/mocks/ --keeptree --case underscore
tests:
				make tests-dep-up
				$(GOTEST) ./...
cover:
				$(GOTEST) -v -coverpkg ./internal/action,./internal/datasource,./internal/dispatcher,./internal/health,./internal/metric,./internal/metricsgroup,./internal/metricsgroupaction,./internal/moove,./internal/plugin ./internal/tests/ -coverprofile=cover.out
				$(GOTOOL) cover -func=cover.out
cover-browser:
				$(GOTEST) -v -coverpkg ./internal/action,./internal/datasource,./internal/dispatcher,./internal/health,./internal/metric,./internal/metricsgroup,./internal/metricsgroupaction,./internal/moove,./internal/plugin ./internal/tests/ -coverprofile=cover.out
				$(GOTOOL) cover -html=cover.out -o cover.html
				open cover.html