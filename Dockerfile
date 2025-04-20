# Base image
FROM golang:1.24-bullseye AS builder

# Working directory
WORKDIR /app

# Install libraries
RUN apt-get update && \
  apt-get install -y --no-install-recommends \
  git \
  bash \
  wkhtmltopdf && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/*

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy code and environment
COPY . .
COPY .env.backup .env

# Build the Go application
RUN go build -o main .

# Expose port 3000
EXPOSE 3000

# Command to run the application
CMD ["./main"]
