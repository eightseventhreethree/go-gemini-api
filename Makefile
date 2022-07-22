# Go parameters
GOCMD=go
#TAG_NAME:=$(shell git describe --abbrev=0 --tags)

all: gen test

vendor: FORCE
	$(GOCMD) mod vendor

tidy: FORCE
	$(GOCMD) mod tidy

gen: FORCE
	$(GOCMD) generate ./...

test: FORCE
	$(GOCMD) test -v ./...

FORCE: ;
