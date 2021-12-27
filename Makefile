# Go parameters
GOCMD=go
TAG_NAME:=$(shell git describe --abbrev=0 --tags)

all: test

test: 
	$(GOCMD) test -v ./...
