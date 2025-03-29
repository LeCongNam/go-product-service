# Use the official Golang image as the base image
FROM golang:tip-alpine3.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main ./cmd/server/main.go

# Final stage: Minimal runtime environment
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the built application from the builder stage
COPY --from=builder /app/main .
# Nếu bạn muốn copy file .env vào image (lưu ý: khi dùng docker-compose env_file cũng được bind vào container)
COPY .env .  

# Expose the port the Gin app runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
