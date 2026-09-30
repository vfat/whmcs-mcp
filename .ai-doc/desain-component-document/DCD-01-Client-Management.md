# DCD-01: Client Management Component

> **Target Codebase**: `whmcs-mcp` (Go 1.22+)  
> **Workspace**: `/home/ubuntu/workspace/plan/whmcs-mcp`  
> **Target Package**: `internal/mcp/tools/clients.go`, `internal/model/client.go`  
> **Tanggal**: 2026-09-30  

---

## 1. Overview

Komponen **Client Management** mengelola seluruh siklus hidup profil pelanggan (*client entity lifecycle*) pada WHMCS. Komponen ini mengekspos fungsi pencarian profil pengguna, pendaftaran akun baru, pembaruan data kontak & penagihan, penutupan akun, serta pengambilan relasi kepemilikan produk hosting, domain, dan faktur.

Target Pengguna:
- **LLM Agent / MCP Client**: Melakukan query data pelanggan dan pendaftaran akun otomatis via protokol MCP.
- **WHMCS API Gateway (`includes/api.php`)**: Sistem tujuan eksekusi aksi database `tblclients`.

---

## 2. Object Identification (Go Architecture)

### 2.1 Boundary
- **MCP Tool Handlers (`internal/mcp/tools/clients.go`)**:
  - `whmcs_get_clients`
  - `whmcs_get_client_details`
  - `whmcs_add_client`
  - `whmcs_update_client`
  - `whmcs_close_client` (*Destructive*)
  - `whmcs_get_client_products`
  - `whmcs_get_client_domains`
  - `whmcs_get_client_invoices`
- **Transport**: JSON-RPC 2.0 via `os.Stdin` / `os.Stdout`.

### 2.2 Control
- **Input Validation**: Tag validasi struct Go & sanitasi parameter.
- **WHMCS Client Dispatcher (`internal/whmcs/client.go`)**:
  - `GetClients(ctx context.Context, req model.GetClientsRequest) (*model.GetClientsResponse, error)`
  - `GetClientDetails(ctx context.Context, req model.GetClientDetailsRequest) (*model.GetClientDetailsResponse, error)`
  - `AddClient(ctx context.Context, req model.AddClientRequest) (*model.AddClientResponse, error)`
  - `UpdateClient(ctx context.Context, req model.UpdateClientRequest) (*model.UpdateClientResponse, error)`
  - `CloseClient(ctx context.Context, req model.CloseClientRequest) (*model.CloseClientResponse, error)`
  - `GetClientProducts(ctx context.Context, req model.GetClientProductsRequest) (*model.GetClientProductsResponse, error)`
  - `GetClientDomains(ctx context.Context, req model.GetClientDomainsRequest) (*model.GetClientDomainsResponse, error)`
  - `GetInvoices(ctx context.Context, req model.GetInvoicesRequest) (*model.GetInvoicesResponse, error)`
- **Parameter Serializer (`internal/whmcs/serializer.go`)**: Mengonversi struct request ke `url.Values` untuk body POST.

### 2.3 Entity
- **Go Structs (`internal/model/client.go`)**:
  - `Client`, `ClientDetails`, `ClientProduct`, `ClientDomain`
  - Request & Response DTOs per action.
- **Remote WHMCS Database Tables**: `tblclients`, `tblhosting`, `tbldomains`, `tblinvoices`.

---

## 3. Use Case List

| No | Use Case Name | Action WHMCS | Sifat | Go Function |
|:---:|:---|:---:|:---:|:---|
| 1 | Cari & Baca Data Klien | `GetClients`, `GetClientsDetails` | Safe | `GetClients`, `GetClientDetails` |
| 2 | Daftarkan & Perbarui Akun Klien | `AddClient`, `UpdateClient` | Safe | `AddClient`, `UpdateClient` |
| 3 | Tutup Akun Klien | `CloseClient` | Destructive | `CloseClient` |
| 4 | Ambil Layanan, Domain, & Faktur | `GetClientsProducts`, `GetClientsDomains`, `GetInvoices` | Safe | `GetClientProducts`, `GetClientDomains`, `GetInvoices` |

---

## 4. Technical Flow & Business Rules

### 4.1 Normal Flow (Query Profil Klien)
1. Host MCP mengirimkan call `whmcs_get_client_details` dengan argument `{"clientid": 101}`.
2. Handler unmarshal argument ke `model.GetClientDetailsRequest`.
3. Handler memanggil `client.GetClientDetails(ctx, req)`.
4. Serializer menambahkan `action=GetClientsDetails`, `responsetype=json`, kredensial, dan query `clientid=101`.
5. Client mengeksekusi HTTP POST ke `includes/api.php`.
6. Bila `result == "success"`, payload di-unmarshal ke `model.GetClientDetailsResponse` dan dikembalikan ke LLM Host via MCP tool result.

### 4.2 Error Handling & Validation
- **Client Not Found**: API WHMCS merespon `result=error&message=Client Not Found`. Go client mendeteksi `result == "error"` dan mengembalikan `ErrClientNotFound` terstruktur.
- **Destructive Action Protection**: `whmcs_close_client` memvalidasi keberadaan `clientid` dan menambahkan catatan audit log pada sistem.
