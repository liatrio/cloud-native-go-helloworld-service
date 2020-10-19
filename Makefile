MAKEFLAGS += --silent

.PHONY: run test

default:
	echo No default target

run:
	go run main.go

test:
	go test ./...
