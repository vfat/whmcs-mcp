# DCD-04: Domain Lifecycle Component

> **Target Codebase**: `whmcs-mcp` (Go 1.22+)  
> **Workspace**: `/home/ubuntu/workspace/plan/whmcs-mcp`  
> **Target Package**: `internal/mcp/tools/domains.go`, `internal/model/domain.go`  
> **Tanggal**: 2026-09-30  

---

## 1. Overview

Komponen **Domain Lifecycle** mengelola siklus hidup pendaftaran dan pemeliharaan nama domain pada WHMCS. Komponen ini berinteraksi dengan API registrar eksternal melalui modul registrar WHMCS untuk registrasi baru, perpanjangan masa aktif, transfer masuk, update nameserver, manajemen proteksi privasi WHOIS, serta status Registrar Lock dan permintaan kode otorisasi EPP.

---

## 2. Object Identification (Go Architecture)

### 2.1 Boundary
- **MCP Tool Handlers (`internal/mcp/tools/domains.go`)**:
  - `whmcs_get_domains`, `whmcs_domain_whois`
  - `whmcs_domain_register`, `whmcs_domain_renew`, `whmcs_domain_transfer`
  - `whmcs_domain_release` (*Destructive*)
  - `whmcs_domain_request_epp`, `whmcs_domain_toggle_id_protect`
  - `whmcs_domain_get_nameservers`, `whmcs_domain_update_nameservers`
  - `whmcs_domain_get_lock_status`, `whmcs_domain_update_lock_status`
- **MCP Resource Handlers (`internal/mcp/resources/domains.go`)**:
  - `whmcs://tld-pricing`

### 2.2 Control
- **WHMCS Client Dispatcher (`internal/whmcs/client.go`)**:
  - `GetDomains`, `DomainWhois`, `DomainRegister`, `DomainRenew`, `DomainTransfer`
  - `DomainRelease`, `DomainRequestEPP`, `DomainToggleIdProtect`
  - `DomainGetNameservers`, `DomainUpdateNameservers`
  - `DomainGetLockingStatus`, `DomainUpdateLockingStatus`, `GetTLDPricing`

### 2.3 Entity
- **Go Structs (`internal/model/domain.go`)**:
  - `Domain`, `Nameservers`, `TLDPricing`, `WhoisResult`
  - Request & Response DTOs.
- **Remote WHMCS Tables**: `tbldomains`, `tblpricing`.

---

## 3. Use Case List

| No | Use Case Name | Action WHMCS | Sifat | Go Function |
|:---:|:---|:---:|:---:|:---|
| 1 | Cek WHOIS & Ketersediaan Domain | `DomainWhois`, `GetClientsDomains` | Safe | `DomainWhois`, `GetDomains` |
| 2 | Registrasi, Perpanjang, & Transfer | `DomainRegister`, `DomainRenew`, `DomainTransfer`, `DomainRelease` | Campuran | `DomainRegister`, `DomainRenew`, `DomainTransfer`, `DomainRelease` |
| 3 | Konfigurasi Nameserver, Lock, & EPP | `DomainGetNameservers`, `DomainUpdateNameservers`, `DomainGetLockingStatus`, `DomainUpdateLockingStatus`, `DomainToggleIdProtect`, `DomainRequestEPP` | Safe | `DomainGetNameservers`, `DomainUpdateNameservers`, `DomainGetLockingStatus`, `DomainUpdateLockingStatus`, `DomainToggleIdProtect`, `DomainRequestEPP` |

---

## 4. Technical Flow & Business Rules

### 4.1 Registrasi & Perpanjangan Domain
- Pemanggilan `whmcs_domain_register` / `whmcs_domain_renew` meneruskan perintah via API registrar yang terhubung di WHMCS.
- Penanganan timeout registrar eksternal: Go context timeout diterapkan (default 60s untuk panggilan modul registrar).

### 4.2 Nameserver Updates
- Input menerima daftar `ns1` s/d `ns5`.
- Serializer Go memetakan argumen ke struktur flat `ns1=...&ns2=...`.
