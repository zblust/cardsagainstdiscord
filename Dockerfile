# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cah-bot ./cmd/cardsagainstdiscord

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -g 1000 cahbot && \
    adduser -D -u 1000 -G cahbot cahbot

WORKDIR /app

# Copy binary from builder
COPY --from=builder /build/cah-bot .

# Change ownership
RUN chown -R cahbot:cahbot /app

# Switch to non-root user
USER cahbot

# Expose pprof port (optional, for debugging)
EXPOSE 7447

# Run the bot
CMD ["./cah-bot"]
