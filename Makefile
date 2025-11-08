# ============================================================
# Makefile for games-api project
# Author: Pedro Ivo
# ============================================================

APP_NAME := gamesapi
BINARY := $(APP_NAME)
IMAGE := $(APP_NAME)
CONTAINER := $(APP_NAME)
COMPOSE := docker-compose

GO := go
GOFILES := $(shell find . -name '*.go' -not -path './vendor/*')

SEED_PATH := data/games_seed.json
PORT := 3000


# ============================================================
##@Build / Local
# ============================================================

## Build the Go binary locally
build:
	@echo "🧱 Building $(BINARY)..."
	$(GO) build -o bin/$(BINARY) ./cmd
	@echo "✅ Binary built at bin/$(BINARY)"

## Run the application locally (without Docker)
run:
	@echo "🚀 Running locally..."
	$(GO) run cmd/main.go

## Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	rm -rf bin
	rm -rf out

## Tidy and verify Go modules
tidy:
	@echo "📦 Tidy and verify modules..."
	$(GO) mod tidy
	$(GO) mod verify

## Run unit tests
test:
	@echo "🧪 Running tests..."
	$(GO) test ./... -v

## Format source code
lint:
	@echo "🔍 Formatting code..."
	$(GO) fmt ./...
	@echo "✅ All code formatted."

# ============================================================
##@Docker
# ============================================================

## Build Docker image - via Dockerfile
docker-build:
	@echo "🐳 Building Docker image $(IMAGE)..."
	docker build -t $(IMAGE) .

## Run Docker container - via Dockerfile
docker-run:
	@echo "🐳 Running Docker image $(IMAGE)..."
	docker run -d --name $(IMAGE) -p $(PORT):$(PORT) $(IMAGE)

## Start containers - via docker-compose
up:
	@echo "⬆️  Starting containers..."
	$(COMPOSE) up --build -d

## Stop and remove containers - via docker-compose
down:
	@echo "⬇️  Stopping containers..."
	$(COMPOSE) down

## Follow container logs
logs:
	@echo "📜 Following logs..."
	$(COMPOSE) logs -f $(CONTAINER)

## Show running containers
ps:
	$(COMPOSE) ps

## Restart the container
restart:
	$(COMPOSE) restart $(CONTAINER)

## Remove Docker image and containers
docker-clean:
	@echo "🔥 Removing image and containers..."
	-docker rm -f $(CONTAINER) || true
	-docker rmi -f $(IMAGE) || true

## Open interactive shell inside the container
shell:
	@echo "🔧 Opening shell inside $(CONTAINER)..."
	docker exec -it $(CONTAINER) /bin/sh

# ============================================================
##@Utilities
# ============================================================

## Check the /health endpoint
health:
	@echo "💡 Checking /health endpoint..."
	curl -s http://localhost:$(PORT)/health | head -n 10

# ============================================================
##@Meta
# ============================================================

## Show this help message
help:
	@echo ""
	@echo -e "\033[1;33m📘 Available commands for $(APP_NAME):\033[0m"
	@awk ' \
		BEGIN { FS = ""; section="General"; } \
		/^##@/ { section=substr($$0,4); gsub(/^[ \t]+|[ \t]+$$/,"",section); printf "\n\033[1;35m%s\033[0m\n", section; next } \
		/^## / { desc=substr($$0,4); getline; if (match($$0,/^[a-zA-Z0-9._-]+/)) { target=substr($$0,1,RLENGTH); printf "  \033[36m%-20s\033[0m %s\n", target, desc } } \
	' $(MAKEFILE_LIST)
	@echo ""

# ============================================================
# Default target
# ============================================================

.DEFAULT_GOAL := help