# Multi-stage Dockerfile for WHMCS MCP Server
# Produces a minimalist, hardened static binary container (~15MB)

# Stage 1: Build Environment
FROM golang:1.22-alpine AS builder

ARG VERSION=v0.1.0

WORKDIR /app

# Install CA certificates for HTTPS WHMCS API calls
RUN apk add --no-cache ca-certificates git

# Cache Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source tree
COPY . .

# Compile static binary without CGO and strip debug symbols
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w -X github.com/vfat/whmcs-mcp/internal/version.Version=${VERSION}" \
    -o /app/whmcs-mcp ./cmd/whmcs-mcp

# Stage 2: Minimalist Production Runner
FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata && \
    addgroup -g 1001 appgroup && \
    adduser -u 1001 -G appgroup -s /bin/sh -D appuser

WORKDIR /app

COPY --from=builder /app/whmcs-mcp /app/whmcs-mcp

USER 1001:1001

ENTRYPOINT ["/app/whmcs-mcp"]
