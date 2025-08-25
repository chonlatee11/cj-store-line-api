.PHONY: run build test lint clean migrate-up migrate-down seed docker-build docker-run

# Variables
APP_NAME=cj-store-line-api
BUILD_DIR=./bin
DOCKER_IMAGE=$(APP_NAME):latest

# Development commands
run:
	@echo "Starting the application..."
	go run main.go

build:
	@echo "Building the application..."
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) main.go

test:
	@echo "Running tests..."
	go test -v -race -cover ./...

lint:
	@echo "Running linter..."
	golangci-lint run

clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)

# Database commands
migrate-up:
	@echo "Running database migrations..."
	go run main.go migrate up

migrate-down:
	@echo "Rolling back database migrations..."
	go run main.go migrate down

seed:
	@echo "Seeding database..."
	go run scripts/seed.go

# Docker commands
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 --env-file .env $(DOCKER_IMAGE)

# Go module commands
mod-download:
	@echo "Downloading Go modules..."
	go mod download

mod-tidy:
	@echo "Tidying Go modules..."
	go mod tidy

# Development setup
setup: mod-download
	@echo "Setting up development environment..."
	cp .env.example .env
	@echo "Please update .env with your configuration"

# Help
help:
	@echo "Available commands:"
	@echo "  run          - Run the application"
	@echo "  build        - Build the application"
	@echo "  test         - Run tests"
	@echo "  lint         - Run linter"
	@echo "  clean        - Clean build artifacts"
	@echo "  migrate-up   - Run database migrations"
	@echo "  migrate-down - Rollback database migrations"
	@echo "  seed         - Seed database"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
	@echo "  mod-download - Download Go modules"
	@echo "  mod-tidy     - Tidy Go modules"
	@echo "  setup        - Setup development environment"
