# Katalog DTO & Pemetaan Endpoint REST API WHMCS

> **Target Codebase**: `whmcs-mcp` (Go 1.22+)  
> **Workspace**: `/home/ubuntu/workspace/plan/whmcs-mcp`  
> **Rujukan Spesifikasi Upstream**: `plan/whmcs-mcp-tool/.ai-doc/rest-api-doc/` (78 Berkas `API-SPEC-001` s/d `API-SPEC-078`)  
> **Target Paket Go**: `internal/model/*.go`  
> **Tanggal**: 2026-09-30  

Dokumen ini memetakan seluruh 78 aksi REST API WHMCS (`includes/api.php`) ke dalam rancangan Struct Go *strongly-typed* (Request DTO & Response DTO) yang akan diimplementasikan pada paket `internal/model/` dan dikonsumsi oleh `internal/whmcs/client.go`.

---

## 1. Konvensi Tipe Data Struct Go

1. **Tagging Struct**:
   - `json:"<field>,omitempty"`: Digunakan untuk serialisasi/deserialisasi payload JSON.
   - `url:"<field>"`: Digunakan oleh serializer form URL-encoded untuk parameter POST.
2. **Respon Dasar WHMCS (`CommonResponse`)**:
   Setiap response struct menyematkan (*embed*) struct status umum WHMCS:
   ```go
   type CommonResponse struct {
       Result  string `json:"result"`            // "success" atau "error"
       Message string `json:"message,omitempty"`  // Pesan error jika gagal
   }
   ```
3. **Paginasi & Sorting**:
   Struct request list mewarisi properti paginasi umum:
   ```go
   type PaginationParams struct {
       LimitStart int    `json:"limitstart,omitempty" url:"limitstart"`
       LimitNum   int    `json:"limitnum,omitempty" url:"limitnum"`
       Sorting    string `json:"sorting,omitempty" url:"sorting"`
   }
   ```

---

## 2. Tabel Pemetaan 78 Aksi WHMCS ke Struct Go

### Domain 1: Client Management (`internal/model/client.go`)

| ID Dokumen REST API | WHMCS Action | Struct Request Go | Struct Response Go | Keterangan & Parameter Utama |
|---|---|---|---|---|
| `API-SPEC-001` | `GetClients` | `GetClientsRequest` | `GetClientsResponse` | Pencarian klien (`search`, `status`, pagination). |
| `API-SPEC-002` | `GetClientsDetails` | `GetClientDetailsRequest` | `GetClientDetailsResponse` | Profil lengkap via `clientid` atau `email`. |
| `API-SPEC-003` | `AddClient` | `AddClientRequest` | `AddClientResponse` | Registrasi akun klien baru (firstname, lastname, email, address). |
| `API-SPEC-004` | `UpdateClient` | `UpdateClientRequest` | `UpdateClientResponse` | Modifikasi profil via `clientid`. |
| `API-SPEC-005` | `CloseClient` | `CloseClientRequest` | `CloseClientResponse` | Penutupan akun (`clientid`). |
| `API-SPEC-006` | `DeleteClient` | `DeleteClientRequest` | `DeleteClientResponse` | Penghapusan akun klien permanen (*destructive*). |
| `API-SPEC-007` | `GetClientsProducts` | `GetClientProductsRequest` | `GetClientProductsResponse` | Daftar layanan hosting aktif milik klien. |
| `API-SPEC-008` | `GetClientsDomains` | `GetClientDomainsRequest` | `GetClientDomainsResponse` | Daftar domain terdaftar milik klien. |

---

### Domain 2: Product Management (`internal/model/product.go`)

| ID Dokumen REST API | WHMCS Action | Struct Request Go | Struct Response Go | Keterangan & Parameter Utama |
|---|---|---|---|---|
| `API-SPEC-009` | `GetProducts` | `GetProductsRequest` | `GetProductsResponse` | Katalog produk hosting (`pid`, `gid`, `module`). |
| `API-SPEC-010` | `GetProductGroups` | `GetProductGroupsRequest` | `GetProductGroupsResponse` | Grup kategori paket hosting. |

---

### Domain 3: Billing & Invoicing (`internal/model/billing.go`)

