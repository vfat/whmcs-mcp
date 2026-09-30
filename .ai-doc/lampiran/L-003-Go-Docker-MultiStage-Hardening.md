# Lampiran L-003: Multi-Stage Build & Container Hardening Go

> **Topik**: Kompilasi Statis Biner Go, Multi-Stage Dockerfile Minimalis, dan Pengamanan Kontainer  
> **Workspace**: `/home/ubuntu/workspace/plan/whmcs-mcp`  
> **Target Sistem**: `whmcs-mcp` (Go 1.22+)  
> **Tanggal**: 2026-09-30  

---

## 1. Keunggulan Kontainerisasi Go vs Node.js

| Karakteristik | Implementasi TypeScript / Node.js | Implementasi Golang |
|:---|:---:|:---:|
| **Ukuran Image** | ~150–200 MB (Node 20 Alpine) | **~15–20 MB** (Alpine) / **~10 MB** (Scratch) |
| **Ketergantungan Runtime** | Node.js engine, npm, node_modules | **Nihil** (Single Static Binary) |
| **Attack Surface** | Ratusan paket dependensi JS pihak ketiga | **Minimal** (Hanya dependensi biner statis) |
| **Cold Start** | ~300–800 ms | **< 10 ms** |
| **Memory RSS** | ~100–180 MB | **~15–25 MB** |

---

## 2. Definisi Multi-Stage Dockerfile

```dockerfile
# Stage 1: Build Environment
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install CA certificates untuk HTTPS calls
RUN apk add --no-cache ca-certificates git

# Cache Go modules
COPY go.mod go.sum ./
RUN go mod download

# Salin source code
COPY . .

# Kompilasi biner statis tanpa CGO
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w" \
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
```

---

## 3. Menjalankan Server via Docker pada MCP Client

Ketika Claude Desktop atau Cursor menjalankan server MCP di dalam Docker:

```json
{
  "mcpServers": {
    "whmcs": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "-e", "WHMCS_URL=https://billing.example.com",
        "-e", "WHMCS_API_IDENTIFIER=my-identifier",
        "-e", "WHMCS_API_SECRET=my-secret",
        "ghcr.io/vfat/whmcs-mcp:latest"
      ]
    }
  }
}
```

Flag `-i` (*interactive*) memastikan stream `stdin` terhubung ke transport stdio biner Go di dalam kontainer.
