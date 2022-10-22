# Go parameters
GOCMD=go

all: test

vendor: FORCE
	$(GOCMD) mod vendor

tidy: FORCE
	$(GOCMD) mod tidy

test: FORCE
	$(GOCMD) test -v ./...

FORCE: ;
