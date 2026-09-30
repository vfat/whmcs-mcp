# Dokumentasi Fitur & Prioritas Porting: WHMCS MCP Server (Golang)

> **Target Codebase**: `whmcs-mcp` (Go 1.22+)  
> **Workspace**: `/home/ubuntu/workspace/plan/whmcs-mcp`  
> **Rujukan Baseline**: `plan/whmcs-mcp-tool/.ai-doc/Dokumentasi-Fitur.md`  
> **Metode**: Feature Inventory & TDD Phased Rollout Plan (`ai-documentor`)  
> **Tanggal**: 2026-09-30  

Dokumen ini mendokumentasikan inventaris lengkap seluruh kapabilitas yang diporting dari implementasi TypeScript ke Golang: **62 MCP Tools**, **8 MCP Prompts**, dan **11 MCP Resources** skema `whmcs://*`. Setiap entri dipetakan ke aksi REST API WHMCS, target paket Go (`internal/mcp/tools/`), target struct DTO (`internal/model/`), serta prioritas siklus TDD.

---

## 1. Strategi Pentahapan Implementasi TDD (Phased Batches)

Untuk menjaga disiplin siklus **RED → GREEN → REFACTOR**, implementasi dibagi menjadi 5 fase/batch:

| Batch | Fokus Domain | Jumlah Tool | Jumlah Resource | Prompt | Prioritas |
|:---:|:---|:---:|:---:|:---:|:---:|
| **Batch 0** | **Fondasi Runtime**: Config, Param Serializer, HTTP Client Core | — | — | — | **P0 (Kritis)** |
| **Batch 1** | **Klien & Administrasi Sistem**: Client Management, System Metrics | 14 Tools | 5 Resources | 2 Prompts | **P0 (Tinggi)** |
| **Batch 2** | **Keuangan & Layanan Bantuan**: Invoices, Payments, Support Tickets | 20 Tools | 4 Resources | 3 Prompts | **P1 (Tinggi)** |
| **Batch 3** | **Siklus Hidup Domain & Pesanan**: Domains, Orders, Quotes | 20 Tools | 2 Resources | 2 Prompts | **P2 (Menengah)** |
| **Batch 4** | **Provisioning Server & Promosi**: Modules, Servers, Affiliates, Promos | 8 Tools | — | 1 Prompt | **P3 (Lanjutan)** |
| **Total** | **Seluruh Sistem** | **62 Tools** | **11 Resources** | **8 Prompts** | **100% Paritas** |

---

## 2. Inventaris Fitur Berdasarkan Domain

### A. Manajemen Klien (*Client Management*) — Batch 1 (P0)

Paket Target: `internal/mcp/tools/clients.go` | DTO: `internal/model/client.go`

| Nama Tool MCP | Tipe | Action WHMCS | Sifat | Deskripsi Fungsional & Parameter Kunci |
|---|---|---|---|---|
| `whmcs_get_clients` | Tool | `GetClients` | Safe | Daftar klien (filter: `search`, `status`, `sorting`, pagination `limitstart`, `limitnum`). |
| `whmcs_get_client_details` | Tool | `GetClientsDetails` | Safe | Rincian profil klien lengkap via `clientid` atau `email`, opsi kalkulasi `stats`. |
| `whmcs_add_client` | Tool | `AddClient` | Safe | Registrasi klien baru (nama, email, alamat, telepon, password, customfields). |
| `whmcs_update_client` | Tool | `UpdateClient` | Safe | Pembaruan profil klien via `clientid` (email, alamat, status, credit balance). |
| `whmcs_close_client` | Tool | `CloseClient` | Destructive | Menutup status akun klien menjadi `Closed` via `clientid`. |
| `whmcs_get_client_products` | Tool | `GetClientsProducts` | Safe | Mengambil daftar produk/layanan hosting aktif milik `clientid`. |
| `whmcs_get_client_domains` | Tool | `GetClientsDomains` | Safe | Mengambil daftar nama domain yang terdaftar atas nama `clientid`. |
| `whmcs_get_client_invoices` | Tool | `GetInvoices` | Safe | Mengambil riwayat faktur tagihan spesifik milik `clientid`. |

