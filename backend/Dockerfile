FROM golang:1.24-alpine

WORKDIR /app

# Install required system dependencies
RUN apk add --no-cache gcc musl-dev

# Copy go mod files first (for better caching)
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the application
RUN go build -o main ./cmd/server

EXPOSE 8080

CMD ["./main"]