export PATH:=$(PATH):$(shell go env GOPATH)/bin
SCRIPT_DIR := ./scripts

include .env

start-dev: go run main.go

test:
	 $(SCRIPT_DIR)/test_excluding.sh

runbuild:
	CGO_ENABLED=0 GOOS=linux go build -o bin/travel-request -tags='no_mysql no_sqlite3 no_ydb' main.go

start-prod: runbuild
	./bin/travel-request start