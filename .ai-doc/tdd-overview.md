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
| Total targets | 19 |
| PLANNED | 0 |
| RED | 0 |
| GREEN | 0 |
| REFACTORING | 0 |
| REFACTORED | 19 |
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
| **TDD-010** | `Models` | Struct DTO Billing, Product & Tickets | Serialisasi model produk, invoice, transaksi, dan tiket bantuan pelanggan | `internal/model/billing_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.005s) | Kontrak DTO Batch 2 |
| **TDD-011** | `MCP Tools` | 9 Tools Billing & Product Handlers | Registrasi dan eksekusi katalog produk, invoice, mutasi pembayaran kas, dan kredit | `internal/mcp/tools/billing_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.046s) | Batch 2 Billing Tools |
| **TDD-012** | `MCP Tools` | 9 Tools Support Ticket Handlers | Registrasi dan eksekusi alur tiket bantuan, balasan staf, dan internal note | `internal/mcp/tools/tickets_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.043s) | Batch 2 Ticket Tools |
| **TDD-013** | `MCP Resources`| 3 Resources Batch 2 (`products`, `support`) | Resource URI katalog produk, departemen bantuan, dan status tiket | `internal/mcp/resources/batch2_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.020s) | Batch 2 Resources |
| **TDD-014** | `MCP Prompts`  | 3 Prompts Batch 2 | Prompt cerdas untuk tiket respon, analisis pendapatan, dan pengingat invoice jatuh tempo | `internal/mcp/prompts/batch2_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.011s) | Batch 2 Prompts |
| **TDD-015** | `Models` | Struct DTO Domain & Order Models | Serialisasi struct domain (whois, nameservers, lock) dan order/quotes | `internal/model/domain_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.017s) | Kontrak DTO Batch 3 |
| **TDD-016** | `MCP Tools` | 9 Tools Domain Management Handlers | Registrasi dan eksekusi register, transfer, renew, whois, nameservers, locking, TLD | `internal/mcp/tools/domains_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.038s) | Batch 3 Domain Tools |
| **TDD-017** | `MCP Tools` | 10 Tools Orders & Quotes Handlers | Registrasi dan eksekusi order lifecycle (accept, cancel, delete, fraud, pending) & quotes | `internal/mcp/tools/orders_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.040s) | Batch 3 Order Tools |
| **TDD-018** | `MCP Resources`| Resource `whmcs://tld-pricing` | Resource URI katalog daftar harga TLD registrar aktif | `internal/mcp/resources/batch3_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.015s) | Batch 3 Resources |
| **TDD-019** | `MCP Prompts`  | 2 Prompts Batch 3 (`domain-expiry`, `fraud`) | Prompt cerdas audit kedaluwarsa domain dan investigasi pesanan terindikasi fraud | `internal/mcp/prompts/batch3_test.go` | `REFACTORED` | `go test -v ./...` PASS (0.012s) | Batch 3 Prompts |

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

### TDD-010 — Domain DTO Models (Billing, Products & Support Tickets)

- **Component:** `internal/model`
- **Use case source:** `.ai-doc/Dokumentasi-Fitur.md` §2.B, §2.C, §2.D
- **Acceptance criteria:**
  1. Struct request mendefinisikan tag `json` dan `url` (termasuk multidimensi array line item `itemdescription[0]` dsb.).
  2. Struct response unmarshaling JSON WHMCS envelope (`invoices`, `products`, `tickets`, `departments`, `statuses`).
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/model/billing_test.go`, `internal/model/ticket_test.go`
- **Test name/target:** `TestBillingModels_Serialization`, `TestProductModels_Serialization`, `TestTicketModels_Serialization`
- **Command:** `go test -v ./internal/model`
- **Exit status:** `1`
- **Failure evidence:** `undefined: model.CreateInvoiceRequest, model.GetInvoicesResponse, model.OpenTicketRequest`
- **Verified at:** `2026-09-30 12:03`

#### GREEN
- **Implementation file(s):** `internal/model/product.go`, `internal/model/billing.go`, `internal/model/ticket.go`
- **Minimal change:** `Implement product, billing, and ticket structs with json and url struct tags`
- **Command:** `go test -v ./internal/model`
- **Exit status:** `0`
- **Passing evidence:** `PASS: 7/7 model tests passed (0.005s)`
- **Verified at:** `2026-09-30 12:04`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Support strong types for invoice line items and ticket note attachments`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: All models verified`
- **Verified at:** `2026-09-30 12:04`