---

### B. Katalog Produk & Layanan (*Product Management*) — Batch 2 (P1)

Paket Target: `internal/mcp/tools/products.go` | DTO: `internal/model/product.go`

| Nama Fitur / URI | Tipe | Action WHMCS | Sifat | Deskripsi Fungsional |
|---|---|---|---|---|
| `whmcs_get_products` | Tool | `GetProducts` | Safe | Katalog paket hosting/produk (filter: `pid`, `gid`, `module`). |
| `whmcs://products` | Resource | `GetProducts` | Read-only | Resource URI instan katalog produk aktif tanpa filter parameter. |

---

### C. Penagihan & Keuangan (*Billing & Invoices*) — Batch 2 (P1)

Paket Target: `internal/mcp/tools/billing.go` | DTO: `internal/model/billing.go`

| Nama Tool / Resource | Tipe | Action WHMCS | Sifat | Deskripsi Fungsional |
|---|---|---|---|---|
| `whmcs_get_invoices` | Tool | `GetInvoices` | Safe | Daftar invoice global dengan filter status (`Paid`, `Unpaid`, `Overdue`, `Cancelled`). |
| `whmcs_get_invoice` | Tool | `GetInvoice` | Safe | Detail rincian invoice (line items, subtotal, tax, credit, payments) via `invoiceid`. |
| `whmcs_create_invoice` | Tool | `CreateInvoice` | Safe | Pembuatan faktur manual baru dengan rincian item, taxrate, dan duedate. |
| `whmcs_update_invoice` | Tool | `UpdateInvoice` | Safe | Memperbarui atribut invoice (status, duedate, paymentmethod, notes). |
| `whmcs_add_invoice_payment` | Tool | `AddInvoicePayment` | Safe | Mencatat pembayaran manual (transid, gateway, date, amount, fees). |
| `whmcs_capture_payment` | Tool | `CapturePayment` | Safe | Menjalankan auto-charge payment gateway kartu kredit untuk faktur tertagih. |
| `whmcs_apply_credit` | Tool | `ApplyCredit` | Safe | Menerapkan saldo kredit akun klien untuk pelunasan invoice tertentu. |
| `whmcs_get_transactions` | Tool | `GetTransactions` | Safe | Riwayat transaksi buku kas / ledger (filter: `clientid`, `invoiceid`, `transid`). |
| `whmcs_add_transaction` | Tool | `AddTransaction` | Safe | Menambahkan pencatatan mutasi kas masuk/keluar ke sistem buku kas. |
| `whmcs_create_billable_item` | Tool | `CreateBillableItem` | Safe | Menjadwalkan item tertagih ke faktur siklus penagihan berikutnya. |
| `whmcs_update_client_product` | Tool | `UpdateClientProduct` | Safe | Mengubah rincian langganan hosting (billingcycle, nextduedate, recurringamount). |
| `whmcs://payment-methods` | Resource | `GetPaymentMethods` | Read-only | Resource URI daftar gateway pembayaran yang aktif. |
| `whmcs://currencies` | Resource | `GetCurrencies` | Read-only | Resource URI konfigurasi mata uang sistem dan nilai tukar aktif. |

---

### D. Tiket Bantuan Pelanggan (*Support Tickets*) — Batch 2 (P1)

Paket Target: `internal/mcp/tools/tickets.go` | DTO: `internal/model/ticket.go`

