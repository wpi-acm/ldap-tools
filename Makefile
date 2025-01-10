GO ?= go
PREFIX ?= /usr/local
BINDIR ?= bin

commands := finduser

all: $(commands)

finduser:
	$(GO) build -o . ./cmd/finduser