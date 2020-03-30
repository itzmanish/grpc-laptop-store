-include .env

VERSION := $(shell git describe --tags --always)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")

# Go related variables.
GOPATH := $(shell echo "$(GOPATH)")
# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

# Redirect error output to a file, so we can show it in development mode.
STDERR := /tmp/.$(PROJECTNAME)-stderr.txt

# PID file will keep the process id of the server
PID := /tmp/.$(PROJECTNAME).pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

##default: Default behaviour
default:
	@echo
	@echo see make help
## gen_pb: Generate protobuffer file from proto file with grpc plugin included.
gen_pb: 
	@echo
	@echo "Generating protobuffer file from proto."
	@protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb 
	@echo

test:
	@echo
	@echo "Running tests on tests folder."
	@go test -cover -race ./...
	@echo

clean:
	@echo 
	@echo "Cleaning protobuffer files"
	@rm pb/*.go
	@echo 

server:
	@echo
	@go run cmd/server/main.go -port 50051
	@echo

client:
	@echo
	@go run cmd/client/main.go -address 0.0.0.0:50051
	@echo

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
