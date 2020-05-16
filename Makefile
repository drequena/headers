CWD := $(shell pwd)

help:  ## Show available commands
	@echo "Available commands:"
	@echo
	@sed -n -E -e 's|^([A-Za-z0-9/_-]+):.+## (.+)|\1@\2|p' $(MAKEFILE_LIST) | column -s '@' -t

build: test ## Run tests, build binary and move to ./bin
	go build
	mkdir -p ./bin 2> /dev/null
	mv headers bin/

test: ## Run unit tests
	go test -v

docker: ## Build docker image called headers
	docker build --no-cache -t "headers:latest" .

check: ## Check kong syntax file
	docker run --rm -e KONG_DATABASE="off" -v "$(CWD)/kong/kong.yaml:/kong.yaml" kong:2.0.3-alpine kong config parse /kong.yaml

all: ## Build binary and docker image
	build 
	docker

.PHONY: build test docker check all help