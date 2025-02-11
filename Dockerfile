# Use the official Go image
FROM golang:1.21

# Set the working directory
WORKDIR /app

# Install an older compatible version of Air (since latest requires Go 1.23+)
RUN go install github.com/cosmtrek/air@v1.43.0 && \
    cp /go/bin/air /usr/local/bin/air

# Copy go.mod and go.sum first (for efficient caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Expose the application port
EXPOSE 8080

# Run Air for hot-reloading
CMD ["air"]