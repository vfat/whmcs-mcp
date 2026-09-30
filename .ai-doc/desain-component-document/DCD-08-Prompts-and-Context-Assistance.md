# DCD-08: Prompts and Context Assistance Component

> **Target Codebase**: `whmcs-mcp` (Go 1.22+)  
> **Workspace**: `/home/ubuntu/workspace/plan/whmcs-mcp`  
> **Target Package**: `internal/mcp/prompts/prompts.go`  
> **Tanggal**: 2026-09-30  

---

## 1. Overview

Komponen **Prompts and Context Assistance** menyediakan 8 templat instruksi agen sistem terstruktur yang dapat dipanggil oleh LLM Client melalui metode JSON-RPC `prompts/get`. Komponen ini menyusun konteks operasional multi-langkah (workflow) untuk memandu LLM menjalankan tugas-tugas administratif secara konsisten.

---

## 2. Object Identification (Go Architecture)

### 2.1 Boundary
- **MCP Prompt Handlers (`internal/mcp/prompts/prompts.go`)**:
  - `client-onboarding`, `client-health-check`
  - `ticket-response`, `revenue-report`, `bulk-invoice-reminder`
  - `domain-expiry-audit`, `fraud-investigation`, `new-product-setup`
- **Transport**: JSON-RPC `prompts/get` dan `prompts/list`.

### 2.2 Control
- **Prompt Formatter & Validator**:
  - Memvalidasi argumen yang diperlukan untuk setiap prompt (misal: `ticket_id`, `client_id`, `period`).
  - Menghasilkan payload pesan `mcp.PromptMessage` dengan peran `user` atau `assistant`.

### 2.3 Entity
- **Go Structs (`internal/mcp/prompts/`)**:
  - Definisi metadata `mcp.Prompt` (Name, Description, Arguments).
  - Teks instruksi templat terstruktur.

---

## 3. Use Case List

| No | Use Case Name | Prompt Name | Input Parameter |
|:---:|:---|:---|:---|
| 1 | Jalankan Orientasi Klien & Setup Produk | `client-onboarding`, `new-product-setup`, `client-health-check` | `client_name`, `email`, `package_name`, `domain`, `client_id` |
| 2 | Audit Keuangan & Penagihan Massal | `revenue-report`, `bulk-invoice-reminder` | `period`, `days_overdue` |
| 3 | Investigasi Fraud & Kedaluwarsa Domain | `fraud-investigation`, `domain-expiry-audit`, `ticket-response` | `order_id`, `days_until_expiry`, `ticket_id` |

---

## 4. Technical Flow & Business Rules

### 4.1 Eksekusi Prompt `prompts/get`
1. Klien mengirim request `prompts/get` dengan nama prompt dan nilai argumen.
2. Handler memeriksa keberadaan parameter wajib.
3. Templat pesan diisi dengan argumen dinamis menggunakan Go `text/template` atau string formatting aman.
4. Dikembalikan sebagai struktur `mcp.GetPromptResult` dengan daftar `mcp.PromptMessage`.
