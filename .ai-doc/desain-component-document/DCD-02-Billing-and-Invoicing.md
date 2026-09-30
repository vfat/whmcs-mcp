# DCD-02: Billing and Invoicing Component

> **Target Codebase**: `whmcs-mcp` (Go 1.22+)  
> **Workspace**: `/home/ubuntu/workspace/plan/whmcs-mcp`  
> **Target Package**: `internal/mcp/tools/billing.go`, `internal/model/billing.go`  
> **Tanggal**: 2026-09-30  

---

## 1. Overview

Komponen **Billing and Invoicing** mengelola siklus penagihan keuangan, faktur tagihan (*invoices*), mutasi transaksi kas, pencatatan pembayaran manual, penagihan kartu kredit otomatis (*merchant capture*), dan aplikasi saldo kredit pada WHMCS.

---

## 2. Object Identification (Go Architecture)

### 2.1 Boundary
- **MCP Tool Handlers (`internal/mcp/tools/billing.go`)**:
  - `whmcs_get_invoices`, `whmcs_get_invoice`, `whmcs_create_invoice`, `whmcs_update_invoice`
  - `whmcs_add_invoice_payment`, `whmcs_capture_payment`, `whmcs_apply_credit`
  - `whmcs_get_transactions`, `whmcs_add_transaction`
  - `whmcs_create_billable_item`, `whmcs_update_client_product`
- **MCP Resource Handlers (`internal/mcp/resources/billing.go`)**:
  - `whmcs://payment-methods`, `whmcs://currencies`

### 2.2 Control
- **WHMCS Client Dispatcher (`internal/whmcs/client.go`)**:
  - `GetInvoices`, `GetInvoice`, `CreateInvoice`, `UpdateInvoice`
  - `AddInvoicePayment`, `CapturePayment`, `ApplyCredit`
  - `GetTransactions`, `AddTransaction`
  - `CreateBillableItem`, `UpdateClientProduct`
  - `GetPaymentMethods`, `GetCurrencies`

### 2.3 Entity
- **Go Structs (`internal/model/billing.go`)**:
  - `Invoice`, `InvoiceItem`, `Transaction`, `PaymentMethod`, `Currency`, `BillableItem`
  - Request & Response DTOs per action.
- **Remote WHMCS Tables**: `tblinvoices`, `tblinvoiceitems`, `tblaccounts`, `tblbillableitems`.

---

## 3. Use Case List

| No | Use Case Name | Action WHMCS | Sifat | Go Function |
|:---:|:---|:---:|:---:|:---|
| 1 | Kelola Faktur & Item Tertagih | `GetInvoices`, `GetInvoice`, `CreateInvoice`, `UpdateInvoice`, `CreateBillableItem` | Safe | `GetInvoices`, `GetInvoice`, `CreateInvoice`, `UpdateInvoice`, `CreateBillableItem` |
| 2 | Catat Pembayaran & Tangkap Transaksi | `AddInvoicePayment`, `CapturePayment`, `GetTransactions`, `AddTransaction` | Safe | `AddInvoicePayment`, `CapturePayment`, `GetTransactions`, `AddTransaction` |
| 3 | Kelola Mutasi & Saldo Kredit Klien | `ApplyCredit` | Safe | `ApplyCredit` |
| 4 | Ubah Siklus Tagihan Produk Klien | `UpdateClientProduct` | Safe | `UpdateClientProduct` |

---

## 4. Technical Flow & Business Rules

### 4.1 Pembuatan Faktur Baru (CreateInvoice)
- Mengonversi rincian item (`itemdescription`, `itemamount`, `itemtaxed`) ke parameter array multidimensi.
- Validasi nilai moneter tidak boleh bernilai negatif kecuali dinyatakan secara eksplisit sebagai item diskon.
- Respon mengembalikan `invoiceid` yang baru dibuat.

### 4.2 Auto-Capture Pembayaran Kartu Kredit
- `whmcs_capture_payment` memicu gateway kartu kredit aktif untuk menagih kartu yang tersimpan pada profil klien.
- Menangani respon asynchronous gateway dan update status faktur menjadi `Paid`.
