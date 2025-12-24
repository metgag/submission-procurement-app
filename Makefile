# Variables
DOCKER_CONTAINER=procurement-db
DB_HOST=localhost
DB_PORT=5436
DB_USER=postgres
DB_PASS=postgres
DB_NAME=procurement

install: ## Install Go dependencies
	@echo "Installing dependencies..."
	go mod tidy
	go mod download
	@echo "Dependencies installed successfully!"

run: ## Run the application
	@echo "Starting application..."
	go run ./cmd/server

docker-up: ## Start PostgreSQL container
	@echo "Starting PostgreSQL container..."
	docker run --name $(DOCKER_CONTAINER) \
		-e POSTGRES_USER=$(DB_USER) \
		-e POSTGRES_PASSWORD=$(DB_PASS) \
		-e POSTGRES_DB=$(DB_NAME) \
		-p $(DB_PORT):5432 \
		-d postgres:16.11
	@echo "PostgreSQL container started!"
	@echo "Waiting for PostgreSQL to be ready..."
	@sleep 2

db-seed: ## Seed database with sample data (via Docker)
	@echo "Seeding database..."
	@docker cp scripts/seed.sql $(DOCKER_CONTAINER):/tmp/seed.sql
	@docker exec -i $(DOCKER_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME) -f /tmp/seed.sql
	@echo "Database seeded successfully!"

clean: ## Clean build artifacts and cache
	@echo "Cleaning..."
	go clean -cache
	rm -rf bin/
	@echo "Clean complete!"
