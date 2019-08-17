.DEFAULT_GOAL := help

help:
	@grep '^[^#\.[:space:]].*:' Makefile | tr -d ':'

build:
	go build cmd/main.go

serve:
	go run cmd/main.go server
