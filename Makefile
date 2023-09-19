# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GORUN = $(GOCMD) run
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GOVET = $(GOCMD) vet
GOMOD = $(GOCMD) mod
GOFMT = $(GOCMD) fmt

# Build target
BINARY_NAME = spacetraders
BUILD_DIR = build

include .env
$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' .env))

.PHONY: help confirm all build clean run test test/cover deps vet cleanmod

# HELPERS

## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' Makefile | column -t -s ':' |  sed -e 's/^/ /'

## confirm: confirm helper
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

## env: print env variables
env:
	@echo $(shell cat .env)

# DEVELOPMENT

## all: clean and build
all: clean build

## build: go build
build:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v

## play: run the binary
play: build
	$(BUILD_DIR)/$(BINARY_NAME) play

## clean: clean directory
clean:
	$(GOFMT) ./...
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

## run: execute go run
run:
	$(GORUN) .

## test: run all tests
test:
	$(GOTEST) -v ./...

## test/cover: run all tests and display coverage
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## deps: get dependencies
deps:
	$(GOGET) ./...
	$(GOMOD) tidy
	$(GOMOD) vendor

## vet: run go vet
vet:
	$(GOVET) ./...

## cleanmod: clean go.mod
cleanmod:
	$(GOCLEAN) -modcache -cache