| ID Dokumen REST API | WHMCS Action | Struct Request Go | Struct Response Go | Keterangan & Parameter Utama |
|---|---|---|---|---|
| `API-SPEC-011` | `GetInvoices` | `GetInvoicesRequest` | `GetInvoicesResponse` | Daftar faktur (filter `status`, `userid`). |
| `API-SPEC-012` | `GetInvoice` | `GetInvoiceRequest` | `GetInvoiceResponse` | Rincian line-item faktur via `invoiceid`. |
| `API-SPEC-013` | `CreateInvoice` | `CreateInvoiceRequest` | `CreateInvoiceResponse` | Buat faktur manual (`itemdescription`, `itemamount`, `itemtaxed`). |
| `API-SPEC-014` | `UpdateInvoice` | `UpdateInvoiceRequest` | `UpdateInvoiceResponse` | Update atribut invoice (`status`, `duedate`, `notes`). |
| `API-SPEC-015` | `AddInvoicePayment` | `AddInvoicePaymentRequest` | `AddInvoicePaymentResponse` | Catat pembayaran (`transid`, `gateway`, `amount`, `fees`). |
| `API-SPEC-016` | `CapturePayment` | `CapturePaymentRequest` | `CapturePaymentResponse` | Eksekusi auto-charge kartu kredit. |
| `API-SPEC-017` | `ApplyCredit` | `ApplyCreditRequest` | `ApplyCreditResponse` | Aplikasikan saldo kredit ke faktur. |
| `API-SPEC-018` | `GetTransactions` | `GetTransactionsRequest` | `GetTransactionsResponse` | Riwayat buku kas / ledger (`invoiceid`, `transid`). |
| `API-SPEC-019` | `AddTransaction` | `AddTransactionRequest` | `AddTransactionResponse` | Tambah transaksi kas masuk/keluar. |
| `API-SPEC-020` | `CreateBillableItem` | `CreateBillableItemRequest`| `CreateBillableItemResponse`| Jadwalkan item tertagih berulang/sekali bayar. |
| `API-SPEC-021` | `UpdateClientProduct` | `UpdateClientProductRequest`| `UpdateClientProductResponse`| Ubah siklus tagihan/harga paket klien. |

---

### Domain 4: Support Tickets (`internal/model/ticket.go`)

| ID Dokumen REST API | WHMCS Action | Struct Request Go | Struct Response Go | Keterangan & Parameter Utama |
|---|---|---|---|---|
| `API-SPEC-022` | `GetTickets` | `GetTicketsRequest` | `GetTicketsResponse` | Antrean tiket (filter `status`, `deptid`, `clientid`). |
| `API-SPEC-023` | `GetTicket` | `GetTicketRequest` | `GetTicketResponse` | Riwayat percakapan thread tiket via `ticketid`. |
| `API-SPEC-024` | `OpenTicket` | `OpenTicketRequest` | `OpenTicketResponse` | Buka tiket baru (`deptid`, `subject`, `message`, `priority`). |
| `API-SPEC-025` | `AddTicketReply` | `AddTicketReplyRequest` | `AddTicketReplyResponse` | Kirim balasan tiket (`ticketid`, `message`, `adminusername`). |
| `API-SPEC-026` | `AddTicketNote` | `AddTicketNoteRequest` | `AddTicketNoteResponse` | Tambah catatan internal staf tanpa terbaca klien. |
| `API-SPEC-027` | `UpdateTicket` | `UpdateTicketRequest` | `UpdateTicketResponse` | Ubah status, eskalasi departemen, staf flag. |
| `API-SPEC-028` | `DeleteTicket` | `DeleteTicketRequest` | `DeleteTicketResponse` | Hapus permanen tiket (*destructive*). |
| `API-SPEC-029` | `GetTicketPredefinedCats` | `GetPredefinedCatsRequest` | `GetPredefinedCatsResponse` | Kategori jawaban cepat. |
| `API-SPEC-030` | `GetTicketPredefinedReplies`| `GetPredefinedRepliesRequest`| `GetPredefinedRepliesResponse`| Templat jawaban cepat via `catid`. |
| `API-SPEC-031` | `GetSupportDepartments` | `GetDepartmentsRequest` | `GetDepartmentsResponse` | Daftar departemen bantuan dan metrik antrean. |
| `API-SPEC-032` | `GetSupportStatuses` | `GetSupportStatusesRequest`| `GetSupportStatusesResponse`| Daftar status tiket sistem. |

---

### Domain 5: Domain Lifecycle (`internal/model/domain.go`)

