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
	$(GOTEST) fibheap -v
	$(GOTEST) -v 

# utilize Go's native benchmark/test system
bench:
	$(GOTEST) fibheap -v -bench=.
	$(GOTEST) -v -bench=.

clean:
	$(GOCMD) clean
	rm -f $(PROG)


