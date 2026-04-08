# ---- Build Stage ----
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy dependency files first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -o app ./cmd/api


# ---- Run Stage ----
FROM alpine:latest

WORKDIR /root/

# Copy compiled binary
COPY --from=builder /app/app .

# Expose app port
EXPOSE 8080

# Run app
CMD ["./app"]