#!/bin/bash

echo "Build ./main"
go build -o ./bin/main ./

echo "Build ./cmd/migrate"
go build -o ./bin/migrate ./cmd/migrate

echo "Build ./cmd/generate-model"
go build -o ./bin/generate-model ./cmd/generate-model

echo "Build ./cmd/seeder"
go build -o ./bin/seeder ./cmd/seeder