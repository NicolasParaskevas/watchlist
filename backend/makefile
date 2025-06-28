APP_NAME := watchlist
CMD_DIR := ./cmd
BUILD_DIR := ./bin

.PHONY: all build run test tidy clean

all: build

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_DIR)/main.go

run:
	go run $(CMD_DIR)/main.go

test:
	go test ./...

tidy:
	go mod tidy

clean:
	rm -rf $(BUILD_DIR)
