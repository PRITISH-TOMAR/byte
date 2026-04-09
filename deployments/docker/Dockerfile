# ---- Build Stage ----
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy dependency files first
COPY go.mod ./

# Download what's available (go.sum will be generated after)
RUN go mod download || true

# Copy all source files
COPY . .

# Now tidy and download with full source context
RUN go mod tidy
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api


# ---- Run Stage ----
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]