| Nama Tool / Resource | Tipe | Action WHMCS | Sifat | Deskripsi Fungsional |
|---|---|---|---|---|
| `whmcs_get_tickets` | Tool | `GetTickets` | Safe | Antrean tiket bantuan (filter: `status`, `deptid`, `clientid`, `priority`). |
| `whmcs_get_ticket` | Tool | `GetTicket` | Safe | Thread lengkap riwayat percakapan tiket beserta internal staff notes via `ticketid`. |
| `whmcs_open_ticket` | Tool | `OpenTicket` | Safe | Membuka tiket baru untuk klien/pengunjung (deptid, subject, message, priority). |
| `whmcs_reply_ticket` | Tool | `AddTicketReply` | Safe | Menambahkan balasan staf atau respon klien ke thread tiket aktif. |
| `whmcs_update_ticket` | Tool | `UpdateTicket` | Safe | Mengubah status tiket, eskalasi departemen, penugasan admin (`flag`), prioritas. |
| `whmcs_get_ticket_predefined_cats` | Tool | `GetTicketPredefinedCats` | Safe | Daftar kategori templat balasan cepat (*predefined replies*). |
| `whmcs_get_ticket_predefined_replies`| Tool | `GetTicketPredefinedReplies`| Safe | Mengambil daftar templat jawaban cepat berdasarkan `catid`. |
| `whmcs_get_support_departments` | Tool | `GetSupportDepartments` | Safe | Mengambil daftar departemen bantuan dan metrik tiket aktif per departemen. |
| `whmcs_get_support_statuses` | Tool | `GetSupportStatuses` | Safe | Daftar status tiket yang terkonfigurasi pada sistem. |
| `whmcs://support/departments` | Resource | `GetSupportDepartments` | Read-only | Resource URI konfigurasi departemen bantuan. |
| `whmcs://support/statuses` | Resource | `GetSupportStatuses` | Read-only | Resource URI konfigurasi status tiket. |

---

### E. Manajemen Domain (*Domain Management*) — Batch 3 (P2)

Paket Target: `internal/mcp/tools/domains.go` | DTO: `internal/model/domain.go`

| Nama Tool / Resource | Tipe | Action WHMCS | Sifat | Deskripsi Fungsional |
|---|---|---|---|---|
| `whmcs_get_domains` | Tool | `GetClientsDomains` | Safe | Daftar domain terdaftar seluruh sistem (filter: `domain`, `status`, `clientid`). |
| `whmcs_domain_whois` | Tool | `DomainWhois` | Safe | Pengecekan ketersediaan domain dan query WHOIS publik. |
| `whmcs_domain_register` | Tool | `DomainRegister` | Safe | Memerintahkan modul registrar untuk registrasi domain baru via `domainid`. |
| `whmcs_domain_renew` | Tool | `DomainRenew` | Safe | Memperpanjang masa aktif domain pada registrar via `domainid`. |
| `whmcs_domain_transfer` | Tool | `DomainTransfer` | Safe | Menjalankan proses transfer masuk domain menggunakan EPP/auth code. |
| `whmcs_domain_release` | Tool | `DomainRelease` | Destructive | Melepaskan domain ke tag registrar baru (khusus domain UK / IPS Tag). |
| `whmcs_domain_request_epp` | Tool | `DomainRequestEPP` | Safe | Mengirimkan kode otorisasi transfer EPP ke email pemilik terdaftar. |
| `whmcs_domain_toggle_id_protect` | Tool | `DomainToggleIdProtect` | Safe | Mengaktifkan/menonaktifkan proteksi identitas WHOIS (ID Protection). |
| `whmcs_domain_get_nameservers` | Tool | `DomainGetNameservers` | Safe | Membaca konfigurasi nameserver aktif (ns1-ns5) dari registrar. |
| `whmcs_domain_update_nameservers`| Tool | `DomainUpdateNameservers`| Safe | Memperbarui alamat nameserver domain pada registrar. |
| `whmcs_domain_get_lock_status` | Tool | `DomainGetLockingStatus` | Safe | Memeriksa status Registrar Lock (terkunci / terbuka). |
| `whmcs_domain_update_lock_status`| Tool | `DomainUpdateLockingStatus`| Safe | Mengubah status penguncian transfer domain (`locked` / `unlocked`). |
| `whmcs://tld-pricing` | Resource | `GetTLDPRicing` | Read-only | Resource URI daftar harga registrasi, transfer, dan renewal seluruh TLD. |

---

### F. Manajemen Pesanan & Penawaran (*Orders & Quotes*) — Batch 3 (P2)

