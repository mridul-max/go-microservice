# Use the official Golang image as the base image
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Initialize the Go module inside the Docker container
RUN go mod init drinks || true

# Download dependencies. go mod tidy will also create go.mod and go.sum if not already present
COPY . . 
RUN go mod tidy

# Install Swagger CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Run the swag init command to generate Swagger docs
RUN swag init

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest  

# Set the working directory in the final image
WORKDIR /root/

# Copy the Pre-built binary file from the builder stage
COPY --from=builder /app/main .

# Make the binary executable
RUN chmod +x ./main

# Expose port 8082 to the outside world
EXPOSE 8082

# Command to run the executable
CMD ["./main"]
