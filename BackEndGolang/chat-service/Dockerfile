# Stage 1: Build the application
FROM golang:1.23.2 AS builder

# Set the working directory
WORKDIR /app

# Copy Go modules manifest
COPY go.mod go.sum ./

# Copy the config file
COPY config/config.yaml /app/config/config.yaml
COPY internal/logs/chat.log /app/internal/logs/chat.log

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o chat-service

# Stage 2: Run the application
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/chat-service .

# Make the binary executable
RUN chmod +x chat-service

# Run the binary
CMD ["./chat-service"]

# Expose the chat service port
EXPOSE 3031 3032 3033