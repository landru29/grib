SHELL := /bin/bash

# The name of the executable (default is current directory name)
TARGET := grib
.DEFAULT_GOAL: $(TARGET)

# These will be provided to the target
VERSION := 1.0.0
BUILD := `git rev-parse HEAD`

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Build=$(BUILD)"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: all build clean install uninstall fmt simplify check run test-all

all: test-all install
	@echo "Built $(TARGET), git sha $(BUILD)"

$(TARGET): $(SRC)
	go build $(LDFLAGS) -o $(TARGET)

build: $(TARGET)
	@true

clean:
	rm -f $(TARGET)
	rm *.png

fmt:
	gofmt -l -w $(SRC)

test:
	go test $(go list ./... | grep -v /mocks) -short ./...

lint:
	go vet ./...

test-all: lint test

gen:
	go generate ./...

strict-check:
	@test -z $(shell gofmt -l main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done
	@go tool vet ${SRC}

run: test-all install
	@$(TARGET)