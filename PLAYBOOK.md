# Vowtera — Playbook

> Baca file ini di awal setiap session. Berisi konteks yang tidak bisa dibaca dari kode.
> Detail teknis lengkap ada di `README.md` (gabungan) dan `backend/DOKUMENTASI.md`.

---

## Project

Admin panel web untuk internal. Stack: **Nuxt 4** (frontend) + **Go/Chi/sqlx** (backend) + **PostgreSQL** (rencana Supabase).

- Repo: https://github.com/richoda/vowtera | Branch: `master`
- Frontend: `localhost:3000` | Backend: `localhost:8080`

---

## Status

### Selesai
- [x] Frontend lengkap — layout, sidebar, header, login, overview, about, contact
- [x] Backend Go — HTTP server (Chi), koneksi DB (sqlx + pgx), JWT auth, model User
- [x] Sistem environment — `.env.{environment}` per stage, semua `.env.example` ada

### Belum Dikerjakan (urutan prioritas)
- [ ] **Setup Supabase** — koneksi PostgreSQL ke cloud, isi `.env` dengan credentials Supabase
- [ ] **Connect login frontend** — ganti hardcode credentials dengan API call ke `POST /api/login`
- [ ] **Deploy frontend** — ke Vercel
- [ ] **Deploy backend** — ke Fly.io atau Railway
- [ ] **GitHub Actions** — CI/CD auto-deploy tiap push ke `master`

---

## Keputusan Teknis

| Keputusan | Alasan |
|-----------|--------|
| Chi bukan Gin/Fiber | Ringan, kompatibel stdlib, tidak over-engineered untuk skala ini |
| sqlx bukan GORM | Developer terbiasa SQL, lebih suka kontrol query langsung |
| JWT stateless | Tidak perlu tabel session di DB, cocok untuk admin panel sederhana |
| `.env.staging/.production` tidak di-commit | Berisi kredensial sensitif |
| `NUXT_PUBLIC_` prefix di frontend | Satu-satunya cara agar env var aman diakses di browser |
| Soft delete manual (`deleted_at`) | Data tidak benar-benar dihapus, ada audit trail |
| Error login selalu pesan generik | Mencegah user enumeration attack |
| Hosting: Vercel + Fly.io + Supabase | Semua gratis untuk tahap awal |

---

## Design System — Frontend

### Warna

Hanya boleh pakai empat token ini. **Jangan pakai kelas Tailwind generik** (`red-500`, `blue-600`, dll.).

| Token | Hex | Digunakan untuk |
|-------|-----|-----------------|
| Navy | `#0F1C33` | Teks utama, background gelap, hover button |
| Amber | `#B8873A` | Aksen aktif, border sidebar, tombol primary |
| Krem | `#F7F4EF` | Background sidebar, input, halaman login |
| Abu | `#B8ADA1` | Border, placeholder, teks sekunder |
| Putih | `#FFFFFF` | Background card, header |

### Tipografi

| Penggunaan | Class |
|------------|-------|
| Heading halaman | `text-2xl font-bold text-[#0F1C33]` |
| Label / subheading | `text-sm font-semibold text-[#0F1C33]` |
| Body teks | `text-sm text-[#0F1C33]` |
| Teks sekunder | `text-sm text-[#B8ADA1]` |
| Teks pada background gelap | `text-white` |

### Komponen

**Card:**
```html
<div class="bg-white rounded-xl border border-[#B8ADA1] p-6">
```

**Tombol primary:**
```html
<button class="bg-[#0F1C33] text-white px-4 py-2 rounded-lg hover:bg-[#B8873A] transition-colors duration-200">
```

**Input:**
```html
<input class="w-full bg-[#F7F4EF] border border-[#B8ADA1] rounded-lg px-4 py-2 text-[#0F1C33] placeholder-[#B8ADA1] focus:outline-none focus:border-[#B8873A]">
```

**Badge / status:**
```html
<span class="text-xs font-medium px-2 py-1 rounded-full bg-[#B8873A]/10 text-[#B8873A]">
```

### Layout & Ukuran

