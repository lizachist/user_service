# Переменные
BINARY_NAME=user_service
MAIN_PATH=./cmd/main.go

# Команды Go
GO=go
GORUN=$(GO) run
GOBUILD=$(GO) build
GOTEST=$(GO) test
GOCLEAN=$(GO) clean

# Команды Docker
DOCKER=docker
DOCKER_COMPOSE=docker-compose

# Цели
.PHONY: all build run test clean docker docker-build docker-run

all: build

build:
    $(GOBUILD) -o $(BINARY_NAME) $(MAIN_PATH)

run:
    $(GORUN) $(MAIN_PATH)

test:
    $(GOTEST) -v ./...

clean:
    $(GOCLEAN)
    rm -f $(BINARY_NAME)

docker-build:
    $(DOCKER) build -t $(BINARY_NAME) .

docker-run:
    $(DOCKER) run -p 8080:8080 $(BINARY_NAME)

docker-compose-up:
    $(DOCKER_COMPOSE) up -d

docker-compose-down:
    $(DOCKER_COMPOSE) down

migrate-up:
    migrate -path ./migrations -database "$${DATABASE_URL}" up

migrate-down:
    migrate -path ./migrations -database "$${DATABASE_URL}" down

lint:
    golangci-lint run