CWD := $(shell pwd)

build:
	go build
	mkdir -p ./bin 2> /dev/null
	mv headers bin/

docker:
	docker build --no-cache -t "headers:latest" .

check:
	docker run --rm -e KONG_DATABASE="off" -v "$(CWD)/kong/kong.yaml:/kong.yaml" kong:2.0.3-alpine kong config parse /kong.yaml

all: build docker

.PHONY: build docker check all