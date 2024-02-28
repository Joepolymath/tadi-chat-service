# Makefile for Go project

# Variables
APP_NAME := app
SRC_DIR := .
BUILD_DIR := ./build
BIN_DIR := bin
PKG_DIR := pkg
MAIN_FILE := ./cmd/main.go

# Commands
GO := go
GO_BUILD := $(GO) build
GO_RUN := $(GO) run
GO_CLEAN := $(GO) clean
GO_TEST := $(GO) test

# Targets
.PHONY: all build run clean test

all: build

brun: build
	$(BUILD_DIR)/$(APP_NAME)

build:
	@echo "Building $(APP_NAME)..."
	$(GO_BUILD) -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)

run:
	@echo "Running $(APP_NAME)..."
	$(SRC_DIR)/$(APP_NAME)

clean:
	@echo "Cleaning up..."
	$(GO_CLEAN)
	rm -rf $(BUILD_DIR) $(BIN_DIR) $(PKG_DIR)

test:
	@echo "Running tests..."
	$(GO_TEST) ./...