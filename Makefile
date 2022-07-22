# Go parameters
GOCMD=go

all: gen test

vendor: FORCE
	$(GOCMD) mod vendor

tidy: FORCE
	$(GOCMD) mod tidy

test: FORCE
	$(GOCMD) test -v ./...

FORCE: ;
