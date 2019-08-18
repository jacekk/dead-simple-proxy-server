.DEFAULT_GOAL := help

help:
	@grep '^[^#\.[:space:]].*:' Makefile | tr -d ':'

clear-cache:
	rm -rf ./cache/*.txt

build:
	go build cmd/main.go

server:
	go run cmd/main.go server

worker:
	go run cmd/main.go worker

run-all:
	go run cmd/main.go all
