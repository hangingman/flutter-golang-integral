# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

BIN=flutter-golang-integral

.PHONY: all

all: build

mod:
	$(GOCMD) mod download

protobuf: mod
	protoc --proto_path ./echo/proto --go_out=plugins=grpc:./pb --go_opt=paths=source_relative ./echo/proto/echo.proto

build: protobuf
	$(GOBUILD) -v -o $(BIN) main.go

clean:
	$(GOCLEAN)

run: build
	./$(BIN)

fmt:
	for go_file in `find . -name \*.go`; do \
		go fmt $${go_file}; \
	done
