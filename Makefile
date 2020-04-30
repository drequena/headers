build:
	go build
	mv headers bin/

docker:
	docker build --no-cache -t "headers:latest" .
