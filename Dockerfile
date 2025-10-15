# --- Build Stage ---
# Use the official Golang image to create a build artifact.
# This is known as a "multi-stage build," which helps keep the final image size small.
FROM golang:1.25-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
# -o /app/server creates the binary in the /app directory
# CGO_ENABLED=0 is important for creating a static binary that doesn't depend on C libraries,
# which is ideal for a minimal container image.
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/server

# --- Final Stage ---
# Use a minimal image for the final container.
# "scratch" is a special Docker image that is empty, providing the smallest possible base.
# Using alpine is also a good choice if you need a shell for debugging.
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the built binary from the "builder" stage
COPY --from=builder /app/server .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./server"]
