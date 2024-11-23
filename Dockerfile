# Stage 1: Development (with air for hot reloading)
FROM golang:1.23-alpine AS development

# Install air for live reloading
RUN go install github.com/air-verse/air@latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to cache dependencies
COPY go.mod go.sum ./

# Download all Go dependencies
RUN go mod download

# Copy the rest of the project files
COPY . .

# Command to run Air for live reloading
CMD ["air"]

# Stage 2: Production (without air, just the Go binary)
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the files
COPY . .

# Build the Go app
RUN go build -o /app/main ./cmd/httpserver/main.go

# Stage 3: Final image (for production)
FROM alpine:latest

# Set working directory in the final image
WORKDIR /app

# Install jq for JSON parsing
RUN apk add --no-cache jq

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy the entrypoint script
COPY entrypoint.sh /app/entrypoint.sh

# Ensure the entrypoint script is executable
RUN chmod +x /app/entrypoint.sh

# Expose the port your app will run on
EXPOSE 8080

# Set the entrypoint to your script
ENTRYPOINT ["/app/entrypoint.sh"]

# Run the main application
CMD ["./main"]
