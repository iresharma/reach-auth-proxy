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
RUN go build -o /gin-app .

# Stage 2: Create a smaller image with just the built binary
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /gin-app .

#Injects variables from railway
ARG POSTGRES
ARG KANBAN_SERVER
ARG PAGE_SERVER
ARG APP_URL
ARG RESEND_API
ARG REDIS
ARG BASE_URL

# Set environment variables
ENV GIN_MODE=release
ENV POSTGRES=$POSTGRES
ENV KANBAN_SERVER=$KANBAN_SERVER
ENV PAGE_SERVER=$PAGE_SERVER
ENV APP_URL=$APP_URL
ENV RESEND_API=$RESEND_API
ENV REDIS=$REDIS
ENV BASE_URL=$BASE_URL

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./gin-app"]
