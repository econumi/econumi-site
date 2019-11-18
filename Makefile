# Overview of Make: https://www.gnu.org/software/make/manual/html_node/index.html#SEC_Contents

SHELL=/bin/sh
CGO_ENABLED=0
GOOS=linux

SRC_DIR:=$(shell pwd)
BIN_DIR:=$(SRC_DIR)/bin
DATA_DIR:=$(SRC_DIR)/data
GO_PATH=$(SRC_DIR)

all: build

# Build binaries

build: $(BIN_DIR)/coin

$(BIN_DIR)/coin:
	go build -o bin/econumi econumi.org/coin

clean:
	@echo cleaning up bin directory
	@rm -rf $(BIN_DIR)/*
	@rm -rf $(DATA_DIR)/*

rm-data:
	@rm -rf $(DATA_DIR)/*

loc:
	find . -name "*.go" -not -path "./src/econumi.org/vendor/*" | xargs wc -l

PHONY: all build clean
