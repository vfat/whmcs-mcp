# DCD-07: System Administration Component

> **Target Codebase**: `whmcs-mcp` (Go 1.22+)  
> **Workspace**: `/home/ubuntu/workspace/plan/whmcs-mcp`  
> **Target Package**: `internal/mcp/tools/system.go`, `internal/model/system.go`, `internal/model/marketing.go`  
> **Tanggal**: 2026-09-30  

---

## 1. Overview

Komponen **System Administration** menyediakan fungsi utilitas operasional harian, audit keamanan log aktivitas, manajemen tugas staf (*to-do items*), pengawasan staf administrator, pengiriman email sistem, serta manajemen program kemitraan (*affiliates*) dan kupon promosi.

---

## 2. Object Identification (Go Architecture)

### 2.1 Boundary
- **MCP Tool Handlers (`internal/mcp/tools/system.go`, `internal/mcp/tools/marketing.go`)**:
  - `whmcs_get_stats`, `whmcs_get_activity_log`, `whmcs_get_admin_users`
  - `whmcs_get_todo_items`, `whmcs_get_todo_item_statuses`, `whmcs_update_todo_item`
  - `whmcs_get_currencies`, `whmcs_get_payment_methods`, `whmcs_send_email`
  - `whmcs_get_affiliates`, `whmcs_affiliate_activate`, `whmcs_get_promotions`
- **MCP Resource Handlers (`internal/mcp/resources/system.go`)**:
  - `whmcs://stats`, `whmcs://admin-users`, `whmcs://admin/todo`, `whmcs://promotions`

### 2.2 Control
- **WHMCS Client Dispatcher (`internal/whmcs/client.go`)**:
  - `GetStats`, `GetActivityLog`, `LogActivity`, `GetAdminUsers`
  - `GetToDoItems`, `GetToDoItemStatuses`, `UpdateToDoItem`
  - `GetCurrencies`, `GetPaymentMethods`, `SendEmail`, `GetEmailTemplates`
  - `GetAffiliates`, `AffiliateActivate`, `GetPromotions`

### 2.3 Entity
- **Go Structs (`internal/model/system.go`, `internal/model/marketing.go`)**:
  - `Stats`, `ActivityLog`, `AdminUser`, `ToDoItem`, `Affiliate`, `Promotion`
  - Request & Response DTOs.
- **Remote WHMCS Tables**: `tblactivitylog`, `tbladmins`, `tbltodolist`, `tblaffiliates`, `tblpromotions`.

---

## 3. Use Case List

| No | Use Case Name | Action WHMCS | Sifat | Go Function |
|:---:|:---|:---:|:---:|:---|
| 1 | Baca Statistik & Log Audit Aktivitas | `GetStats`, `GetActivityLog`, `GetAdminUsers`, `GetCurrencies`, `GetPaymentMethods` | Safe | `GetStats`, `GetActivityLog`, `GetAdminUsers`, `GetCurrencies`, `GetPaymentMethods` |
| 2 | Kelola Tugas Administratif (To-Do) | `GetToDoItems`, `GetToDoItemStatuses`, `UpdateToDoItem` | Safe | `GetToDoItems`, `GetToDoItemStatuses`, `UpdateToDoItem` |
| 3 | Kirim Email Notifikasi Langsung | `SendEmail` | Safe | `SendEmail` |
| 4 | Kelola Afiliasi & Promosi | `GetAffiliates`, `AffiliateActivate`, `GetPromotions` | Safe | `GetAffiliates`, `AffiliateActivate`, `GetPromotions` |

---

## 4. Technical Flow & Business Rules

### 4.1 System Stats (GetStats)
- Mengembalikan metrik snapshot pendapatan hari ini, bulan ini, total unpaid invoices, tiket aktif, dan pending orders.
- Dijadikan *health check* fungsional awal konektivitas API WHMCS.

### 4.2 Audit Log Reading
- Parameter pagination dan filtering berdasarkan rentang tanggal atau user staf.
- Membantu agen AI menganalisis kejadian keamanan atau perubahan data penting.