---

### TDD-011 — Billing & Product Tools Handlers (9 Tools)

- **Component:** `internal/mcp/tools`
- **Use case source:** `.ai-doc/Dokumentasi-Fitur.md` §2.B & §2.C, `plan/whmcs-mcp-tool/src/index.ts`
- **Acceptance criteria:**
  1. Mendaftarkan 9 tools: `whmcs_get_products`, `whmcs_get_product_groups`, `whmcs_get_invoices`, `whmcs_get_invoice`, `whmcs_create_invoice`, `whmcs_update_invoice`, `whmcs_add_payment`, `whmcs_apply_credit`, `whmcs_get_transactions`.
  2. Validasi input arguments wajib (`invoiceid`, `userid`, `transid`, `gateway`, `amount`).
  3. Memanggil action API WHMCS yang tepat dan mengembalikan tool result.
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/mcp/tools/products_test.go`, `internal/mcp/tools/billing_test.go`
- **Test name/target:** `TestRegisterBillingTools_AllPresent`, `TestBillingTools_Execution`, `TestRegisterProductTools_AllPresent`, `TestProductTools_Execution`
- **Command:** `go test -v ./internal/mcp/tools`
- **Exit status:** `1`
- **Failure evidence:** `undefined: tools.RegisterBillingTools, tools.RegisterProductTools`
- **Verified at:** `2026-09-30 12:04`

#### GREEN
- **Implementation file(s):** `internal/mcp/tools/products.go`, `internal/mcp/tools/billing.go`
- **Minimal change:** `Implement RegisterProductTools and RegisterBillingTools with validation and API dispatching`
- **Command:** `go test -v ./internal/mcp/tools`
- **Exit status:** `0`
- **Passing evidence:** `PASS: 19/19 test cases in internal/mcp/tools passed (0.046s)`
- **Verified at:** `2026-09-30 12:05`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Unified formatJSONResult helper across tools`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: All tool tests pass`
- **Verified at:** `2026-09-30 12:05`

---

### TDD-012 — Support Ticket Tools Handlers (9 Tools)

- **Component:** `internal/mcp/tools`
- **Use case source:** `.ai-doc/Dokumentasi-Fitur.md` §2.D, `plan/whmcs-mcp-tool/src/index.ts`
- **Acceptance criteria:**
  1. Mendaftarkan 9 tools: `whmcs_get_tickets`, `whmcs_get_ticket`, `whmcs_open_ticket`, `whmcs_add_ticket_reply`, `whmcs_add_ticket_note`, `whmcs_update_ticket`, `whmcs_delete_ticket`, `whmcs_get_support_departments`, `whmcs_get_support_statuses`.
  2. Validasi input arguments wajib (`ticketid`, `deptid`, `subject`, `message`).
  3. Memanggil action API WHMCS dan mengembalikan JSON tool result.
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/mcp/tools/tickets_test.go`
- **Test name/target:** `TestRegisterTicketTools_AllPresent`, `TestTicketTools_Execution`
- **Command:** `go test -v ./internal/mcp/tools`
- **Exit status:** `1`
- **Failure evidence:** `undefined: tools.RegisterTicketTools`
- **Verified at:** `2026-09-30 12:05`

#### GREEN
- **Implementation file(s):** `internal/mcp/tools/tickets.go`
- **Minimal change:** `Implement RegisterTicketTools with all 9 support ticket tool definitions and handlers`
- **Command:** `go test -v ./internal/mcp/tools`
- **Exit status:** `0`
- **Passing evidence:** `PASS: 14/14 subtests in TestTicketTools_Execution passed (0.043s)`
- **Verified at:** `2026-09-30 12:05`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Clean validation checks for note content and reply payloads`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: 28/28 tools across clients, system, billing, products, tickets verified`
- **Verified at:** `2026-09-30 12:06`

---

### TDD-013 — Batch 2 MCP Resources (3 Resources)

