# Use Go base image
FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go app
RUN go build -o main ./cmd/main.go

# Run the app
CMD ["./main"]
