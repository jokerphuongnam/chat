# Use a minimal base image
FROM golang:1.23.2 AS builder

# Set the working directory
WORKDIR /app

# Copy Go modules manifest
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy