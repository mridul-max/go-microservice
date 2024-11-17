# Use the official Golang image as the base image
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Initialize the Go module inside the Docker container
RUN go mod init drinks || true

# Download dependencies
COPY . .
RUN go mod tidy

# Install Swagger CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Run the swag init command to generate Swagger docs
RUN swag init

# Build the Go app
RUN go build -o main .

# Set executable permissions in the build stage to ensure they're carried over
RUN chmod +x ./main

# Start a new stage from scratch
FROM alpine:latest  

# Set the working directory in the final image
WORKDIR /root/

# Copy the Pre-built binary file and Swagger docs from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

# Explicitly set execute permissions in the final stage
RUN chmod +x ./main

# Expose port 8082 to the outside world
EXPOSE 8082

# Command to run the executable
CMD ["./main"]
