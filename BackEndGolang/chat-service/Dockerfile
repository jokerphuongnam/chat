# Stage 1: Build the application
FROM golang:1.23.2 AS builder

# Set the working directory
WORKDIR /app

# Copy Go modules manifest
COPY go.mod go.sum ./

# Assuming your Go source code and internal packages are in a parent directory
COPY ../internal /app/internal

# Copy the full Go source code into the container
COPY . .

# Download dependencies
RUN go mod tidy
RUN go mod download

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o chat-service

# Copy the source code
COPY . .

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