FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o home-automation

# Start a new stage from scratch
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/home-automation /app/home-automation
# Copy the web folder
COPY /web /app/web
# Make sure the binary is executable
RUN chmod +x /app/home-automation

# Create entrypoint script
RUN echo -e '#!/bin/sh\nexec /app/home-automation' > /entrypoint.sh && chmod +x /entrypoint.sh

# Command to run the entrypoint script
ENTRYPOINT ["/entrypoint.sh"]
