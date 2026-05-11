# Vowtera

Admin panel berbasis web. Frontend: **Nuxt 4**. Backend: **Go + Chi + GORM**.

> Repo: https://github.com/richoda/vowtera | Branch utama: `master`

---

## Tech Stack

| Bagian | Teknologi |
|--------|-----------|
| Frontend | Nuxt 4, Vue 3, Tailwind CSS, Nuxt Icon (MDI) |
| Backend | Go, Chi v5, sqlx, pgx |
| Database | PostgreSQL (rencana: Supabase) |
| Auth | JWT (golang-jwt v5) |

---

## Struktur Folder

```
vowtera/
├── frontend/
│   ├── layouts/         # default (sidebar+header), plain (login)
│   ├── components/      # SideBar.vue, Header.vue
│   ├── composables/     # useSidebar.ts
│   └── pages/           # index, login, overview, about, contact
│
└── backend/
    ├── main.go
    ├── config/          # load env per environment
    ├── database/        # koneksi PostgreSQL + auto-migrate
    ├── models/          # struct User
    ├── middleware/      # JWT auth
    ├── handlers/        # Login, Me
    └── routes/          # semua route
```

---

## Cara Jalankan

```bash
# Frontend — http://localhost:3000
cd frontend
npm install
npm run dev

# Backend — http://localhost:8080
cd backend
go run main.go
```

---

## Environment

Salin file example lalu isi nilainya:

```bash
cp frontend/.env.example frontend/.env.development
cp backend/.env.example  backend/.env.development
```

**Frontend** (`frontend/.env.example`):

| Variabel | Keterangan |
|----------|------------|
| `NUXT_PUBLIC_API_BASE_URL` | URL backend (default: `http://localhost:8080`) |
| `NUXT_PUBLIC_APP_NAME` | Nama aplikasi |
| `NUXT_PUBLIC_APP_ENV` | `development` / `staging` / `production` |

**Backend** (`backend/.env.example`):

| Variabel | Keterangan |
|----------|------------|
| `APP_ENV` | Nama environment (default: `development`) |
| `APP_PORT` | Port server (default: `8080`) |
| `DB_HOST`, `DB_PORT`, `DB_NAME` | Koneksi PostgreSQL |
| `DB_USER`, `DB_PASSWORD` | Kredensial database |
| `JWT_SECRET` | Secret key JWT — **ganti di production!** |

---

## API Endpoints

| Method | Path | Auth | Keterangan |
|--------|------|------|------------|
| `GET` | `/health` | — | Cek server hidup |
| `POST` | `/api/login` | — | Login, dapat JWT token |
| `GET` | `/api/me` | Bearer token | Info user yang login |

---

## Design System (Frontend)

| Token | Hex | Digunakan untuk |
|-------|-----|-----------------|
| Navy | `#0F1C33` | Teks utama, background gelap |
| Amber | `#B8873A` | Aksen, border aktif, tombol |
| Krem | `#F7F4EF` | Background sidebar & input |
| Abu | `#B8ADA1` | Border, placeholder |

Semua ikon pakai `<Icon name="mdi:nama-ikon" />`.

---

## Menambah Halaman Baru (Frontend)

1. Buat `pages/nama.vue`
2. Tambah entry di `SideBar.vue`
3. Tambah entry di object `pages` di `Header.vue`

## Menambah Endpoint Baru (Backend)

1. Buat handler di `handlers/nama.go`
2. Daftarkan di `routes/routes.go`
3. Jika butuh model baru, tambahkan ke `AutoMigrate` di `database/database.go`

---

## Status

- [x] Frontend: layout, sidebar, header, login, overview, about, contact
- [x] Backend: HTTP server, routing Chi, koneksi GORM, JWT auth
- [ ] Database: setup Supabase
- [ ] Login frontend connect ke backend
- [ ] Deploy (Vercel + Fly.io/Railway)
- [ ] GitHub Actions CI/CD
