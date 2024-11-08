# Stage 1: Build the Go application
FROM golang:1.23.2-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main cmd/myservice/main.go

# Stage 2: Create a lightweight image for running the application
FROM alpine:3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the service will run on
EXPOSE 8080

# Command to run the application
CMD ["./main"]

