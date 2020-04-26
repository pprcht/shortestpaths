#!/bin/make -f

#binary name
PROG := shortestpath

#Go commands
GOCMD := go
GOBUILD := $(GOCMD) build
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get

all: test build

build:
	$(GOBUILD) -o $(PROG) -v

test:
	$(GOTEST) -v 

clean:
	$(GOCMD) clean
	rm -f $(PROG)


