# Stage 1: Build the Go application
FROM golang:1.26.1-alpine AS builder

WORKDIR /app 

# install git rquired for go mod tidy
RUN apk add --no-cache git  

# Set Go proxy fallback to direct to avoid issues with private modules
RUN go env -w GOPROXY=https://proxy.golang.org,direct

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
# Download dependencies
RUN go mod tidy
# Copy the source code
COPY . .
# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# ---------- Stage 2: Minimal Runtime ----------
FROM alpine:latest

WORKDIR /root/

# Install certs (important if your app calls HTTPS APIs)
RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /app/main .

# Copy config (only if needed at runtime)
COPY config.yml .

# Run app
CMD ["./main"]