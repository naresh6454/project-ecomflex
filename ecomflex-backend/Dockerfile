FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git for dependencies
RUN apk add --no-cache git

# Copy workspace and module files first
COPY go.work* ./
COPY go.mod go.sum ./

# Force module update
RUN go mod tidy

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ecomflex-api ./cmd/api

# Final stage
FROM alpine:3.17

WORKDIR /app

# Add required packages
RUN apk add --no-cache ca-certificates tzdata dos2unix

# Copy binary and config
COPY --from=builder /app/ecomflex-api /app/
COPY --from=builder /app/config /app/config
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /app/.env /app/.env
COPY scripts/entrypoint.sh /app/entrypoint.sh

# Make sure the entrypoint script is executable and fix line endings
RUN dos2unix /app/entrypoint.sh && chmod +x /app/entrypoint.sh

EXPOSE 8000

# Use entrypoint script instead of direct command
ENTRYPOINT ["/app/entrypoint.sh"]