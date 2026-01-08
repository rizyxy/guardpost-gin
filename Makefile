# Variables
BINARY_NAME=guardpost
DOCKER_IMAGE=guardpost-gin

.PHONY: build run test docker-up docker-down clean

build:
	go build -o $(BINARY_NAME) main.go

run: build
	./$(BINARY_NAME)

test:
	go test ./...

docker-up:
	docker-compose up --build -d

docker-down:
	docker-compose down

clean:
	rm -f $(BINARY_NAME)
	rm -rf data/