- **Component:** `internal/mcp/resources`
- **Use case source:** `.ai-doc/Dokumentasi-Fitur.md` §2.B & §2.D, `plan/whmcs-mcp-tool/src/index.ts`
- **Acceptance criteria:**
  1. Mendaftarkan 3 resources: `whmcs://products`, `whmcs://support/departments`, `whmcs://support/statuses`.
  2. Handler mengeksekusi API call dan mengembalikan `TextResourceContents` JSON.
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/mcp/resources/batch2_test.go`
- **Test name/target:** `TestRegisterBatch2Resources_AllPresent`, `TestBatch2Resources_Read`
- **Command:** `go test -v ./internal/mcp/resources`
- **Exit status:** `1`
- **Failure evidence:** `undefined: resources.RegisterBatch2Resources`
- **Verified at:** `2026-09-30 12:06`

#### GREEN
- **Implementation file(s):** `internal/mcp/resources/batch2.go`
- **Minimal change:** `Implement RegisterBatch2Resources with product and support resource definitions`
- **Command:** `go test -v ./internal/mcp/resources`
- **Exit status:** `0`
- **Passing evidence:** `PASS: 8/8 resources across Batch 1 and Batch 2 passed (0.020s)`
- **Verified at:** `2026-09-30 12:06`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Shared formatResourceJSON error recovery`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: ok github.com/vfat/whmcs-mcp/internal/mcp/resources`
- **Verified at:** `2026-09-30 12:06`

---

### TDD-014 — Batch 2 MCP Prompts (3 Prompts)

- **Component:** `internal/mcp/prompts`
- **Use case source:** `.ai-doc/Dokumentasi-Fitur.md` §2.J, `plan/whmcs-mcp-tool/src/index.ts`
- **Acceptance criteria:**
  1. Mendaftarkan 3 prompts: `ticket-response`, `revenue-report`, `bulk-invoice-reminder`.
  2. Validasi argumen wajib (`ticketId`, `issueType`, `period`, `daysOverdue`).
  3. Memformulasikan panduan interaktif komprehensif.
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/mcp/prompts/batch2_test.go`
- **Test name/target:** `TestRegisterBatch2Prompts_AllPresent`, `TestBatch2Prompts_Execution`
- **Command:** `go test -v ./internal/mcp/prompts`
- **Exit status:** `1`
- **Failure evidence:** `undefined: prompts.RegisterBatch2Prompts`
- **Verified at:** `2026-09-30 12:06`

#### GREEN
- **Implementation file(s):** `internal/mcp/prompts/batch2.go`
- **Minimal change:** `Implement RegisterBatch2Prompts with prompt definitions and message builders`
- **Command:** `go test -v ./internal/mcp/prompts`
- **Exit status:** `0`
- **Passing evidence:** `PASS: 5/5 prompts across Batch 1 and Batch 2 passed (0.011s)`
- **Verified at:** `2026-09-30 12:07`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Composition root wiring in cmd/whmcs-mcp/main.go updated with all Batch 2 components`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: 52/52 tests passed across entire repository`
- **Verified at:** `2026-09-30 12:07`

---

### TDD-015 — Domain DTO Models: Domains, Orders & Quotes

- **Component:** `internal/model`
- **Use case source:** `.ai-doc/Dokumentasi-Fitur.md` §E & §F, `plan/whmcs-mcp-tool/src/whmcs-client.ts`
- **Acceptance criteria:**
  1. Struct Request/Response untuk Domain: `DomainWhois`, `DomainRegister`, `DomainTransfer`, `DomainRenew`, `DomainGetNameservers`, `DomainUpdateNameservers`, `DomainGetLockingStatus`, `DomainUpdateLockingStatus`, `GetTLDPricing`.
  2. Struct Request/Response untuk Orders & Quotes: `GetOrders`, `AcceptOrder`, `CancelOrder`, `DeleteOrder`, `FraudOrder`, `PendingOrder`, `GetQuotes`, `CreateQuote`, `AcceptQuote`, `DeleteQuote`.
  3. Validasi serialisasi dan parsing JSON roundtrip.
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/model/domain_test.go`, `internal/model/order_test.go`
- **Test name/target:** `TestDomainModels_Serialization`, `TestOrderModels_Serialization`
- **Command:** `go test -v ./internal/model`
- **Exit status:** `1`
- **Failure evidence:** `undefined: model.DomainRegisterRequest`
- **Verified at:** `2026-09-30 12:10`

#### GREEN
- **Implementation file(s):** `internal/model/domain.go`, `internal/model/order.go`
- **Minimal change:** `Define structs with url and json tags for all Batch 3 operations`
- **Command:** `go test -v ./internal/model`
- **Exit status:** `0`
- **Passing evidence:** `PASS: 10/10 tests passed (0.017s)`
- **Verified at:** `2026-09-30 12:11`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Add godoc comments and ensure pointer semantics on optional query parameters`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: ok github.com/vfat/whmcs-mcp/internal/model`
- **Verified at:** `2026-09-30 12:12`

---

### TDD-016 — 9 Tools Domain Management Handlers

- **Component:** `internal/mcp/tools`
- **Use case source:** `.ai-doc/Dokumentasi-Fitur.md` §E
- **Acceptance criteria:**
  1. Registrasi 9 tool: `whmcs_register_domain`, `whmcs_transfer_domain`, `whmcs_renew_domain`, `whmcs_get_domain_whois`, `whmcs_get_domain_nameservers`, `whmcs_update_domain_nameservers`, `whmcs_get_domain_lock_status`, `whmcs_update_domain_lock_status`, `whmcs_get_tld_pricing`.
  2. Handler mengeksekusi aksi WHMCS terkait dan merespons format JSON mcp TextContent.
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/mcp/tools/domains_test.go`
- **Test name/target:** `TestDomainTools_Execution`
- **Command:** `go test -v ./internal/mcp/tools -run TestDomainTools`
- **Exit status:** `1`
- **Failure evidence:** `undefined: RegisterDomainTools`
- **Verified at:** `2026-09-30 12:11`

