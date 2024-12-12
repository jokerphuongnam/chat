# Stage 1: Build the application
FROM golang:1.23.2 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules manifest
COPY go.mod go.sum ./

# Assuming your Go source code and internal packages are in a parent directory
COPY ../internal /app/internal

# Copy the full Go source code into the container
COPY . .

# Download dependencies
RUN go mod tidy
RUN go mod download

# Copy the full Go source code into the container
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o chat-backend

# Copy the source code
COPY . .

# Stage 2: Run the application
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/api-gateway .

# Make the binary executable
RUN chmod +x chat-backend

# Run the binary
CMD ["./chat-backend"]

# Expose the API gateway port
EXPOSE 8080