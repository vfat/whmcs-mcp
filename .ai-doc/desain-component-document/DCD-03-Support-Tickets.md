# DCD-03: Support Tickets Component

> **Target Codebase**: `whmcs-mcp` (Go 1.22+)  
> **Workspace**: `/home/ubuntu/workspace/plan/whmcs-mcp`  
> **Target Package**: `internal/mcp/tools/tickets.go`, `internal/model/ticket.go`  
> **Tanggal**: 2026-09-30  

---

## 1. Overview

Komponen **Support Tickets** mengelola operasional *helpdesk* pelanggan pada WHMCS, mencakup pembuatan tiket baru, pemuatan thread balasan, pengiriman respon staf/klien, eskalasi departemen, penyesuaian status tiket, dan pemanfaatan templat balasan cepat (*predefined replies*).

---

## 2. Object Identification (Go Architecture)

### 2.1 Boundary
- **MCP Tool Handlers (`internal/mcp/tools/tickets.go`)**:
  - `whmcs_get_tickets`, `whmcs_get_ticket`, `whmcs_open_ticket`, `whmcs_reply_ticket`, `whmcs_update_ticket`
  - `whmcs_get_ticket_predefined_cats`, `whmcs_get_ticket_predefined_replies`
  - `whmcs_get_support_departments`, `whmcs_get_support_statuses`
- **MCP Resource Handlers (`internal/mcp/resources/tickets.go`)**:
  - `whmcs://support/departments`, `whmcs://support/statuses`

### 2.2 Control
- **WHMCS Client Dispatcher (`internal/whmcs/client.go`)**:
  - `GetTickets`, `GetTicket`, `OpenTicket`, `AddTicketReply`, `UpdateTicket`
  - `GetTicketPredefinedCats`, `GetTicketPredefinedReplies`
  - `GetSupportDepartments`, `GetSupportStatuses`

### 2.3 Entity
- **Go Structs (`internal/model/ticket.go`)**:
  - `Ticket`, `TicketReply`, `TicketNote`, `SupportDepartment`, `SupportStatus`, `PredefinedReply`
  - Request & Response DTOs.
- **Remote WHMCS Tables**: `tbltickets`, `tblticketreplies`, `tblticketnotes`, `tblticketdepartments`.

---

## 3. Use Case List

| No | Use Case Name | Action WHMCS | Sifat | Go Function |
|:---:|:---|:---:|:---:|:---|
| 1 | Buka & Balas Tiket Bantuan | `GetTickets`, `GetTicket`, `OpenTicket`, `AddTicketReply` | Safe | `GetTickets`, `GetTicket`, `OpenTicket`, `AddTicketReply` |
| 2 | Kelola Status & Eskalasi Departemen | `UpdateTicket`, `GetSupportDepartments`, `GetSupportStatuses` | Safe | `UpdateTicket`, `GetSupportDepartments`, `GetSupportStatuses` |
| 3 | Ambil Templat Balasan Cepat | `GetTicketPredefinedCats`, `GetTicketPredefinedReplies` | Safe | `GetTicketPredefinedCats`, `GetTicketPredefinedReplies` |

---

## 4. Technical Flow & Business Rules

### 4.1 Buka Tiket Baru (OpenTicket)
- Validasi input: `deptid`, `subject`, `message`, serta identitas pengirim (`clientid` atau `email` & `name`).
- Mengembalikan `ticketid` dan `tid` (tracking ticket mask number, misal: `#123456`).

### 4.2 Balas Tiket (AddTicketReply)
- Memastikan `ticketid` valid.
- Mendukung pengiriman atas nama staf (`adminusername`) atau klien.
- Otomatis memperbarui status tiket menjadi `Answered` bila dibalas oleh admin.
