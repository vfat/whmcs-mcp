# Makefile for WHMCS MCP Server

VERSION ?= v0.1.0
LDFLAGS := -s -w -X github.com/vfat/whmcs-mcp/internal/version.Version=$(VERSION)

.PHONY: all build test test-coverage build-all clean docker-build

all: test build

build:
	@mkdir -p bin
	CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o bin/whmcs-mcp ./cmd/whmcs-mcp

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

# Cross compilation for multi-platform distribution
build-all:
	@mkdir -p dist
	@echo "Building Linux amd64..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o dist/whmcs-mcp_linux_amd64 ./cmd/whmcs-mcp
	@echo "Building Linux arm64..."
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o dist/whmcs-mcp_linux_arm64 ./cmd/whmcs-mcp
	@echo "Building macOS amd64 (Intel)..."
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o dist/whmcs-mcp_darwin_amd64 ./cmd/whmcs-mcp
	@echo "Building macOS arm64 (Apple Silicon)..."
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o dist/whmcs-mcp_darwin_arm64 ./cmd/whmcs-mcp
	@echo "Building Windows amd64..."
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o dist/whmcs-mcp_windows_amd64.exe ./cmd/whmcs-mcp
	@echo "All binaries successfully compiled in dist/"

docker-build:
	docker build --build-arg VERSION=$(VERSION) -t whmcs-mcp:$(VERSION) -t whmcs-mcp:latest .

clean:
	rm -rf bin dist coverage.out
