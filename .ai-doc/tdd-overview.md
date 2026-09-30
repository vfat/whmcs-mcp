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
| Total targets | 9 |
| PLANNED | 0 |
| RED | 0 |
| GREEN | 0 |
| REFACTORING | 0 |
| REFACTORED | 9 |
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
| **TDD-005** | `Models` | Struct DTO Client & System Models | Serialisasi struct request dengan tag `url` dan parsing JSON response dari 14 action WHMCS | `internal/model/client_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.002s) | Kontrak DTO Batch 1 |
| **TDD-006** | `MCP Tools` | 8 Tools Client Management Handlers | Registrasi dan eksekusi 8 tools klien (`whmcs_get_clients`, `whmcs_add_client`, dsb.) ke MCP | `internal/mcp/tools/clients_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.031s) | Batch 1 Client Tools |
| **TDD-007** | `MCP Tools` | 9 Tools System Administration Handlers | Registrasi dan eksekusi tools sistem (`whmcs_get_stats`, `whmcs_get_admin_users`, dsb.) | `internal/mcp/tools/system_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.029s) | Batch 1 System Tools |
| **TDD-008** | `MCP Resources`| 5 Resources Skema `whmcs://*` | Membaca resource data sistem instan (`stats`, `admin-users`, `currencies`, dll.) | `internal/mcp/resources/resources_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.021s) | Batch 1 Resources |
| **TDD-009** | `MCP Prompts`  | 2 Prompts Interaktif Klien | Menghasilkan instruksi terstruktur untuk onboarding dan audit kesehatan klien | `internal/mcp/prompts/prompts_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.011s) | Batch 1 Prompts |

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

### TDD-005 — Domain DTO Models (Client & System)

- **Component:** `internal/model`
- **Use case source:** `.ai-doc/Dokumentasi-Fitur.md` §2.A & §2.H, `.ai-doc/katalog-dto-endpoint.md`
- **Acceptance criteria:**
  1. Struct request mendefinisikan tag `json` dan `url` secara akurat sesuai form parameters WHMCS API.
  2. Struct response mendefinisikan tag `json` untuk unmarshaling payload WHMCS (envelope `clients`, `products`, `domains`, dll.).
  3. Embedding `CommonResponse` untuk menangkap status `result: "success" | "error"`.
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/model/client_test.go`, `internal/model/system_test.go`
- **Test name/target:** `TestClientModels_GetClientsSerialization`, `TestClientModels_AddClientSerialization`, `TestSystemModels_GetStatsResponse`, `TestSystemModels_SendEmailSerialization`
- **Command:** `go test -v ./internal/model`
- **Exit status:** `1`
- **Failure evidence:** `github.com/vfat/whmcs-mcp/internal/model: no non-test Go files`
- **Verified at:** `2026-09-30 10:58`

#### GREEN
- **Implementation file(s):** `internal/model/common.go`, `internal/model/client.go`, `internal/model/system.go`
- **Minimal change:** `Implement client request/response structs, system request/response structs, and common response envelope`
- **Command:** `go test -v ./internal/model`
- **Exit status:** `0`
- **Passing evidence:** `PASS: 4/4 tests passed (0.002s)`
- **Verified at:** `2026-09-30 11:00`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Enrich GetActivityLogRequest and GetToDoItemsRequest with complete WHMCS API parameters; add GetClientInvoicesRequest`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: 20/20 tests passed across config, whmcs, mcp, model`
- **Verified at:** `2026-09-30 11:13`

---

### TDD-006 — Client Management Tools Handlers (8 Tools)

