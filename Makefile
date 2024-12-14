BUILD_DIR = bin
APP_NAME = ecom
MAIN_FILE = main.go
MIGRATE_DIR = cmd/migrate/migrations
MIGRATE_MAIN = cmd/migrate/main.go
DB_URL = postgres://username:password@localhost:5432/your_database_name?sslmode=disable

.PHONY: build run test migration migrate-up migrate-down clean

build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

run: build
	@echo "Running $(APP_NAME)..."
	@./$(BUILD_DIR)/$(APP_NAME)

test:
	@echo "Running tests..."
	@go test -v ./...

migration:
	@echo "Creating a new migration..."
	@migrate create -ext sql -dir $(MIGRATE_DIR) -seq $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@echo "Applying migrations..."
	@go run $(MIGRATE_MAIN) up

migrate-down:
	@echo "Reverting last migration..."
	@go run $(MIGRATE_MAIN) down

clean:
	@echo "Cleaning build files..."
	@rm -rf $(BUILD_DIR)

%:
	@:
