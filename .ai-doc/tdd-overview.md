# TDD Overview

> Pusat kontrol TDD project `whmcs-mcp` (Golang).
> Source of truth status adalah bukti eksekusi test nyata yang tercatat di sini.

## 1. Metadata

- **Project:** `whmcs-mcp` (Golang Implementation)
- **TDD Policy:** `Enabled`
- **Scope:** `Greenfield development`
- **Activated at:** `2026-09-30`
- **Constitution:** `.ai-doc/constitution.md`
- **Last updated:** `2026-09-30 09:45`
- **Overall status:** `Active`
- **Test Runner:** `go test -v ./...`

---

## 2. Progress Summary

| Metric | Count |
|---|---:|
| Total targets | 4 |
| PLANNED | 0 |
| RED | 0 |
| GREEN | 0 |
| REFACTORING | 0 |
| REFACTORED | 4 |
| BLOCKED | 0 |
| EXCEPTION | 0 |

---

## 3. TDD Registry

| ID | Component | Use Case / Behavior | Acceptance Criteria | Test File | Current Status | Last Evidence | Notes |
|---|---|---|---|---|---|---|---|
| **TDD-001** | `Config` | Memuat & memvalidasi konfigurasi env WHMCS | Mengembalikan struct `Config` valid bila env lengkap, error jika URL/kredensial kosong | `internal/config/config_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.013s) | Fondasi konfigurasi tuntas |
| **TDD-002** | `WHMCS Client` | Serialisasi parameter hierarki/array ke `url.Values` | Menghasilkan encoding `application/x-www-form-urlencoded` yang tepat untuk nested array WHMCS | `internal/whmcs/serializer_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.004s) | Serializer parameter tuntas |
| **TDD-003** | `WHMCS Client` | Eksekusi API call & unmarshaling JSON response | Mengirim POST ke mock WHMCS server, inject auth headers/body, dan parse `result: success\|error` | `internal/whmcs/client_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.013s) | HTTP Client core tuntas |
| **TDD-004** | `MCP Engine` | Registrasi tools & eksekusi stdio RPC dispatch | Server MCP mengenali list tool (Client, Billing, Ticket, dll.) dan memanggil handler yang tepat | `internal/mcp/server_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.010s) | Dispatcher protokol MCP tuntas |

---

## 4. Cycle Detail

### TDD-001 — Config Loading and Validation

- **Component:** `internal/config`
- **Use case source:** `.ai-doc/project-overview.md` §6.2
- **Acceptance criteria:**
  1. `LoadConfig()` membaca `WHMCS_URL`, `WHMCS_API_IDENTIFIER`, `WHMCS_API_SECRET`, `WHMCS_API_ACCESS_KEY`, `WHMCS_TIMEOUT`.
  2. Memberikan fallback default timeout (30s) jika tidak ditentukan.
  3. Mengembalikan error informatif jika `WHMCS_URL` tidak valid atau kredensial kosong.
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/config/config_test.go`
- **Test name/target:** `TestLoad_Success`, `TestLoad_DefaultTimeout`, `TestLoad_MissingURL`, `TestLoad_InvalidURL`, `TestLoad_MissingIdentifier`, `TestLoad_MissingSecret`
- **Command:** `go test -v ./internal/config`
- **Exit status:** `1`
- **Failure evidence:** `github.com/vfat/whmcs-mcp/internal/config: no non-test Go files in /home/ubuntu/workspace/plan/whmcs-mcp/internal/config`
- **Verified at:** `2026-09-30 10:31`

#### GREEN
- **Implementation file(s):** `internal/config/config.go`
- **Minimal change:** `Implement struct Config and Load() with environment parsing, url.ParseRequestURI validation, and DefaultTimeout fallback`
- **Command:** `go test -v ./internal/config`
- **Exit status:** `0`
- **Passing evidence:** `PASS: 6/6 tests passed (0.013s)`
- **Verified at:** `2026-09-30 10:31`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Add idiomatic doc comments, clean up imports, ensure trimmed strings for all inputs`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: ok github.com/vfat/whmcs-mcp/internal/config (cached)`
- **Verified at:** `2026-09-30 10:32`

---

### TDD-002 — Parameter Serializer untuk WHMCS Form URL-Encoded

- **Component:** `internal/whmcs`
- **Use case source:** `.ai-doc/project-overview.md` §7, `plan/whmcs-mcp-tool/src/whmcs-client.ts` (`flattenParams`)
- **Acceptance criteria:**
  1. Parameter skalar dipetakan langsung ke key-value string.
  2. Parameter array/slice (seperti `customfields` atau list ID) diserialisasi dengan pola indeks (misal: `customfields[0]`).
  3. Parameter map/nested struct diserialisasi dengan key nested (misal: `nameservers[ns1]`).
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/whmcs/serializer_test.go`
- **Test name/target:** `TestSerializeParams_Flat`, `TestSerializeParams_NestedMap`, `TestSerializeParams_Slice`, `TestSerializeParams_Struct`
- **Command:** `go test -v ./internal/whmcs`
- **Exit status:** `1`
- **Failure evidence:** `github.com/vfat/whmcs-mcp/internal/whmcs: no non-test Go files in /home/ubuntu/workspace/plan/whmcs-mcp/internal/whmcs`
- **Verified at:** `2026-09-30 10:40`

