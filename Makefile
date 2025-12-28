.PHONY: build run clean

build:
	 go build -o ExpressTS.out ./main.go

run:
	go run ./main.go