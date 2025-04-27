.PHONY: migrate-create migrate-up migrate-down run proto

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

proto:
	# Install protoc generators if not exist
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

	# Generate proto files
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/auth/*.proto