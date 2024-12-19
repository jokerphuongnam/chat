# Stage 1: Build the application
FROM golang:1.23.2 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules manifest
COPY chat-backend/go.mod chat-backend/go.sum ./

# Download dependencies
RUN go mod tidy
RUN go mod download

# Copy the Go source code into the container
COPY chat-backend/ ./chat-backend/
COPY internal/ ./internal/

# Set the working directory to the chat-backend directory
WORKDIR /app/chat-backend

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o chat-backend

# Final stage: create a lightweight runtime image
FROM gcr.io/distroless/base-debian11

# Copy the source code
COPY . .

# Stage 2: Run the application
FROM alpine:latest

# Set the working directory
WORKDIR /app/chat-backend

# Copy the binary from the builder stage
COPY --from=builder /app/chat-backend .

# Make the binary executable
RUN chmod +x chat-backend

# Run the binary
CMD ["./chat-backend"]

# Expose the API gateway port
EXPOSE 3300