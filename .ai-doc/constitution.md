# Constitution for `whmcs-mcp` (Golang Implementation)

## Purpose

Konstitusi ini mendefinisikan aturan dan standar dokumentasi serta perencanaan berbasis bukti (*evidence-based documentation & planning*) untuk proyek **whmcs-mcp** berbasis bahasa pemrograman **Golang**.

Proyek ini merupakan implementasi ulang / transformasi greenfield dari arsitektur referensi TypeScript (`whmcs-mcp-tool`) ke dalam runtime Golang yang berkinerja tinggi, berukuran biner mandiri (*single binary*), hemat memori, dan mendukung konkurensi native.

Baseline mengadopsi standar umum dari `ai-documentor`.

---

## 1. Workspace Rules

- Ruang kerja dokumentasi dan perancangan berada di `.ai-doc/`.
- File control plane:
  - `.ai-doc/3p.md` (Progress, Plan, Pending tracker)
  - `.ai-doc/constitution.md` (Aturan dan pedoman lokal proyek)
  - `.ai-doc/project-overview.md` (Dokumen perancangan awal sistem greenfield)
- Artefak turunan (`Dokumentasi-Fitur.md`, `Dokumentasi-Komponen-Usecase.md`, `C4-Component-Diagrams.md`, `desain-component-document/`, `desain-database-document/`, `rest-api-doc/`) **hanya dibuat bila diminta secara eksplisit oleh user** (*Only-On-Request Rule*).
- Sumber rujukan spesifikasi awal bersumber dari audit artefak komprehensif di `plan/whmcs-mcp-tool/.ai-doc/` (78 API Action specs, DCD 01-08, Annex L-001 s/d L-003).

---

## 2. Progress Tracking

- File `3p.md` harus selalu dibaca dan diperbarui setelah menyelesaikan langkah bermakna.
- Struktur standar `3p.md`:
  - `Progress`
  - `Plan`
  - `Pending`

---

## 3. Evidence Rules

- Tidak boleh mendokumentasikan klaim atau asumsi tanpa bukti.
- Pada fase Greenfield Planning:
  - Setiap kebutuhan fungsional dan kontrak tool harus traceable ke dokumen spesifikasi referensi di `plan/whmcs-mcp-tool/.ai-doc/`.
- Pada fase Implementasi & Testing:
  - Setiap klaim harus traceable ke source code Go (`*.go`), modul `go.mod`, unit/integration test, atau konfigurasi nyata.
- Unknowns harus ditandai secara tegas: `Draft`, `Partial`, `Placeholder`, `Asumsi`, atau `Perlu Dikonfirmasi`.

---

## 4. Klasifikasi Komponen & Objek (Arsitektur Golang)

- **Boundary**: Antarmuka transport komunikasi (stdio transport JSON-RPC 2.0 via `os.Stdin` / `os.Stdout`, CLI flags/environment loader).
- **Control**: Engine orkestrasi MCP Server, dispatcher tool/prompt/resource, validator parameter, serta HTTP Client WHMCS (`WhmcsClient`) yang menangani POST ke `includes/api.php` dan serialisasi `url.Values`.
- **Entity**: Data transfer objects (DTO), struct skema request/response WHMCS, metadata MCP Tool/Prompt/Resource, dan error definitions. API endpoint WHMCS diklasifikasikan sebagai `Entity`.
- **Infrastructure**: Go runtime (Go 1.22+), container scratch/alpine minimalis, OS process pipe.

---

## 5. TDD Policy (Enabled)

- **Status**: **Enabled** (Dikonfirmasi oleh user pada 2026-09-30).
- **Scope**: Pengembangan Greenfield Golang `whmcs-mcp`.
- **Test Runner**: Go built-in testing (`go test -v ./...`).
- **Siklus**: **RED → GREEN → REFACTOR**:
  1. *RED*: Tidak ada production code untuk behavior baru sebelum test ditulis dan dipastikan gagal dengan bukti nyata (*failure evidence*).
  2. *GREEN*: Tulis implementasi seminimal mungkin hingga test berhasil lulus (*passing evidence*).
  3. *REFACTOR*: Rapikan struktur dan performa kode tanpa mengubah behavior, dengan verifikasi regresi test tetap lulus.
- **Pusat Kontrol TDD**: Dilacak dan diverifikasi secara berkala di `.ai-doc/tdd-overview.md`.
- **Integritas Bukti**: Status RED/GREEN wajib menyertakan command, exit code, dan cuplikan output nyata. Tidak boleh mengklaim status tanpa hasil eksekusi riil.