Paket Target: `internal/mcp/tools/orders.go` | DTO: `internal/model/order.go`

| Nama Tool MCP | Tipe | Action WHMCS | Sifat | Deskripsi Fungsional |
|---|---|---|---|---|
| `whmcs_get_orders` | Tool | `GetOrders` | Safe | Daftar pesanan pelanggan (filter status: `Pending`, `Active`, `Fraud`, `Cancelled`). |
| `whmcs_add_order` | Tool | `AddOrder` | Safe | Membuat order pesanan baru produk/domain lengkap dengan opsi gateway pembayaran. |
| `whmcs_accept_order` | Tool | `AcceptOrder` | Safe | Menyetujui order pending dan memicu pembuatan invoice/layanan otomatis. |
| `whmcs_cancel_order` | Tool | `CancelOrder` | Destructive | Membatalkan order pelanggan yang berstatus tertunda. |
| `whmcs_fraud_order` | Tool | `FraudOrder` | Destructive | Menandai order sebagai fraud dan membatalkan transaksi terkait. |
| `whmcs_delete_order` | Tool | `DeleteOrder` | Destructive | Menghapus permanen rekaman order dari basis data WHMCS. |
| `whmcs_get_quotes` | Tool | `GetQuotes` | Safe | Mengambil daftar surat penawaran (*quotes*) berdasarkan subjek atau status. |
| `whmcs_accept_quote` | Tool | `AcceptQuote` | Safe | Menyetujui surat penawaran dan otomatis mengonversinya menjadi faktur/tagihan. |

---

### G. Server & Modul Provisioning (*Modules & Servers*) — Batch 4 (P3)

Paket Target: `internal/mcp/tools/provisioning.go` | DTO: `internal/model/provisioning.go`

| Nama Tool / Resource | Tipe | Action WHMCS | Sifat | Deskripsi Fungsional |
|---|---|---|---|---|
| `whmcs_module_create` | Tool | `ModuleCreate` | Safe | Mengeksekusi pembuatan akun hosting otomatis pada server kontrol panel target. |
| `whmcs_module_suspend` | Tool | `ModuleSuspend` | Destructive | Menangguhkan akun hosting pada server kontrol panel beserta alasan penangguhan. |
| `whmcs_module_unsuspend` | Tool | `ModuleUnsuspend` | Safe | Mengaktifkan kembali akun hosting yang ditangguhkan. |
| `whmcs_module_terminate` | Tool | `ModuleTerminate` | Destructive | Menghapus permanen akun hosting dari server target (destructive). |
| `whmcs_module_change_package` | Tool | `ModuleChangePackage` | Safe | Mengubah paket/kuota hosting pada server kontrol panel (upgrade/downgrade). |
| `whmcs_module_change_password`| Tool | `ModuleChangePassword`| Safe | Mengubah password akses akun hosting pada server kontrol panel. |
| `whmcs_get_servers` | Tool | `GetServers` | Safe | Mengambil daftar server hosting (IP, status aktif, kuota akun, modul). |
| `whmcs_get_health_status` | Tool | `GetHealthStatus` | Safe | Memeriksa indikator kesehatan sistem WHMCS dan cron runner. |
| `whmcs://servers` | Resource | `GetServers` | Read-only | Resource URI konfigurasi seluruh server hosting. |

---

### H. Sistem & Administrasi (*System & Admin*) — Batch 1 (P0)

Paket Target: `internal/mcp/tools/system.go` | DTO: `internal/model/system.go`