- **Component:** `internal/mcp/tools`
- **Use case source:** `.ai-doc/Dokumentasi-Fitur.md` §2.A, `plan/whmcs-mcp-tool/src/index.ts`
- **Acceptance criteria:**
  1. Mendaftarkan 8 tools: `whmcs_get_clients`, `whmcs_get_client_details`, `whmcs_add_client`, `whmcs_update_client`, `whmcs_close_client`, `whmcs_get_client_products`, `whmcs_get_client_domains`, `whmcs_get_client_invoices`.
  2. Validasi input arguments (e.g., `clientid` mandatory pada invoice/close, first/last name required pada add client).
  3. Memanggil WHMCS Client dengan parameter terserialisasi dan mengembalikan CallToolResult berformat JSON.
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/mcp/tools/clients_test.go`
- **Test name/target:** `TestRegisterClientTools_AllToolsPresent`, `TestClientTools_GetClientsCall`, `TestClientTools_GetClientDetails`, `TestClientTools_CRUDAndSubResources`
- **Command:** `go test -v ./internal/mcp/tools`
- **Exit status:** `1`
- **Failure evidence:** `github.com/vfat/whmcs-mcp/internal/mcp/tools: no non-test Go files`
- **Verified at:** `2026-09-30 11:02`

#### GREEN
- **Implementation file(s):** `internal/mcp/tools/clients.go`
- **Minimal change:** `Implement RegisterClientTools registering all 8 client tools with JSON schema and tool execution handlers`
- **Command:** `go test -v ./internal/mcp/tools`
- **Exit status:** `0`
- **Passing evidence:** `PASS: TestRegisterClientTools_AllToolsPresent, TestClientTools_GetClientsCall, TestClientTools_GetClientDetails, TestClientTools_CRUDAndSubResources (0.031s)`
- **Verified at:** `2026-09-30 11:12`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Refactor whmcs_get_client_invoices to use GetClientInvoicesRequest with url:"userid" tag for type-safe parameter serialization`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: 28/28 tests passed across all packages`
- **Verified at:** `2026-09-30 11:13`

---

### TDD-007 — System Administration Tools Handlers (9 Tools)

- **Component:** `internal/mcp/tools`
- **Use case source:** `.ai-doc/Dokumentasi-Fitur.md` §2.H, `plan/whmcs-mcp-tool/src/index.ts`
- **Acceptance criteria:**
  1. Mendaftarkan 9 tools: `whmcs_get_stats`, `whmcs_get_activity_log`, `whmcs_get_admin_users`, `whmcs_get_todo_items`, `whmcs_get_todo_item_statuses`, `whmcs_update_todo_item`, `whmcs_get_currencies`, `whmcs_get_payment_methods`, `whmcs_send_email`.
  2. Validasi input arguments (e.g. `itemid` wajib pada `whmcs_update_todo_item`).
  3. Memanggil action API WHMCS yang bersesuaian dan mengembalikan JSON tool result.
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/mcp/tools/system_test.go`
- **Test name/target:** `TestRegisterSystemTools_AllToolsPresent`, `TestSystemTools_Execution`
- **Command:** `go test -v ./internal/mcp/tools`
- **Exit status:** `1`
- **Failure evidence:** `internal/mcp/tools/system_test.go: undefined: tools.RegisterSystemTools`
- **Verified at:** `2026-09-30 11:14`

#### GREEN
- **Implementation file(s):** `internal/mcp/tools/system.go`
- **Minimal change:** `Implement RegisterSystemTools with 9 system tool definitions and execution handlers`
- **Command:** `go test -v ./internal/mcp/tools`
- **Exit status:** `0`
- **Passing evidence:** `PASS: 10/10 test functions passed in internal/mcp/tools (0.029s)`
- **Verified at:** `2026-09-30 11:15`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Clean formatJSONResult error handling and parameter unmarshaling`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: All tool and model tests pass`
- **Verified at:** `2026-09-30 11:16`

---

### TDD-008 — Batch 1 MCP Resources (5 Resources)

