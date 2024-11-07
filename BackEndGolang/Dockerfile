# Stage 1: Build
FROM golang:1.23.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Copy the config file
COPY config/config.yaml /app/config/config.yaml
COPY internal/logs/chat.log /app/internal/logs/chat.log

# Download the dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

# Stage 2: Run
FROM alpine:latest

# Install curl
RUN apk --no-cache add curl

# Set the working directory in the new stage
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Make the binary executable
RUN chmod +x main

# Command to run the binary
CMD ["./main"]

EXPOSE 3031
