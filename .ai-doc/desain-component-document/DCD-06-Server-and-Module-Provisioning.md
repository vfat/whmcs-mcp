# DCD-06: Server and Module Provisioning Component

> **Target Codebase**: `whmcs-mcp` (Go 1.22+)  
> **Workspace**: `/home/ubuntu/workspace/plan/whmcs-mcp`  
> **Target Package**: `internal/mcp/tools/provisioning.go`, `internal/model/provisioning.go`  
> **Tanggal**: 2026-09-30  

---

## 1. Overview

Komponen **Server and Module Provisioning** menjembatani instruksi pengendalian akun hosting ke server remote (kontrol panel seperti cPanel/WHM, Plesk, DirectAdmin, Proxmox) melalui modul WHMCS, serta menyediakan inventaris pemantauan server fisik/virtual dan status kesehatan cron sistem.

---

## 2. Object Identification (Go Architecture)

### 2.1 Boundary
- **MCP Tool Handlers (`internal/mcp/tools/provisioning.go`)**:
  - `whmcs_module_create`, `whmcs_module_suspend` (*Destructive*), `whmcs_module_unsuspend`
  - `whmcs_module_terminate` (*Destructive*), `whmcs_module_change_package`, `whmcs_module_change_password`
  - `whmcs_get_servers`, `whmcs_get_health_status`
- **MCP Resource Handlers (`internal/mcp/resources/servers.go`)**:
  - `whmcs://servers`

### 2.2 Control
- **WHMCS Client Dispatcher (`internal/whmcs/client.go`)**:
  - `ModuleCreate`, `ModuleSuspend`, `ModuleUnsuspend`, `ModuleTerminate`
  - `ModuleChangePackage`, `ModuleChangePassword`, `ModuleCustom`
  - `GetServers`, `GetHealthStatus`

### 2.3 Entity
- **Go Structs (`internal/model/provisioning.go`)**:
  - `Server`, `HealthStatus`, `ModuleResult`
  - Request & Response DTOs.
- **Remote WHMCS Tables**: `tblservers`, `tblhosting`.

---

## 3. Use Case List

| No | Use Case Name | Action WHMCS | Sifat | Go Function |
|:---:|:---|:---:|:---:|:---|
| 1 | Provisioning Akun Hosting | `ModuleCreate`, `ModuleSuspend`, `ModuleUnsuspend`, `ModuleTerminate` | Campuran | `ModuleCreate`, `ModuleSuspend`, `ModuleUnsuspend`, `ModuleTerminate` |
| 2 | Ubah Paket Kuota & Password | `ModuleChangePackage`, `ModuleChangePassword` | Safe | `ModuleChangePackage`, `ModuleChangePassword` |
| 3 | Pantau Status Server & Kesehatan | `GetServers`, `GetHealthStatus` | Safe | `GetServers`, `GetHealthStatus` |

---

## 4. Technical Flow & Business Rules

### 4.1 Eksekusi Modul Hosting (ModuleCreate / ModuleSuspend / ModuleTerminate)
- Menerima `accountid` (service ID pada `tblhosting`).
- Untuk `ModuleSuspend`, parameter `suspendreason` diwajibkan untuk audit transparansi.
- Untuk `ModuleTerminate`, operasi bersifat permanen dan menghapus akun beserta seluruh data file/database pada server remote target.

### 4.2 Health Status & Server Inventory
- `whmcs_get_health_status` memeriksa indikator cron runtime, pembaruan versi WHMCS, status lisensi, dan direktori izin penyimpanan.
