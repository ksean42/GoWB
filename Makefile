run:
	go run -v ./cmd

build:
	go build -o service ./cmd
	go build -o pubMock ./pubServer
	
.DEFAULT_GOAL := run