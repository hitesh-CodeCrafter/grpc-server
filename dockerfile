# Use the official Golang image as a build stage
FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application from the /cmd/server directory
RUN go build -o main ./cmd/server

# Verify that the binary was created
RUN ls -la /app/main

# Use a minimal base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Install necessary libraries for running Go binaries
RUN apk add --no-cache libc6-compat

# Copy the binary from the build stage
COPY --from=builder /app/main .

# Verify that the binary was copied correctly
RUN ls -la /app/main

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
