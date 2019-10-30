SOURCE_FILES?=./...

export PATH := ./bin:$(PATH)
export GO111MODULE := on
export GOPATH := $(shell go env GOPATH)
export GOPROXY := https://goproxy.io,direct
export GOPRIVATE := github.com/jinmukeji/*
export GOVERSION := $(shell go version | awk '{print $$3}')
# GORELEASER is the path to the goreleaser binary.
export GORELEASER := $(shell which goreleaser)

# all is the default target
all: release
.PHONY: all

# Install all the build and lint dependencies
setup:
	# TODO: 官方 golangci-lint 发行包不兼容 go 1.13，需要使用手动编译的版本
	#curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	[ -d "./bin" ] || mkdir -p ./bin
	cp $(GOPATH)/bin/golangci-lint ./bin/
	# curl -L https://git.io/misspell | sh
	go mod download
.PHONY: setup

# Update go packages
go-update:
	@echo "Updating go packages..."
	@go get -u all
	@echo "go mod tidy..."
	@$(MAKE) go-mod-tidy
.PHONY: go-update

# Clean go.mod
go-mod-tidy:
	@go mod tidy -v
	# @git --no-pager diff HEAD
	# @git --no-pager diff-index --quiet HEAD
.PHONY: go-mod-tidy

generate:
	@go generate ./...
.PHONY: generate

# Format go files
format:
	@goimports -w ./
.PHONY: format

# Run all the linters
lint:
	@./bin/golangci-lint run
.PHONY: lint

# Go build all
build:
	@go build ./... > /dev/null
.PHONY: build

# Go test all
test:
	@go test -v ./...
.PHONY: test

# Run all code checks
ci: generate format lint build test
.PHONY: ci

# Release wia goreleaser
release:
	@[ -x "$(GORELEASER)" ] || ( echo "goreleaser not installed"; exit 1)
	@goreleaser --snapshot --skip-publish --rm-dist
.PHONY: release

.DEFAULT_GOAL := all
