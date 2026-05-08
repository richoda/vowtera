# Vowtera — Playbook

Dokumen ini dibuat agar session Claude berikutnya langsung paham konteks project tanpa penjelasan ulang.

---

## Tentang Project

**Vowtera** adalah admin panel berbasis web.

| | |
|---|---|
| **Repo GitHub** | https://github.com/richoda/vowtera |
| **Akun GitHub** | richoda |
| **Branch utama** | master |

---

## Tech Stack

| Bagian | Teknologi |
|---|---|
| Frontend | Nuxt 4, Vue 3, Tailwind CSS, Nuxt Icon (MDI) |
| Backend | Go (masih awal, baru `main.go`) |
| Database | Belum ada (rencana PostgreSQL) |
| Hosting | Belum deploy (rencana Vercel untuk frontend) |

---

## Struktur Folder

```
vowtera/
├── frontend/        # Nuxt 4 app
├── backend/         # Go API (masih awal)
├── .gitignore
└── PLAYBOOK.md      # file ini
```

---

## Status Project

### Sudah Selesai

- [x] Frontend Nuxt: layout, sidebar, header, halaman login, overview, about, contact
- [x] Git init + push ke GitHub
- [x] Setup environment: `.env.example`, `.env.staging`, `.env.production`
- [x] `nuxt.config.ts` sudah pakai `runtimeConfig` untuk env vars
- [x] Script npm per environment (`dev:staging`, `build:staging`, `build:production`)
- [x] Dokumentasi di `DOKUMENTASI.md` dan `README.md` sudah diupdate

### Belum Dikerjakan

- [ ] Backend Go: routing, koneksi database, API endpoints
- [ ] Database: setup PostgreSQL (rencana pakai Supabase gratis)
- [ ] Autentikasi: login masih hardcode, belum connect ke backend
- [ ] Deploy: belum ke server manapun (rencana Vercel untuk frontend)
- [ ] GitHub Actions: belum ada CI/CD

---

## Keputusan Teknis yang Sudah Dibuat

| Keputusan | Alasan |
|---|---|
| `.env.staging` dan `.env.production` tidak di-commit | Berisi kredensial sensitif |
| Hanya `.env.example` yang masuk git | Sebagai template untuk developer lain |
| Pakai `NUXT_PUBLIC_` prefix untuk env vars | Agar aman diakses di sisi browser |
| Rencana hosting: Vercel (frontend) + Fly.io/Railway (backend) + Supabase (DB) | Gratis untuk tahap awal |

---

## Cara Jalankan Lokal

```bash
# Frontend
cd frontend
npm install
npm run dev          # http://localhost:3000

# Backend
cd backend
go run main.go       # http://localhost:8080
```

---

## Variabel Environment

Salin dari example lalu isi nilainya:

```bash
cp frontend/.env.example frontend/.env.staging
cp frontend/.env.example frontend/.env.production
cp backend/.env.example backend/.env.staging
cp backend/.env.example backend/.env.production
```

| Variabel Frontend | Keterangan |
|---|---|
| `NUXT_PUBLIC_API_BASE_URL` | URL backend API |
| `NUXT_PUBLIC_APP_NAME` | Nama aplikasi |
| `NUXT_PUBLIC_APP_ENV` | `development` / `staging` / `production` |

| Variabel Backend | Keterangan |
|---|---|
| `APP_ENV` | Nama environment |
| `APP_PORT` | Port server (default 8080) |
| `DB_HOST`, `DB_NAME`, `DB_USER`, `DB_PASSWORD` | Koneksi database |
| `JWT_SECRET` | Secret key JWT |

---

## Lanjutan yang Disarankan

Urutan pengerjaan berikutnya yang direkomendasikan:

1. **Setup backend Go** — routing HTTP, koneksi ke database
2. **Setup database** — pakai Supabase (gratis, PostgreSQL)
3. **Sambungkan login** — ganti hardcode dengan API call ke backend
4. **Deploy frontend** — ke Vercel
5. **Deploy backend** — ke Fly.io atau Railway
6. **Setup GitHub Actions** — auto-deploy setiap push

---

## Cara Pakai Playbook Ini

Di awal session baru, cukup bilang ke Claude:

> *"Baca PLAYBOOK.md dan lanjutkan project Vowtera"*

Claude akan baca file ini dan langsung paham konteks tanpa perlu penjelasan ulang.
