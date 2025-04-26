FROM golang:1.24-alpine AS builder

RUN apk update && apk add git
RUN apk add --no-cache --upgrade bash

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

# Stage 2: Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .

# Expose port
EXPOSE 8080

ENTRYPOINT go mod download && ./scripts/build.sh && ./bin/migrate up && ./bin/main