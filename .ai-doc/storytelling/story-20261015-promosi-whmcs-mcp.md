# Story Output

**Created:** 2026-10-15  
**Author:** Sato (Master Copywriter ✍️)  

---

## Story Information

* **Story Type:** Pitch Narrative  
* **Framework Used:** Pitch Narrative (Persuasive #11)  
* **Category:** Persuasive Narratives  
* **Purpose:** Promosi repositori open-source `whmcs-mcp` (Go) untuk adopsi tim hosting dan developer AI agent  
* **Target Audience:** Cloud/Hosting Providers, Sysadmins WHMCS, dan AI Agent/MCP Engineers  
* **Constraint:** Maksimal 2.500 karakter untuk narasi inti, 100% evidence-based sesuai konstitusi  

---

## Story Structure

### Opening Hook
Kontras antara beban manual operasional WHMCS dengan overhead wrapper JavaScript konvensional.

### Core Narrative
Kelahiran `whmcs-mcp` berbasis arsitektur Go mandiri melalui metodologi TDD disiplin, menghadirkan biner 8.7 MB super ringan dengan paritas fungsional penuh.

### Key Story Beats
1. *The Friction*: Ribetnya tab admin WHMCS manual dan beban berat `node_modules`.
2. *The Breakthrough*: Transformasi Greenfield ke Golang (single static binary ~8.7 MB, <15 MB RAM, 0ms startup).
3. *The Architecture*: Isolasi ketat `os.Stderr` (logging) vs `os.Stdout` (murni JSON-RPC 2.0).
4. *The Proof (v0.2.0)*: 66 Tools, 11 Resources, 8 Prompts, HTTP Basic Auth, dan flattened `vmid` (default "000").
5. *The Call to Action*: Link repository dan rilis multi-platform open-source.

### Emotional Arc
Frustrasi operasional → Kejernihan solusi rekayasa → Keyakinan berbasis bukti (165+ tests PASS) → Dorongan adopsi.

### Resolution / Call to Action
Instalasi biner mandiri langsung ke Claude Desktop / Cursor tanpa friksi dependensi.

---

## Complete Story (Narasi Inti — 1.940 Karakter)

Bayangkan ini: server WHMCS Anda sedang melayani ribuan transaksi, tiket keluhan masuk beruntun di tengah malam, dan Anda masih harus membuka puluhan tab browser hanya untuk memeriksa status satu VPS.

Saat AI assistant seperti Claude Desktop dan Cursor hadir, janji otomatisasi terdengar manis. Namun kenyataan teknisnya sering mengecewakan: wrapper API lama menuntut ratusan megabyte dependensi `node_modules`, proses startup lambat, dan konsumsi RAM yang boros.

Inilah alasan **`whmcs-mcp`** dibangun ulang dari nol dengan arsitektur Golang.

Bukan sekadar porting biasa. Proyek ini mematuhi disiplin konstitusi teknis ketat berbasis bukti (*evidence-based*) dan 100% dibangun dengan metodologi Test-Driven Development (TDD).

Apa yang membedakannya?

1. **Biner Tunggal Mandiri ~8.7 MB**: 
   Kompilasi native Go 1.22+. Tanpa interpreter eksternal, tanpa `node_modules`, konsumsi memori kurang dari 15 MB RAM, dan *cold start* instan dalam hitungan milidetik.
2. **Protokol Stdio Murni & Stabil**: 
   Mematuhi standar MCP 2024-11-05 secara presisi. Seluruh log dialirkan ke `os.Stderr`, menjaga `os.Stdout` 100% steril untuk JSON-RPC 2.0 tanpa risiko korupsi paket data.
3. **Senjata Lengkap (66 Tools, 11 Resources, 8 Prompts)**: 
   Mencakup seluruh siklus hosting: manajemen klien, penagihan invoice, domain, tiket support, hingga otomatisasi modul server cPanel dan Proxmox.
4. **Teruji Nyata di Versi v0.2.0**: 
   Dilengkapi dukungan HTTP Basic Auth untuk menembus web server Nginx yang terproteksi, serta ekstraksi otomatis field `vmid` (fallback `"000"`) yang membebaskan LLM dari parsing nested array PHP yang rumit.

Dengan 165+ automated test assertions berstatus lulus 100%, Anda tidak lagi membeli janji fiktif. Anda mendapatkan keandalan nyata yang siap tempur di lini produksi.

Cukup katakan: *"Claude, audit tagihan yang jatuh tempo dan periksa VMID aktif klien 109524."* 
Semua tereksekusi dalam kedipan mata langsung dari desktop Anda.

Biner pra-kompilasi multi-platform (Linux, macOS, Windows) kini tersedia bebas secara open-source.

Kunjungi dan pasang sekarang:  
👉 https://github.com/vfat/whmcs-mcp

---

## Story Elements Analysis

* **Character / Voice:** Sato (The Word Crafter) — Percakapan langsung, lugas, anti-jargon, dan mengedepankan pembuktian nyata.
* **Conflict / Tension:** Beban kerja admin hosting + beratnya runtime stack JS vs kebutuhan integrasi AI yang cepat dan stabil.
* **Transformation:** Dari pengelolaan manual & skrip berat menjadi otomatisasi AI instan melalui biner mandiri ringan.
* **Key Messages:** Go native performance, zero node_modules, TDD-backed reliability, v0.2.0 production features.

---

## Variations & Adaptations

### Short Version (LinkedIn / Twitter / Social Pitch — 420 Karakter)
> Kelola WHMCS lewat Claude Desktop tanpa `node_modules`! 🚀  
> 
> **`whmcs-mcp` (Go)** mengubah instalasi WHMCS Anda menjadi ekosistem AI-ready:  
> ⚡ Biner mandiri ~8.7 MB (RAM <15 MB)  
> 🛠️ 66 Tools, 11 Resources, 8 Prompts  
> 🛡️ 100% lolos TDD (165+ tests)  
> 🚀 Rilis v0.2.0: HTTP Basic Auth & auto-flatten VMID VPS  
> 
> Open-source & multi-platform: https://github.com/vfat/whmcs-mcp

### Medium Version (Developer Forum / Community Post — 1.100 Karakter)
> **WHMCS MCP Server (Go) v0.2.0 Released: Single Binary AI Management for Hosting Providers**
> 
> Capek bolak-balik buka admin panel WHMCS untuk cek tiket, tagihan invoice, atau provisioning server?
> 
> Kami baru saja merilis `whmcs-mcp` versi v0.2.0—server Model Context Protocol (MCP) berbasis Golang untuk menghubungkan Claude Desktop, Cursor, dan AI Agent langsung ke WHMCS Anda:
> 
> • **Ultra Lightweight**: Single binary ~8.7 MB, hemat memori (<15 MB RAM), zero runtime dependencies.
> • **Strict Protocol Isolation**: Output JSON-RPC murni di stdout, logging aman di stderr.
> • **66 Tools & 11 Resources**: Cakupan 100% untuk Client, Invoice, Tiket, Domain, Order, dan Server Module.
> • **New in v0.2.0**: Support HTTP Basic Auth (Nginx htpasswd bypass) dan ekstraksi langsung top-level `vmid` (fallback "000") untuk VPS.
> • **Rock-Solid TDD**: Diverifikasi dengan 165+ automated test assertions.
> 
> Tersedia biner siap pakai untuk Linux, macOS (termasuk Apple Silicon M-series), dan Windows.
> 
> Repo: https://github.com/vfat/whmcs-mcp

---

## Usage Guidelines

* **Best Channels:** GitHub Readme spotlight, Reddit (r/webhosting, r/golang, r/selfhosted), Dev.to, LinkedIn tech groups, Telegram hosting channels.
* **Tone & Voice Notes:** Tegas, berbasis bukti riil (*evidence-based*), profesional, tidak melebih-lebihkan hal yang tidak teruji di kode.

---

_Story crafted using the AI Documentor storytelling add-on & constitution of whmcs-mcp._