| Nama Tool / Resource | Tipe | Action WHMCS | Sifat | Deskripsi Fungsional |
|---|---|---|---|---|
| `whmcs_get_stats` | Tool | `GetStats` | Safe | Ringkasan metrik statistik operasional (income, open tickets, orders, cancel). |
| `whmcs_get_activity_log` | Tool | `GetActivityLog` | Safe | Membaca log audit aktivitas administratif untuk kebutuhan security/investigasi. |
| `whmcs_get_admin_users` | Tool | `GetAdminUsers` | Safe | Mengambil daftar user staf administrator yang memiliki akses sistem. |
| `whmcs_get_todo_items` | Tool | `GetToDoItems` | Safe | Mengambil daftar tugas administratif internal staf (*To-Do items*). |
| `whmcs_get_todo_item_statuses`| Tool | `GetToDoItemStatuses` | Safe | Mengambil daftar status tugas to-do yang tersedia. |
| `whmcs_update_todo_item` | Tool | `UpdateToDoItem` | Safe | Memperbarui status, tanggal, atau penugasan staf pada tugas to-do. |
| `whmcs_get_currencies` | Tool | `GetCurrencies` | Safe | Mengambil daftar mata uang penagihan dan kurs nilai tukar. |
| `whmcs_get_payment_methods` | Tool | `GetPaymentMethods` | Safe | Mengambil daftar metode pembayaran yang diaktifkan. |
| `whmcs_send_email` | Tool | `SendEmail` | Safe | Mengirimkan email kustom atau memicu templat email sistem kepada klien. |
| `whmcs://stats` | Resource | `GetStats` | Read-only | Resource URI ringkasan statistik operasional WHMCS secara instan. |
| `whmcs://admin-users` | Resource | `GetAdminUsers` | Read-only | Resource URI daftar user staf administrator. |
| `whmcs://admin/todo` | Resource | `GetToDoItems` | Read-only | Resource URI antrean tugas to-do staf. |

---

### I. Pemasaran & Pertumbuhan (*Affiliates & Promotions*) — Batch 4 (P3)

Paket Target: `internal/mcp/tools/marketing.go` | DTO: `internal/model/marketing.go`

| Nama Tool / Resource | Tipe | Action WHMCS | Sifat | Deskripsi Fungsional |
|---|---|---|---|---|
| `whmcs_get_affiliates` | Tool | `GetAffiliates` | Safe | Daftar akun afiliasi, saldo komisi tertagih, dan jumlah klik referal. |
| `whmcs_affiliate_activate` | Tool | `AffiliateActivate` | Safe | Mengaktifkan akun afiliasi baru untuk klien tertentu. |
| `whmcs_get_promotions` | Tool | `GetPromotions` | Safe | Daftar kode kupon diskon/promosi aktif dan persentase potongan. |
| `whmcs://promotions` | Resource | `GetPromotions` | Read-only | Resource URI daftar promosi aktif. |

---

### J. Templat Prompt MCP (*Prompts Suite*) — Batch 1 s/d Batch 4

Paket Target: `internal/mcp/prompts/prompts.go`

| Nama Prompt | Batch | Parameter Input | Deskripsi Skenario Penggunaan |
|---|:---:|---|---|
| `client-onboarding` | Batch 1 | `client_name`, `email`, `package_name`, `domain` | Panduan langkah orientasi klien baru (buat akun, pesanan, invoice, kirim email). |
| `client-health-check` | Batch 1 | `client_id` | Audit akun klien (riwayat invoice macet, tiket bantuan, paket aktif). |
| `ticket-response` | Batch 2 | `ticket_id`, `context` (opsional) | Formulasi respon balasan tiket pelanggan berdasarkan thread percakapan sebelumnya. |
| `revenue-report` | Batch 2 | `period` (`day`, `week`, `month`, `year`) | Analisis pendapatan, tagihan outstanding, perbandingan tren, dan ringkasan keuangan. |
| `bulk-invoice-reminder`| Batch 2 | `days_overdue` (opsional) | Strategi pengingat invoice jatuh tempo massal dan draf pesan penagihan persuasif. |
| `domain-expiry-audit` | Batch 3 | `days_until_expiry` (opsional) | Audit domain menjelang kedaluwarsa beserta verifikasi autorenew dan registrar lock. |
| `fraud-investigation` | Batch 3 | `order_id` | Investigasi kecurigaan transaksi fraud berdasarkan IP address, alamat, dan pola order. |
| `new-product-setup` | Batch 4 | `product_type`, `target_market` (opsional) | Panduan konfigurasi paket hosting baru (harga, kuota disk/bandwidth, cPanel module). |
