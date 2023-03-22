APP_NAME := mvn-snapshot-cleaner
APP_VERSION := 1.2
BUILD_TIME := $(shell date "+%F %T %Z")
PWD := $(shell pwd)
OUTDIR ?= $(PWD)/dist
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
GO_BUILD := go build -ldflags "-s -w" -ldflags "-X 'main.appName=$(APP_NAME)' -X 'main.appVersion=$(APP_VERSION)' -X 'main.buildTime=$(BUILD_TIME)'"

.PHONY: build clean rebuild linux/amd64 linux/arm64 linux/arm windows/amd64 darwin/amd64 darwin/arm64 linux windows darwin buildall

build:
	mkdir -p $(OUTDIR)/ && $(GO_BUILD) -o $(OUTDIR)/$(APP_NAME)_$(GOOS)_$(GOARCH)$(GOEXE) .

rebuild: clean build

linux/amd64: export GOOS=linux
linux/amd64: export GOARCH=amd64
linux/amd64: build

linux/arm64: export GOOS=linux
linux/arm64: export GOARCH=arm64
linux/arm64: build

linux/arm: export GOOS=linux
linux/arm: export GOARCH=arm
linux/arm: build

windows/amd64: export GOOS=windows
windows/amd64: export GOARCH=amd64
windows/amd64: export GOEXE=.exe
windows/amd64: build

darwin/amd64: export GOOS=darwin
darwin/amd64: export GOARCH=amd64
darwin/amd64: build

darwin/arm64: export GOOS=darwin
darwin/arm64: export GOARCH=arm64
darwin/arm64: build

linux:
	make linux/amd64
	make linux/arm64
	make linux/arm

windows:
	make windows/amd64

darwin:
	make darwin/amd64
	make darwin/arm64

buildall: linux windows darwin

rebuildall: clean buildall

clean:
	rm -rf $(OUTDIR)