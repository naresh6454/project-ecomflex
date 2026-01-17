.PHONY: build run test clean docker-up docker-down migrate

# Variables
APP_NAME = ecomflex-api
GO = go
DOCKER_COMPOSE = docker-compose
MIGRATE = migrate

# Build the application
build:
	$(GO) build -o $(APP_NAME) ./cmd/api

# Run the application
run:
	$(GO) run ./cmd/api

# Run tests
test:
	$(GO) test -v ./...

# Clean build artifacts
clean:
	rm -f $(APP_NAME)

# Start Docker containers
docker-up:
	$(DOCKER_COMPOSE) up -d

# Stop Docker containers
docker-down:
	$(DOCKER_COMPOSE) down

# Run migrations
migrate:
	$(GO) run ./migrations

# Create a new migration
migrate-create:
	@read -p "Migration name: " name; \
	$(MIGRATE) create -ext sql -dir migrations -seq $$name

# Help
help:
	@echo "Available commands:"
	@echo "  make build       - Build the application"
	@echo "  make run         - Run the application"
	@echo "  make test        - Run tests"
	@echo "  make clean       - Clean build artifacts"
	@echo "  make docker-up   - Start Docker containers"
	@echo "  make docker-down - Stop Docker containers"
	@echo "  make migrate     - Run migrations"
	@echo "  make migrate-create - Create a new migration"