# Dockerfile

# ===== STAGE 1: BUILD =====
FROM golang:1.20-alpine AS builder

# Create and switch to the app directory
WORKDIR /app

# Copy go.mod and go.sum first to leverage caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the binary (name it 'optimus' for convenience)
RUN go build -o /bin/optimus main.go

# ===== STAGE 2: RUN =====
FROM alpine:3.18

WORKDIR /app

# Copy the compiled binary from STAGE 1
COPY --from=builder /bin/optimus /app/optimus

# Expose the port your app will listen on (8080, in your case)
EXPOSE 8080

# Run the binary
CMD ["/app/optimus"]