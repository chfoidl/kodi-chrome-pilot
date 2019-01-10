# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=kodi-chrome-pilot

all: build-linux

build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./build/linux-amd64/bin/$(BINARY_NAME) -v .