| Elemen | Nilai | Class |
|--------|-------|-------|
| Sidebar expanded | 256px | `w-64` |
| Sidebar collapsed | 96px | `w-24` |
| Header height | 80px (top-20) | `h-20` |
| Content margin-top | 80px | `mt-20` |
| Content margin-left (expanded) | 256px | `ml-64` |
| Content margin-left (collapsed) | 96px | `ml-24` |

### Ikon

Semua ikon via Nuxt Icon dengan prefix `mdi:`:
```html
<Icon name="mdi:nama-ikon" class="text-xl" />
```
Referensi: https://icon-sets.iconify.design/mdi/

### Animasi

- Hover interaktif selalu pakai `transition-colors duration-200`
- Pergantian gambar (logo sidebar) pakai `transition-opacity duration-300` — jangan `v-if/v-else`
- Jangan animasi tanpa `transition`

---

## Design System — Backend

### Format Response

Semua endpoint mengembalikan JSON. Gunakan format konsisten berikut:

**Sukses (data tunggal):**
```json
{ "token": "...", "user": { ... } }
```

**Sukses (data list) — untuk endpoint list di masa depan:**
```json
{ "data": [ ... ], "total": 10 }
```

**Error:**
```
Plain text — bukan JSON
Contoh: "email atau password salah"
```

### HTTP Status Code

| Situasi | Kode |
|---------|------|
| Sukses GET / operasi berhasil | `200 OK` |
| Request body tidak valid / field kurang | `400 Bad Request` |
| Token tidak ada / tidak valid / expired | `401 Unauthorized` |
| Akses ke resource milik user lain | `403 Forbidden` |
| Data tidak ditemukan di DB | `404 Not Found` |
| Kesalahan server / query gagal | `500 Internal Server Error` |

### Penamaan

| Konteks | Konvensi | Contoh |
|---------|----------|--------|
| Nama kolom DB | `snake_case` | `created_at`, `user_id` |
| JSON response key | `snake_case` | `"created_at"`, `"user_id"` |
| Nama file Go | `snake_case` | `auth.go`, `user_handler.go` |
| Nama fungsi / variabel Go | `camelCase` | `ListUsers`, `jwtSecret` |
| URL path | `kebab-case` | `/api/user-roles` |
| Placeholder SQL | `$1`, `$2` | `WHERE id = $1` |

### Pola Route (RESTful)

| Operasi | Method | Path |
|---------|--------|------|
| List semua | `GET` | `/api/users` |
| Detail satu | `GET` | `/api/users/{id}` |
| Buat baru | `POST` | `/api/users` |
| Update | `PUT` | `/api/users/{id}` |
| Hapus | `DELETE` | `/api/users/{id}` |

### Query SQL

- Selalu tambahkan `AND deleted_at IS NULL` pada SELECT dan UPDATE
- Gunakan `$1, $2, ...` sebagai placeholder — **bukan** `?` (itu sintaks MySQL)
- Soft delete: `UPDATE users SET deleted_at = NOW() WHERE id = $1`
- Update: selalu sertakan `updated_at = NOW()` dalam query UPDATE

---

## Konvensi Kode

**Frontend:**
- State sidebar selalu lewat `useSidebar()` — jangan buat `ref()` lokal
- Gunakan `<script setup>` di semua komponen — bukan Options API
- Navigasi internal: `<NuxtLink>` | Redirect programatik: `navigateTo()`
- Halaman baru tanpa sidebar: `definePageMeta({ layout: 'plain' })`

**Backend:**
- Handler baru → `handlers/`, daftarkan di `routes/routes.go`
- Model baru → `models/`, tambahkan ke fungsi `migrate()` di `database/database.go`
- Route yang butuh auth → masuk ke `r.Group` dengan `middleware.Auth`
- Jangan baca `os.Getenv` di handler — selalu lewat struct `Config`

---

## Catatan untuk Claude

- Login frontend saat ini **hardcode** — belum connect ke backend. Jangan terkecoh seolah auth sudah terintegrasi.
- Backend belum pernah dijalankan dengan database nyata — masih menunggu Supabase di-setup.
- `go.mod` module name adalah `admin-api`, bukan `vowtera` — pakai ini untuk import antar package.
- Saat tambah endpoint baru, selalu tanya apakah perlu protected (masuk group JWT) atau public.
- Jika ada perubahan signifikan, update bagian **Status** di file ini.
