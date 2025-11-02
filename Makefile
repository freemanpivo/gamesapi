# ============================================================
# Makefile for games-api project
# Author: Pedro Ivo
# ============================================================

APP_NAME := games-api
BINARY := $(APP_NAME)
IMAGE := $(APP_NAME):local
CONTAINER := $(APP_NAME)
COMPOSE := docker-compose

GO := go
GOFILES := $(shell find . -name '*.go' -not -path './vendor/*')

SEED_PATH := data/games_seed.json
PORT := 3000

# ============================================================
# Commands
# ============================================================

build:
	@echo "🧱 Building $(BINARY)..."
	$(GO) build -o bin/$(BINARY) ./main.go
	@echo "✅ Binary built at bin/$(BINARY)"

run:
	@echo "🚀 Running locally..."
	$(GO) run main.go

tidy:
	@echo "📦 Tidy and verify modules..."
	$(GO) mod tidy
	$(GO) mod verify

clean:
	@echo "🧹 Cleaning..."
	rm -rf bin
	rm -rf out

# ============================================================
# Docker commands
# ============================================================

docker-build:
	@echo "🐳 Building Docker image $(IMAGE)..."
	docker build -t $(IMAGE) .

up:
	@echo "⬆️  Starting containers..."
	$(COMPOSE) up --build -d

down:
	@echo "⬇️  Stopping containers..."
	$(COMPOSE) down

logs:
	@echo "📜 Following logs..."
	$(COMPOSE) logs -f $(CONTAINER)

ps:
	$(COMPOSE) ps

restart:
	$(COMPOSE) restart $(CONTAINER)

docker-clean:
	@echo "🔥 Removing image and containers..."
	-docker rm -f $(CONTAINER) || true
	-docker rmi -f $(IMAGE) || true

shell:
	@echo "🔧 Opening shell inside $(CONTAINER)..."
	docker exec -it $(CONTAINER) /bin/sh

# ============================================================
# Utility / Shortcuts
# ============================================================

test:
	@echo "🧪 Running tests..."
	$(GO) test ./... -v

lint:
	@echo "🔍 Formatting code..."
	$(GO) fmt ./...
	@echo "✅ All code formatted."

health:
	@echo "💡 Checking /games health endpoint..."
	curl -s http://localhost:$(PORT)/games | head -n 10

# ============================================================
# Help
# ============================================================

help:
	@echo ""
	@echo "📘 Available commands:"
	@echo ""
	@grep -E '^##' $(MAKEFILE_LIST) | sed 's/## //'
	@echo ""

# ============================================================
# Default
# ============================================================

.DEFAULT_GOAL := help