| ID Dokumen REST API | WHMCS Action | Struct Request Go | Struct Response Go | Keterangan & Parameter Utama |
|---|---|---|---|---|
| `API-SPEC-033` | `DomainWhois` | `DomainWhoisRequest` | `DomainWhoisResponse` | Lookup ketersediaan dan data WHOIS publik. |
| `API-SPEC-034` | `DomainRegister` | `DomainRegisterRequest` | `DomainRegisterResponse` | Perintahkan registrasi domain baru ke registrar. |
| `API-SPEC-035` | `DomainRenew` | `DomainRenewRequest` | `DomainRenewResponse` | Perpanjangan masa aktif domain. |
| `API-SPEC-036` | `DomainTransfer` | `DomainTransferRequest` | `DomainTransferResponse` | Inisiasi transfer masuk domain via EPP code. |
| `API-SPEC-037` | `DomainRelease` | `DomainReleaseRequest` | `DomainReleaseResponse` | Lepas domain UK ke IPS tag baru (*destructive*). |
| `API-SPEC-038` | `DomainRequestEPP` | `DomainRequestEPPRequest` | `DomainRequestEPPResponse` | Minta pengiriman kode otorisasi EPP. |
| `API-SPEC-039` | `DomainToggleIdProtect` | `ToggleIdProtectRequest` | `ToggleIdProtectResponse` | Toggle status proteksi privasi WHOIS. |
| `API-SPEC-040` | `DomainGetNameservers` | `GetNameserversRequest` | `GetNameserversResponse` | Ambil konfigurasi ns1-ns5 aktif dari registrar. |
| `API-SPEC-041` | `DomainUpdateNameservers` | `UpdateNameserversRequest` | `UpdateNameserversResponse` | Update alamat nameserver pada registrar. |
| `API-SPEC-042` | `DomainGetLockingStatus` | `GetLockStatusRequest` | `GetLockStatusResponse` | Ambil status registrar lock (`locked`/`unlocked`). |
| `API-SPEC-043` | `DomainUpdateLockingStatus`| `UpdateLockStatusRequest` | `UpdateLockStatusResponse` | Ubah status registrar lock. |
| `API-SPEC-044` | `GetTLDPricing` | `GetTLDPricingRequest` | `GetTLDPricingResponse` | Harga registrasi, perpanjangan, transfer per TLD. |

---

### Domain 6: Orders & Quotes (`internal/model/order.go`)

| ID Dokumen REST API | WHMCS Action | Struct Request Go | Struct Response Go | Keterangan & Parameter Utama |
|---|---|---|---|---|
| `API-SPEC-045` | `GetOrders` | `GetOrdersRequest` | `GetOrdersResponse` | Daftar pesanan pelanggan (filter status). |
| `API-SPEC-046` | `AddOrder` | `AddOrderRequest` | `AddOrderResponse` | Buat order baru (clientid, pid, domain, paymentmethod). |
| `API-SPEC-047` | `AcceptOrder` | `AcceptOrderRequest` | `AcceptOrderResponse` | Setujui pesanan dan picu provisioning. |
| `API-SPEC-048` | `PendingOrder` | `PendingOrderRequest` | `PendingOrderResponse` | Kembalikan order ke status Pending. |
| `API-SPEC-049` | `CancelOrder` | `CancelOrderRequest` | `CancelOrderResponse` | Batalkan pesanan (*destructive*). |
| `API-SPEC-050` | `FraudOrder` | `FraudOrderRequest` | `FraudOrderResponse` | Tandai pesanan sebagai fraud (*destructive*). |
| `API-SPEC-051` | `DeleteOrder` | `DeleteOrderRequest` | `DeleteOrderResponse` | Hapus permanen pesanan (*destructive*). |
| `API-SPEC-052` | `GetQuotes` | `GetQuotesRequest` | `GetQuotesResponse` | Daftar surat penawaran (quotes). |
| `API-SPEC-053` | `CreateQuote` | `CreateQuoteRequest` | `CreateQuoteResponse` | Buat surat penawaran biaya baru. |
| `API-SPEC-054` | `AcceptQuote` | `AcceptQuoteRequest` | `AcceptQuoteResponse` | Setujui quote dan konversi ke faktur. |
| `API-SPEC-055` | `DeleteQuote` | `DeleteQuoteRequest` | `DeleteQuoteResponse` | Hapus penawaran quote (*destructive*). |

---

### Domain 7: Server & Module Provisioning (`internal/model/provisioning.go`)

