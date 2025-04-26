.PHONY: migrate-create migrate-up migrate-down run

# Create new migration
# Usage: make migrate-create name=create_users_table
migrate-create:
	@if [ "$(name)" = "" ]; then \
		echo "Error: name is required. Usage: make migrate-create name=migration_name"; \
		exit 1; \
	fi
	@go run cmd/migrate/main.go $(name)

# Run migrations up
migrate-up:
	go run cmd/migrate/main.go

# Run migrations down
migrate-down:
	go run cmd/migrate/main.go down

# Run the application
run:
	go run cmd/api/main.go