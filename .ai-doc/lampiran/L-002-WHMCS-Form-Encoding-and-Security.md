# Lampiran L-002: Serialisasi Form WHMCS & Keamanan Kredensial

> **Topik**: Serialisasi Parameter Nested/Array ke `application/x-www-form-urlencoded` dan Standar Pengamanan Kredensial API  
> **Workspace**: `/home/ubuntu/workspace/plan/whmcs-mcp`  
> **Target Sistem**: `whmcs-mcp` (Go 1.22+)  
> **Tanggal**: 2026-09-30  

---

## 1. Masalah Encoding Parameter WHMCS

API WHMCS (`includes/api.php`) dibangun dengan bahasa PHP yang mengharapkan input body berupa `application/x-www-form-urlencoded`. Ketika parameter memiliki struktur hierarki (array atau map), PHP menggunakan konvensi penamaan kurung siku (`[]`), misalnya:

```text
customfields[0]=Value1&customfields[1]=Value2
nameservers[ns1]=ns1.example.com&nameservers[ns2]=ns2.example.com
```

Pustaka standar Go `url.Values` secara bawaan hanya menangani pemetaan flat `string -> []string`. Untuk itu, paket `internal/whmcs/serializer.go` bertugas melakukan *flattening* rekursif terhadap struct atau map Go menjadi `url.Values`.

---

## 2. Aturan Serializer (`flattenParams`)

1. **Skalar (String, Int, Bool)**:
   - Dipetakan langsung: `key=value` (contoh: `action=GetClientsDetails&clientid=101`).
2. **Slice / Array**:
   - Jika slice berisi skalar, diserialisasi dengan indeks: `prefix[0]=val0&prefix[1]=val1`.
   - Jika slice berisi objek/struct (misal line items invoice), diserialisasi dengan notasi indeks dan properti: `itemdescription[0]=Item1&itemamount[0]=10.00`.
3. **Map / Struct Bertingkat**:
   - Diserialisasi dengan nested key: `prefix[field]=value`.

---

## 3. Protokol Keamanan & Whitelisting IP

1. **Metode Autentikasi Modern**:
   - WHMCS 7.2+ menggunakan pasangan `identifier` dan `secret` (API Role Credentials) yang dibuat di menu *Setup > Staff Management > Administrator Roles / API Roles*.
2. **Penggunaan `$api_access_key`**:
   - Jika file `configuration.php` WHMCS mendefinisikan variabel `$api_access_key`, maka setiap request API wajib menyertakan nilai tersebut pada parameter `accesskey`.
3. **IP Access Restriction**:
   - WHMCS menerapkan pembatasan IP ketat (*IP Whitelisting*) pada menu *Setup > General Settings > Security > API IP Access Restriction*.
   - Server biner Go yang berjalan di mesin lokal atau VPS wajib didaftarkan IP publik/keluarnya pada daftar whitelist WHMCS.
4. **Pencegahan Kebocoran Kredensial**:
   - Kredensial dibaca dari environment variable (`WHMCS_API_IDENTIFIER`, `WHMCS_API_SECRET`, `WHMCS_API_ACCESS_KEY`).
   - Tidak pernah menuliskan kredensial ke berkas log atau pesan error JSON-RPC.
