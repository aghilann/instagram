# Stage 1: Build the Go app
FROM golang:1.23-alpine AS builder

# Install necessary C libraries for CGO
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Enable CGO for Go SQLite3 support
ENV CGO_ENABLED=1

# Set the current working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies first
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app with CGO enabled
RUN go build -o instagram ./cmd/main.go

# Stage 2: Create a lightweight runtime container for the Go app
FROM alpine:latest

# Install necessary libraries for SQLite at runtime
RUN apk add --no-cache sqlite-libs

# Set the current working directory inside the container
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/instagram .

# Copy any required files, like SQLite DB if necessary
COPY instagram.db .

# Expose the port the app listens on (if your app runs on port 8080)
EXPOSE 8080

# Run the Go app
CMD ["./instagram"]