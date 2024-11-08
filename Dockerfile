# Stage 1: Build the Go application
FROM golang:1.19-alpine AS builder

# Install git and other dependencies
RUN apk update && apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Set the working directory for the build context
WORKDIR /app/cmd

# Build the Go application
RUN go build -o /reach-proxy .

# Stage 2: Create a smaller image with just the built binary
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./reach-proxy"]