#### GREEN
- **Implementation file(s):** `internal/mcp/tools/domains.go`
- **Minimal change:** `Implement RegisterDomainTools and all 9 tool handlers`
- **Command:** `go test -v ./internal/mcp/tools -run TestDomainTools`
- **Exit status:** `0`
- **Passing evidence:** `PASS: 9/9 domain tools tests passed (0.038s)`
- **Verified at:** `2026-09-30 12:11`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Clean up arguments unmarshaling via parseArguments helper`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: ok github.com/vfat/whmcs-mcp/internal/mcp/tools`
- **Verified at:** `2026-09-30 12:12`

---

### TDD-017 — 10 Tools Orders & Quotes Handlers

- **Component:** `internal/mcp/tools`
- **Use case source:** `.ai-doc/Dokumentasi-Fitur.md` §F
- **Acceptance criteria:**
  1. Registrasi 10 tool: `whmcs_get_orders`, `whmcs_accept_order`, `whmcs_cancel_order`, `whmcs_delete_order`, `whmcs_fraud_order`, `whmcs_pending_order`, `whmcs_get_quotes`, `whmcs_create_quote`, `whmcs_accept_quote`, `whmcs_delete_quote`.
  2. Handler mengeksekusi aksi WHMCS terkait dan merespons format JSON mcp TextContent.
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/mcp/tools/orders_test.go`
- **Test name/target:** `TestOrderTools_Execution`
- **Command:** `go test -v ./internal/mcp/tools -run TestOrderTools`
- **Exit status:** `1`
- **Failure evidence:** `undefined: RegisterOrderTools`
- **Verified at:** `2026-09-30 12:11`

#### GREEN
- **Implementation file(s):** `internal/mcp/tools/orders.go`
- **Minimal change:** `Implement RegisterOrderTools and all 10 tool handlers`
- **Command:** `go test -v ./internal/mcp/tools -run TestOrderTools`
- **Exit status:** `0`
- **Passing evidence:** `PASS: 10/10 order & quote tools tests passed (0.040s)`
- **Verified at:** `2026-09-30 12:11`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Ensure safe type conversions and structured descriptions`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: ok github.com/vfat/whmcs-mcp/internal/mcp/tools`
- **Verified at:** `2026-09-30 12:12`

---

### TDD-018 — Resource `whmcs://tld-pricing`

- **Component:** `internal/mcp/resources`
- **Use case source:** `.ai-doc/Dokumentasi-Fitur.md` §E
- **Acceptance criteria:**
  1. Handler resource membaca `GetTLDPricing` dari WHMCS API.
  2. Mengembalikan data dalam MIME `application/json`.
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/mcp/resources/batch3_test.go`
- **Test name/target:** `TestBatch3Resources_Execution`
- **Command:** `go test -v ./internal/mcp/resources -run TestBatch3Resources`
- **Exit status:** `1`
- **Failure evidence:** `undefined: RegisterBatch3Resources`
- **Verified at:** `2026-09-30 12:11`

#### GREEN
- **Implementation file(s):** `internal/mcp/resources/batch3.go`
- **Minimal change:** `Implement RegisterBatch3Resources for whmcs://tld-pricing`
- **Command:** `go test -v ./internal/mcp/resources -run TestBatch3Resources`
- **Exit status:** `0`
- **Passing evidence:** `PASS: 2/2 batch 3 resource test cases passed (0.015s)`
- **Verified at:** `2026-09-30 12:11`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Shared JSON response formatter with Batch 1/2`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: ok github.com/vfat/whmcs-mcp/internal/mcp/resources`
- **Verified at:** `2026-09-30 12:12`

