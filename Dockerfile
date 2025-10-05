# Use official Go image
FROM golang:1.24-alpine

# Set working directory inside container
WORKDIR /urlshortner

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go app
RUN go build -o url-shortener main.go

# Expose app port
EXPOSE 8080

# Run the app
CMD ["./url-shortener"]
