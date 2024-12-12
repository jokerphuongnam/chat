# Use a minimal base image
FROM golang:1.23.2 AS builder

# Set the working directory
WORKDIR /app

# Copy Go modules manifest
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the log file
COPY config/config.yaml .

# Optional: Define default CMD if required
CMD ["cat", "config.yaml"]