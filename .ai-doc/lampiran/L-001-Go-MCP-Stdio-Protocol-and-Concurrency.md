# Lampiran L-001: Protokol Stdio MCP & Konkurensi Go

> **Topik**: Integrasi Transport Stdio JSON-RPC 2.0, Pemisahan I/O Stream, dan Manajemen Konkurensi di Go  
> **Workspace**: `/home/ubuntu/workspace/plan/whmcs-mcp`  
> **Target Sistem**: `whmcs-mcp` (Go 1.22+)  
> **Tanggal**: 2026-09-30  

---

## 1. Aturan Mutlak Protokol Stdio MCP

Dalam arsitektur Model Context Protocol (MCP) dengan transport `stdio`:
1. **Kanal Standar Output (`os.Stdout`)** HANYA boleh berisi frame JSON-RPC 2.0 yang valid, satu pesan per baris (*newline-delimited JSON* atau *Content-Length framing*).
2. **Kanal Standar Error (`os.Stderr`)** adalah SATU-SATUNYA saluran untuk:
   - Pesan log aplikasi (debug, info, warn, error).
   - Traceback atau stack trace panic.
   - Diagnostik HTTP request/response.
3. **Pelanggaran**: Penulisan teks biasa (seperti `fmt.Println("Starting server...")`) ke `os.Stdout` akan merusak framing JSON-RPC dan menyebabkan klien LLM (seperti Claude Desktop atau Cursor) mengalami *crash* atau pemutusan koneksi seketika.

---

## 2. Implementasi Logging Aman di Go

Konfigurasi logging standar harus secara eksplisit mengarahkan output ke `os.Stderr`:

```go
package main

import (
    "log/slog"
    "os"
)

func initLogger(debug bool) *slog.Logger {
    level := slog.LevelInfo
    if debug {
        level = slog.LevelDebug
    }
    
    // Handler diarahkan ke os.Stderr
    handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
        Level: level,
    })
    return slog.New(handler)
}
```

---

## 3. Manajemen Konkurensi & Siklus Hidup Proses

1. **Context Cancellation**:
   Setiap HTTP request ke WHMCS API harus menerima `context.Context` yang dapat dibatalkan saat klien MCP menutup koneksi atau membatalkan request (*AbortSignal*):
   ```go
   req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL, body)
   ```
2. **Graceful Shutdown**:
   Aplikasi menangani sinyal OS `SIGINT` dan `SIGTERM` secara elegan:
   ```go
   ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
   defer stop()
   ```
3. **Goroutine Safety**:
   Koleksi handler tool MCP berjalan secara paralel bila klien mengirimkan beberapa RPC request secara serentak. Instans `WhmcsClient` bersifat thread-safe dan menggunakan koneksi pooling bawaan `http.Transport`.
