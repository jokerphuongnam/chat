# Use a minimal base image
FROM golang:1.23.2 AS builder

# Set the working directory
WORKDIR /app

# Copy Go modules manifest
COPY go.mod go.sum ./

# Copy the source code
COPY . .

# Download dependencies
RUN go mod tidy
RUN go mod download

# Ensure the logs directory exists and create the chat.log file if it doesn't exist
RUN mkdir -p logs && [ ! -f logs/chat.log ] && echo "This is a sample log entry" > logs/chat.log || echo "chat.log already exists"

# Verify the file creation (optional, for debugging)
RUN [ -f "logs/chat.log" ] && echo "chat.log verified"

# Copy the log file
COPY logs/chat.log .

# Copy the source code
COPY . .

# Optional: Define default CMD if required
CMD ["cat", "chat.log"]