# Use a newer Go version as the base image
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Initialize the Go module inside the Docker container
RUN go mod init drink || true

# Download dependencies
COPY . .
RUN go mod tidy

# Install Swagger CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Generate Swagger docs
RUN swag init

# Build the Go app as a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Start a new stage from scratch
FROM alpine:latest

# Create a non-root user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the working directory in the final image
WORKDIR /app

# Copy the Pre-built binary file and Swagger docs
COPY --from=builder /app/main /usr/local/bin/main
COPY --from=builder /app/docs /app/docs

# Ensure the binary has executable permissions
RUN chmod +x /usr/local/bin/main

# Change ownership to the non-root user
RUN chown -R appuser:appgroup /app

# Switch to the non-root user
USER appuser

# Expose the application port
EXPOSE 8082

# Command to run the executable
CMD ["main"]
