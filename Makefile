.PHONY: build run clean build-all

APP_NAME := ExpressTS
DIST_DIR := dist
MAIN := ./main.go

build:
	go build -o $(APP_NAME).out $(MAIN)

run:
	go run $(MAIN)

build-all:
	mkdir -p $(DIST_DIR)
	GOOS=linux   GOARCH=amd64 go build -o $(DIST_DIR)/$(APP_NAME)-linux-amd64 $(MAIN)
	GOOS=linux   GOARCH=arm64 go build -o $(DIST_DIR)/$(APP_NAME)-linux-arm64 $(MAIN)
	GOOS=darwin  GOARCH=amd64 go build -o $(DIST_DIR)/$(APP_NAME)-macos-amd64 $(MAIN)
	GOOS=darwin  GOARCH=arm64 go build -o $(DIST_DIR)/$(APP_NAME)-macos-arm64 $(MAIN)
	GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/$(APP_NAME)-windows-amd64.exe $(MAIN)

clean:
	rm -rf $(DIST_DIR) $(APP_NAME).out
