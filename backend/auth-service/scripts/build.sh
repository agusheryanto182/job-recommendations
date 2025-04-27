#!/bin/bash

echo "Generating proto files..."
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/auth/auth.proto 

echo "Building main service..."
go build -o ./bin/main ./

echo "Building migration tool..."
go build -o ./bin/migrate ./cmd/migrate

echo "Building model generator..."
go build -o ./bin/generate-model ./cmd/generate-model

echo "Building seeder..."
go build -o ./bin/seeder ./cmd/seeder