| ID Dokumen REST API | WHMCS Action | Struct Request Go | Struct Response Go | Keterangan & Parameter Utama |
|---|---|---|---|---|
| `API-SPEC-056` | `ModuleCreate` | `ModuleCreateRequest` | `ModuleCreateResponse` | Eksekusi pembuatan akun hosting di kontrol panel. |
| `API-SPEC-057` | `ModuleSuspend` | `ModuleSuspendRequest` | `ModuleSuspendResponse` | Tangguhkan akun hosting di server modul (*destructive*). |
| `API-SPEC-058` | `ModuleUnsuspend` | `ModuleUnsuspendRequest`| `ModuleUnsuspendResponse`| Aktifkan kembali akun hosting yang disuspend. |
| `API-SPEC-059` | `ModuleTerminate` | `ModuleTerminateRequest`| `ModuleTerminateResponse`| Hapus akun hosting dari server target (*destructive*). |
| `API-SPEC-060` | `ModuleChangePackage` | `ChangePackageRequest` | `ChangePackageResponse` | Ubah paket hosting di kontrol panel. |
| `API-SPEC-061` | `ModuleChangePassword`| `ChangePasswordRequest` | `ChangePasswordResponse` | Ubah password akun hosting di kontrol panel. |
| `API-SPEC-062` | `ModuleCustom` | `ModuleCustomRequest` | `ModuleCustomResponse` | Panggil fungsi kustom modul hosting pihak ketiga. |
| `API-SPEC-063` | `GetServers` | `GetServersRequest` | `GetServersResponse` | Daftar seluruh server hosting terkonfigurasi. |
| `API-SPEC-064` | `GetHealthStatus` | `GetHealthStatusRequest` | `GetHealthStatusResponse` | Status kesehatan sistem WHMCS dan cron. |

---

### Domain 8: System Administration & Marketing (`internal/model/system.go`, `internal/model/marketing.go`)

| ID Dokumen REST API | WHMCS Action | Struct Request Go | Struct Response Go | Keterangan & Parameter Utama |
|---|---|---|---|---|
| `API-SPEC-065` | `GetStats` | `GetStatsRequest` | `GetStatsResponse` | Ringkasan metrik statistik operasional sistem. |
| `API-SPEC-066` | `GetActivityLog` | `GetActivityLogRequest` | `GetActivityLogResponse` | Log audit aktivitas administratif. |
| `API-SPEC-067` | `LogActivity` | `LogActivityRequest` | `LogActivityResponse` | Tambahkan entri log audit baru secara manual. |
| `API-SPEC-068` | `GetPaymentMethods` | `GetPaymentMethodsRequest`| `GetPaymentMethodsResponse`| Daftar gateway pembayaran aktif. |
| `API-SPEC-069` | `GetCurrencies` | `GetCurrenciesRequest` | `GetCurrenciesResponse` | Daftar mata uang dan kurs nilai tukar. |
| `API-SPEC-070` | `GetAdminUsers` | `GetAdminUsersRequest` | `GetAdminUsersResponse` | Daftar user staf administrator. |
| `API-SPEC-071` | `GetToDoItems` | `GetToDoItemsRequest` | `GetToDoItemsResponse` | Daftar tugas administratif to-do item staf. |
| `API-SPEC-072` | `GetToDoItemStatuses`| `GetToDoStatusesRequest`| `GetToDoStatusesResponse`| Daftar status to-do yang valid. |
| `API-SPEC-073` | `UpdateToDoItem` | `UpdateToDoItemRequest` | `UpdateToDoItemResponse` | Perbarui status dan rincian to-do item. |
| `API-SPEC-074` | `SendEmail` | `SendEmailRequest` | `SendEmailResponse` | Kirim email kustom atau picu templat email sistem. |
| `API-SPEC-075` | `GetEmailTemplates` | `GetEmailTemplatesRequest`| `GetEmailTemplatesResponse`| Daftar templat email yang tersedia di sistem. |
| `API-SPEC-076` | `GetAffiliates` | `GetAffiliatesRequest` | `GetAffiliatesResponse` | Daftar akun afiliasi dan saldo komisi referal. |
| `API-SPEC-077` | `AffiliateActivate` | `AffiliateActivateRequest`| `AffiliateActivateResponse`| Aktivasi akun afiliasi baru untuk klien. |
| `API-SPEC-078` | `GetPromotions` | `GetPromotionsRequest` | `GetPromotionsResponse` | Daftar kode kupon promosi aktif. |