---

### TDD-019 — 2 Prompts Batch 3: Domain Expiry Audit & Fraud Investigation

- **Component:** `internal/mcp/prompts`
- **Use case source:** `.ai-doc/Dokumentasi-Fitur.md` §J
- **Acceptance criteria:**
  1. Prompt `domain-expiry-audit` memformat instruksi audit kedaluwarsa domain.
  2. Prompt `fraud-investigation` memformat instruksi investigasi keamanan pesanan.
- **Current status:** `REFACTORED`

#### RED
- **Test file:** `internal/mcp/prompts/batch3_test.go`
- **Test name/target:** `TestBatch3Prompts_Execution`
- **Command:** `go test -v ./internal/mcp/prompts -run TestBatch3Prompts`
- **Exit status:** `1`
- **Failure evidence:** `undefined: RegisterBatch3Prompts`
- **Verified at:** `2026-09-30 12:11`

#### GREEN
- **Implementation file(s):** `internal/mcp/prompts/batch3.go`
- **Minimal change:** `Implement RegisterBatch3Prompts with prompt definitions and message builders`
- **Command:** `go test -v ./internal/mcp/prompts -run TestBatch3Prompts`
- **Exit status:** `0`
- **Passing evidence:** `PASS: 3/3 prompt test cases passed (0.012s)`
- **Verified at:** `2026-09-30 12:11`

#### REFACTOR
- **Status:** `REFACTORED`
- **Changes:** `Composition root wiring in cmd/whmcs-mcp/main.go updated with all Batch 3 components`
- **Regression command:** `go test -v ./...`
- **Exit status:** `0`
- **Regression evidence:** `PASS: 137/137 tests passed across entire repository`
- **Verified at:** `2026-09-30 12:14`

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
| `2026-09-30` | TDD-010 | `Complete` | Domain DTO Models Billing, Product & Ticket | `internal/model/billing_test.go` PASS |
| `2026-09-30` | TDD-011 | `Complete` | 9 Tools Billing & Product Handlers | `internal/mcp/tools/billing_test.go` PASS |
| `2026-09-30` | TDD-012 | `Complete` | 9 Tools Support Ticket Handlers | `internal/mcp/tools/tickets_test.go` PASS |
| `2026-09-30` | TDD-013 | `Complete` | 3 Resources Batch 2 (`products`, `support`) | `internal/mcp/resources/batch2_test.go` PASS |
| `2026-09-30` | TDD-014 | `Complete` | 3 Prompts Batch 2 (`ticket-response`, `revenue-report`, `bulk-invoice-reminder`) | `internal/mcp/prompts/batch2_test.go` PASS |
| `2026-09-30` | TDD-015 | `Complete` | Domain DTO Models Domain & Order | `internal/model/domain_test.go` PASS |
| `2026-09-30` | TDD-016 | `Complete` | 9 Tools Domain Management Handlers | `internal/mcp/tools/domains_test.go` PASS |
| `2026-09-30` | TDD-017 | `Complete` | 10 Tools Orders & Quotes Handlers | `internal/mcp/tools/orders_test.go` PASS |
| `2026-09-30` | TDD-018 | `Complete` | Resource `whmcs://tld-pricing` | `internal/mcp/resources/batch3_test.go` PASS |
| `2026-09-30` | TDD-019 | `Complete` | 2 Prompts Batch 3 (`domain-expiry-audit`, `fraud-investigation`) | `internal/mcp/prompts/batch3_test.go` PASS |

---

## 7. Operating Rules

- Test ditulis sebelum production code untuk behavior baru.
- Status `RED` membutuhkan test yang gagal karena behavior belum ada, bukan karena typo/setup rusak.
- Status `GREEN` membutuhkan passing evidence setelah implementasi minimal.
- Status `REFACTORED` membutuhkan test terkait dan regression test tetap lulus.
- Jika command tidak dapat dijalankan, gunakan `BLOCKED` dan jelaskan alasannya.
- Update file ini dan `.ai-doc/3p.md` setelah setiap transisi bermakna.
