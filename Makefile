GO ?= go
PREFIX ?= /usr/local
BINDIR ?= bin

commands := finduser

all: clean $(commands)

finduser:
	$(GO) build -o . ./cmd/finduser


clean:
	rm -f $(commands)