#### GREEN
- **Implementation file(s):** `internal/whmcs/serializer.go`
- **Minimal change:** `Implement SerializeParams with recursive reflection to handle maps, slices, and struct url tags`
- **Command:** `go test -v ./internal/whmcs`
- **Exit status:** `0`
- **Passing evidence:** `PASS: 4/4 tests passed (0.004s)`
- **Verified at:** `2026-09-30 10:41`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Modular scalar formatter, proper omitempty handling, clean export comments`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: ok github.com/vfat/whmcs-mcp/internal/config, ok github.com/vfat/whmcs-mcp/internal/whmcs`
- **Verified at:** `2026-09-30 10:41`

---

### TDD-003 — WHMCS API HTTP Client Execution

- **Component:** `internal/whmcs`
- **Use case source:** `.ai-doc/project-overview.md` §6.3, `plan/whmcs-mcp-tool/.ai-doc/rest-api-doc/`
- **Acceptance criteria:**
  1. Mengirim POST `application/x-www-form-urlencoded` ke `/includes/api.php` dengan parameter wajib `action`, `responsetype=json`, serta kredensial auth.
  2. Mengembalikan error terstruktur bila WHMCS merespon `result: error` (`message` diekstrak).
  3. Berhasil decode JSON payload response generic/spesifik pada `result: success`.
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/whmcs/client_test.go`
- **Test name/target:** `TestClient_ExecuteSuccess`, `TestClient_ExecuteApiError`, `TestClient_ExecuteHttpError`, `TestClient_ExecuteContextCanceled`
- **Command:** `go test ./internal/whmcs`
- **Exit status:** `1`
- **Failure evidence:** `internal/whmcs/client_test.go:74:18: undefined: whmcs.NewClient`
- **Verified at:** `2026-09-30 10:43`

#### GREEN
- **Implementation file(s):** `internal/whmcs/client.go`
- **Minimal change:** `Implement Client struct and Execute method with auth injection, context propagation, status check, and JSON decode`
- **Command:** `go test -v ./internal/whmcs`
- **Exit status:** `0`
- **Passing evidence:** `PASS: 8/8 tests passed (0.013s)`
- **Verified at:** `2026-09-30 10:43`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Add connection pooling in http.Transport, clean error wrapping, context cancellation support`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: ok github.com/vfat/whmcs-mcp/internal/config, ok github.com/vfat/whmcs-mcp/internal/whmcs (14/14 tests passing)`
- **Verified at:** `2026-09-30 10:43`

---

### TDD-004 — MCP Server Engine & Tool Dispatcher

- **Component:** `internal/mcp`
- **Use case source:** `.ai-doc/project-overview.md` §6.1, `plan/whmcs-mcp-tool/.ai-doc/lampiran/L-001`
- **Acceptance criteria:**
  1. Server terinisialisasi dengan metadata nama dan versi.
  2. Registrasi tool menghasilkan definisi skema JSON-RPC yang valid.
  3. Eksekusi tool memanggil handler terkait dan mengembalikan teks/JSON result format MCP.
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/mcp/server_test.go`
- **Test name/target:** `TestNewServer`, `TestRegisterTool_Execution`
- **Command:** `go test ./internal/mcp`
- **Exit status:** `1`
- **Failure evidence:** `github.com/vfat/whmcs-mcp/internal/mcp: no non-test Go files in /home/ubuntu/workspace/plan/whmcs-mcp/internal/mcp`
- **Verified at:** `2026-09-30 10:48`

#### GREEN
- **Implementation file(s):** `internal/mcp/server.go`
- **Minimal change:** `Implement Server struct with NewServer, MCPServer, Client, and ListTools wrapping mcp-go server`
- **Command:** `go test -v ./internal/mcp`
- **Exit status:** `0`
- **Passing evidence:** `PASS: 2/2 tests passed (0.010s)`
- **Verified at:** `2026-09-30 10:51`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Support stdio serve loop, safe type assertions on tool call arguments`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: ok github.com/vfat/whmcs-mcp/internal/config, ok github.com/vfat/whmcs-mcp/internal/whmcs, ok github.com/vfat/whmcs-mcp/internal/mcp (16/16 tests passing)`
- **Verified at:** `2026-09-30 10:54`

---

## 5. Blockers and Exceptions

*Tidak ada blocker aktif saat ini.*

---

## 6. Change Log

| Date | Target | Phase | Change | Evidence / Reference |
|---|---|---|---|---|
| `2026-09-30` | All | `Activation` | Inisialisasi TDD overview dan pendaftaran 4 target inti (TDD-001 s/d TDD-004) | `.ai-doc/constitution.md` |

---

## 7. Operating Rules

- Test ditulis sebelum production code untuk behavior baru.
- Status `RED` membutuhkan test yang gagal karena behavior belum ada, bukan karena typo/setup rusak.
- Status `GREEN` membutuhkan passing evidence setelah implementasi minimal.
- Status `REFACTORED` membutuhkan test terkait dan regression test tetap lulus.
- Jika command tidak dapat dijalankan, gunakan `BLOCKED` dan jelaskan alasannya.
- Update file ini dan `.ai-doc/3p.md` setelah setiap transisi bermakna.
