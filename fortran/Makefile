#!/bin/make -f

#binary name
PROG := shortestpath

#source files
SRCS := main.f90
SRCS += graphf.f90
SRCS += dijkstra.f90
SRCS += floyd-warshall.f90

#create object files
OBJS := $(patsubst %, build/%.o, $(SRCS))

# configuration
FC := gfortran
LD := $(FC)
RM := rm -rf
BUILD := build
vpath %.f90 .:lib
FFLAGS := -O -ffree-form -ffree-line-length-none
FFLAGS += -J$(BUILD)
LIBS := 

.PHONY: all
all: build $(PROG)

$(PROG): $(OBJS) 
	$(FC) $^ $(LIBS) -o $@

$(OBJS): $(BUILD)/%.o: %
	$(FC) $(FFLAGS) -c $< -o $@

# dependencies
$(BUILD)/main.f90.o: $(BUILD)/graphf.f90.o\
                     $(BUILD)/dijkstra.f90.o\
                     $(BUILD)/floyd-warshall.f90.o
$(BUILD)/dijkstra.f90.o: $(BUILD)/graphf.f90.o
$(BUILD)/floyd-warshall.f90.o: $(BUILD)/graphf.f90.o

.PHONY: build clean 
# make a build directory
build: $(BUILD) 
$(BUILD):
	mkdir -p $@

clean:
	$(RM) $(wildcard $(BUILD)/*.o) $(wildcard $(BUILD)/*.mod) $(PROG) $(BUILD)