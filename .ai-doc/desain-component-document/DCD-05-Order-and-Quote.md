# DCD-05: Order and Quote Component

> **Target Codebase**: `whmcs-mcp` (Go 1.22+)  
> **Workspace**: `/home/ubuntu/workspace/plan/whmcs-mcp`  
> **Target Package**: `internal/mcp/tools/orders.go`, `internal/model/order.go`  
> **Tanggal**: 2026-09-30  

---

## 1. Overview

Komponen **Order and Quote** mengelola alur penerimaan pesanan baru pelanggan (paket hosting, add-on, domain), persetujuan pesanan (*accept order*), pembatalan, mitigasi fraud, serta pembuatan dan penerimaan surat penawaran biaya (*quotes*).

---

## 2. Object Identification (Go Architecture)

### 2.1 Boundary
- **MCP Tool Handlers (`internal/mcp/tools/orders.go`)**:
  - `whmcs_get_orders`, `whmcs_add_order`, `whmcs_accept_order`
  - `whmcs_cancel_order` (*Destructive*), `whmcs_fraud_order` (*Destructive*), `whmcs_delete_order` (*Destructive*)
  - `whmcs_get_quotes`, `whmcs_accept_quote`

### 2.2 Control
- **WHMCS Client Dispatcher (`internal/whmcs/client.go`)**:
  - `GetOrders`, `AddOrder`, `AcceptOrder`, `CancelOrder`, `FraudOrder`, `DeleteOrder`
  - `GetQuotes`, `AcceptQuote`

### 2.3 Entity
- **Go Structs (`internal/model/order.go`)**:
  - `Order`, `LineItem`, `Quote`
  - Request & Response DTOs per action.
- **Remote WHMCS Tables**: `tblorders`, `tblquotes`, `tblquoteitems`.

---

## 3. Use Case List

| No | Use Case Name | Action WHMCS | Sifat | Go Function |
|:---:|:---|:---:|:---:|:---|
| 1 | Buat & Setujui Pesanan Layanan | `GetOrders`, `AddOrder`, `AcceptOrder` | Safe | `GetOrders`, `AddOrder`, `AcceptOrder` |
| 2 | Tandai Fraud & Hapus Pesanan | `CancelOrder`, `FraudOrder`, `DeleteOrder` | Destructive | `CancelOrder`, `FraudOrder`, `DeleteOrder` |
| 3 | Kelola Penawaran Biaya (Quotes) | `GetQuotes`, `AcceptQuote` | Safe | `GetQuotes`, `AcceptQuote` |

---

## 4. Technical Flow & Business Rules

### 4.1 Pemesanan Baru (AddOrder)
- Menerima parameter `clientid`, `pid` (product ID), `domain`, `billingcycle`, `paymentmethod`.
- Mengembalikan `orderid` dan `invoiceid` terkait.

### 4.2 Persetujuan Pesanan (AcceptOrder)
- Memverifikasi status order `Pending`.
- Memicu otomatisasi modul hosting jika opsi modul auto-setup aktif pada konfigurasi produk.
