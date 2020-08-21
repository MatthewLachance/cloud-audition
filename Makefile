# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
BINARY_NAME=cloudaudition
BINARY_UNIX=$(BINARY_NAME)_unix
LINTER=golangci-lint

export GO111MODULE=on

all: setup test build

setup:
		$(GOMOD) download
		$(GOMOD) tidy

build: clean
		$(GOBUILD) -o build/bin/$(BINARY_NAME) -v main.go

clean:
		$(GOCLEAN)
		rm -f build/bin/$(BINARY_NAME)
		rm -f build/bin/$(BINARY_UNIX)
test:
		rm -f coverage.out
		$(GOTEST) ./... -coverprofile=coverage.out

show-coverage: test
		$(GOCMD) tool cover -html=coverage.out

build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o build/bin/$(BINARY_UNIX) -v

build-image:
		docker build -t cloudaudition:test .

lint-local:
		docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:v1.30.0 $(LINTER) run -v

lint:
		$(LINTER) run

swag:
		swag init -g main.go

.PHONY: all build clean setup build-linux test show-coverage build-image lint-local lint swag