- **Component:** `internal/mcp/resources`
- **Use case source:** `.ai-doc/Dokumentasi-Fitur.md` §2.H, `plan/whmcs-mcp-tool/src/index.ts`
- **Acceptance criteria:**
  1. Mendaftarkan 5 resources: `whmcs://stats`, `whmcs://admin-users`, `whmcs://currencies`, `whmcs://payment-methods`, `whmcs://admin/todo`.
  2. Handler membaca URI spesifik, mengeksekusi WHMCS API call, dan mengembalikan `TextResourceContents` dengan MIME `application/json`.
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/mcp/resources/resources_test.go`
- **Test name/target:** `TestRegisterBatch1Resources_AllPresent`, `TestBatch1Resources_Read`
- **Command:** `go test -v ./internal/mcp/resources`
- **Exit status:** `1`
- **Failure evidence:** `github.com/vfat/whmcs-mcp/internal/mcp/resources: no non-test Go files`
- **Verified at:** `2026-09-30 11:47`

#### GREEN
- **Implementation file(s):** `internal/mcp/resources/resources.go`
- **Minimal change:** `Implement RegisterBatch1Resources with resource declarations and read handlers returning JSON text contents`
- **Command:** `go test -v ./internal/mcp/resources`
- **Exit status:** `0`
- **Passing evidence:** `PASS: TestRegisterBatch1Resources_AllPresent, TestBatch1Resources_Read (0.021s)`
- **Verified at:** `2026-09-30 11:48`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Shared formatResourceJSON helper with fallback error JSON for resilient read operations`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: 36/36 tests passed across the repository`
- **Verified at:** `2026-09-30 11:50`

---

### TDD-009 — Batch 1 MCP Prompts (2 Prompts)

- **Component:** `internal/mcp/prompts`
- **Use case source:** `.ai-doc/Dokumentasi-Fitur.md` §2.J, `plan/whmcs-mcp-tool/src/index.ts`
- **Acceptance criteria:**
  1. Mendaftarkan 2 prompts: `client-onboarding` dan `client-health-check`.
  2. Memvalidasi argumen wajib (`clientName`, `email`, `clientId`).
  3. Memformulasikan pesan prompt panduan terstruktur langkah demi langkah.
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/mcp/prompts/prompts_test.go`
- **Test name/target:** `TestRegisterBatch1Prompts_AllPresent`, `TestBatch1Prompts_Execution`
- **Command:** `go test -v ./internal/mcp/prompts`
- **Exit status:** `1`
- **Failure evidence:** `github.com/vfat/whmcs-mcp/internal/mcp/prompts: no non-test Go files`
- **Verified at:** `2026-09-30 11:50`

#### GREEN
- **Implementation file(s):** `internal/mcp/prompts/prompts.go`
- **Minimal change:** `Implement RegisterBatch1Prompts with argument declarations and prompt message builders`
- **Command:** `go test -v ./internal/mcp/prompts`
- **Exit status:** `0`
- **Passing evidence:** `PASS: TestRegisterBatch1Prompts_AllPresent, TestBatch1Prompts_Execution (0.011s)`
- **Verified at:** `2026-09-30 11:51`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Clean string formatting with strings.Builder and complete composition root wiring in cmd/whmcs-mcp/main.go`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: 39/39 tests passed across all packages`
- **Verified at:** `2026-09-30 11:52`

---

## 5. Blockers and Exceptions

*Tidak ada blocker aktif saat ini.*

---

## 6. Change Log

| Date | Target | Phase | Change | Evidence / Reference |
|---|---|---|---|---|
| `2026-09-30` | All | `Activation` | Inisialisasi TDD overview dan pendaftaran 4 target inti (TDD-001 s/d TDD-004) | `.ai-doc/constitution.md` |
| `2026-09-30` | TDD-005 | `Complete` | Domain DTO Models untuk Client & System | `internal/model/client_test.go` PASS |
| `2026-09-30` | TDD-006 | `Complete` | 8 Tools Client Management Handlers | `internal/mcp/tools/clients_test.go` PASS |
| `2026-09-30` | TDD-007 | `Complete` | 9 Tools System Administration Handlers | `internal/mcp/tools/system_test.go` PASS |
| `2026-09-30` | TDD-008 | `Complete` | 5 Resources Skema `whmcs://*` | `internal/mcp/resources/resources_test.go` PASS |
| `2026-09-30` | TDD-009 | `Complete` | 2 Prompts Interaktif Klien | `internal/mcp/prompts/prompts_test.go` PASS |

---

## 7. Operating Rules

- Test ditulis sebelum production code untuk behavior baru.
- Status `RED` membutuhkan test yang gagal karena behavior belum ada, bukan karena typo/setup rusak.
- Status `GREEN` membutuhkan passing evidence setelah implementasi minimal.
- Status `REFACTORED` membutuhkan test terkait dan regression test tetap lulus.
- Jika command tidak dapat dijalankan, gunakan `BLOCKED` dan jelaskan alasannya.
- Update file ini dan `.ai-doc/3p.md` setelah setiap transisi bermakna.
