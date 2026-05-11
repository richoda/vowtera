# Dokumentasi Backend Vowtera

Dokumentasi teknis lengkap untuk backend Vowtera — Go REST API dengan Chi, GORM, dan JWT.

---

## Daftar Isi

1. [Gambaran Umum](#1-gambaran-umum)
2. [Tech Stack & Dependency](#2-tech-stack--dependency)
3. [Struktur Folder](#3-struktur-folder)
4. [Alur Kerja Aplikasi](#4-alur-kerja-aplikasi)
5. [File per File](#5-file-per-file)
   - [main.go](#51-maingo)
   - [config/config.go](#52-configconfiggo)
   - [database/database.go](#53-databasedatabasego)
   - [models/user.go](#54-modelsusergo)
   - [middleware/auth.go](#55-middlewareauthgo)
   - [handlers/auth.go](#56-handlersauthgo)
   - [routes/routes.go](#57-routesroutesgo)
6. [API Endpoints](#6-api-endpoints)
7. [Sistem Environment](#7-sistem-environment)
8. [Konsep Penting](#8-konsep-penting)
9. [Cara Menambah Fitur Baru](#9-cara-menambah-fitur-baru)

---

## 1. Gambaran Umum

Backend Vowtera adalah REST API yang melayani frontend Nuxt. Tugasnya:
- Menerima request HTTP dari browser/frontend
- Memproses logika bisnis (autentikasi, data, dll.)
- Berkomunikasi dengan database PostgreSQL
- Mengembalikan response dalam format JSON

```
Browser/Frontend
      │
      │ HTTP Request
      ▼
  Chi Router  ──► Middleware (Logger, JWT, dll.)
      │
      ▼
   Handler     ──► Model / Database
      │
      │ JSON Response
      ▼
Browser/Frontend
```

---

## 2. Tech Stack & Dependency

### Dependency Utama (`go.mod`)

| Package | Versi | Fungsi |
|---------|-------|--------|
| `github.com/go-chi/chi/v5` | v5.2.5 | HTTP router — mendefinisikan URL path dan method |
| `github.com/jmoiron/sqlx` | v1.4.0 | Ekstensi `database/sql` — scan result ke struct, named query |
| `github.com/jackc/pgx/v5` | v5.6.0 | Driver PostgreSQL native, dipakai via `pgx/stdlib` |
| `github.com/golang-jwt/jwt/v5` | v5.3.1 | Generate dan verifikasi JWT token |
| `github.com/joho/godotenv` | v1.5.1 | Load file `.env` ke dalam environment OS |
| `golang.org/x/crypto` | v0.31.0 | Bcrypt untuk hash password |

### Kenapa Chi bukan Gin/Fiber?

Chi dipilih karena:
- **Ringan** — hanya router, tidak ada magic tersembunyi
- **Kompatibel stdlib** — middleware Chi bisa dipakai di semua project Go lain
- **Mudah dipelajari** — pola `r.Get`, `r.Post`, `r.Group` sangat intuitif

### Kenapa sqlx bukan GORM?

sqlx dipilih karena developer terbiasa dengan SQL. Keunggulannya:
- Tulis SQL langsung — tidak ada query yang tersembunyi di balik ORM
- `db.Get(&user, "SELECT ...", arg)` — scan satu row ke struct secara otomatis
- `db.Select(&users, "SELECT ...", arg)` — scan banyak row ke slice struct
- Pakai `$1, $2` sebagai placeholder PostgreSQL (parameterized, aman dari SQL injection)

---

## 3. Struktur Folder

```
backend/
├── main.go                  # Entry point — titik mulai program
├── go.mod                   # Daftar dependency (seperti package.json di Node)
├── go.sum                   # Checksum dependency (jangan diedit manual)
│
├── .env.example             # Template env — DI-COMMIT ke git
├── .env.staging             # Env staging — TIDAK di-commit
├── .env.production          # Env production — TIDAK di-commit
│
├── config/
│   └── config.go            # Baca env vars, bungkus dalam struct Config
│
├── database/
│   └── database.go          # Buka koneksi ke PostgreSQL, jalankan AutoMigrate
│
├── models/
│   └── user.go              # Definisi tabel users (struct + method password)
│
├── middleware/
│   └── auth.go              # Cek JWT token sebelum request masuk ke handler
│
├── handlers/
│   └── auth.go              # Logika Login dan Me
│
└── routes/
    └── routes.go            # Peta semua URL → handler
```

**Prinsip struktur ini:** setiap folder punya satu tanggung jawab. `models` hanya tahu soal data, `handlers` hanya tahu soal request/response, `routes` hanya tahu soal URL mapping.

---

## 4. Alur Kerja Aplikasi

### Startup (saat `go run main.go`)

```
main.go
  │
  ├─ 1. config.Load()        → baca APP_ENV, load .env.{env}, bungkus ke struct
  │
  ├─ 2. database.Connect()   → buka koneksi PostgreSQL, jalankan AutoMigrate
  │
  ├─ 3. handlers.SetJWTSecret() → simpan JWT secret ke package handlers
  │
  ├─ 4. routes.Setup()       → buat router Chi, daftarkan semua middleware & route
  │
  └─ 5. http.ListenAndServe() → mulai terima request HTTP
```

### Request Login (`POST /api/login`)

```
Request masuk
  │
  ├─ Chi Router → cocokkan path /api/login
  │
  ├─ Middleware: Logger (catat ke terminal)
  ├─ Middleware: Recoverer (tangkap panic)
  ├─ Middleware: RequestID (beri ID unik)
  │
  ├─ handlers.Login()
  │    ├─ Decode JSON body → loginRequest{email, password}
  │    ├─ Query DB: SELECT * FROM users WHERE email = ?
  │    ├─ Cek password dengan bcrypt
  │    ├─ Generate JWT token (berlaku 24 jam)
  │    └─ Encode response JSON → {token, user}
  │
  └─ Response 200 OK + JSON
```

### Request Protected (`GET /api/me`)

```
Request masuk
  │
  ├─ Chi Router → cocokkan path /api/me
  │
  ├─ Middleware: Logger, Recoverer, RequestID
  ├─ Middleware: Auth (CEK JWT) ──► kalau invalid → 401, STOP
  │    ├─ Ambil header Authorization: Bearer <token>
  │    ├─ Parse & verifikasi token
  │    ├─ Ambil user ID dari claims
  │    └─ Simpan user ID ke context request
  │
  ├─ handlers.Me()
  │    ├─ Ambil user ID dari context
  │    ├─ Query DB: SELECT * FROM users WHERE id = ?
  │    └─ Encode response JSON → {user}
  │
  └─ Response 200 OK + JSON
```

---

## 5. File per File

### 5.1 `main.go`

```go
package main

import (
    "fmt"
    "log"
    "net/http"

    "admin-api/config"
    "admin-api/database"
    "admin-api/handlers"
    "admin-api/routes"
)

func main() {
    cfg := config.Load()

    database.Connect(cfg)

    handlers.SetJWTSecret(cfg.JWTSecret)

    r := routes.Setup(cfg.JWTSecret)

    addr := fmt.Sprintf(":%s", cfg.AppPort)
    log.Printf("Server %s berjalan di %s [env=%s]", cfg.AppName, addr, cfg.AppEnv)

    if err := http.ListenAndServe(addr, r); err != nil {
        log.Fatalf("Server error: %v", err)
    }
}
```

**Penjelasan baris per baris:**

- `cfg := config.Load()` — baca semua env vars dan bungkus jadi satu struct `Config`. Seluruh bagian aplikasi menerima konfigurasi dari struct ini, bukan baca env sendiri-sendiri.
- `database.Connect(cfg)` — buka koneksi ke PostgreSQL. Jika gagal, program berhenti (`log.Fatalf`). Tidak ada gunanya server berjalan tanpa database.
- `handlers.SetJWTSecret(cfg.JWTSecret)` — simpan JWT secret ke package `handlers` supaya bisa dipakai saat generate token.
- `routes.Setup(cfg.JWTSecret)` — buat router Chi dengan semua route terdaftar. `secret` diteruskan ke middleware auth.
- `http.ListenAndServe(addr, r)` — mulai server HTTP. `r` adalah router Chi yang sudah tahu semua URL yang valid.

---

### 5.2 `config/config.go`

```go
package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    AppEnv    string
    AppPort   string
    AppName   string
    DBHost    string
    DBPort    string
    DBName    string
    DBUser    string
    DBPass    string
    JWTSecret string
}

func Load() *Config {
    env := os.Getenv("APP_ENV")
    if env == "" {
        env = "development"
    }

    envFile := ".env." + env
    if err := godotenv.Load(envFile); err != nil {
        log.Printf("env file %s tidak ditemukan, pakai env OS", envFile)
    }

    return &Config{
        AppEnv:    getEnv("APP_ENV", "development"),
        AppPort:   getEnv("APP_PORT", "8080"),
        AppName:   getEnv("APP_NAME", "admin-api"),
        DBHost:    getEnv("DB_HOST", "localhost"),
        DBPort:    getEnv("DB_PORT", "5432"),
        DBName:    getEnv("DB_NAME", "vowtera_dev"),
        DBUser:    getEnv("DB_USER", "postgres"),
        DBPass:    getEnv("DB_PASSWORD", ""),
        JWTSecret: getEnv("JWT_SECRET", "changeme"),
    }
}

func getEnv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}
```

**Penjelasan:**

**Struct `Config`** — mewakili semua konfigurasi yang dibutuhkan aplikasi. Dengan struct ini, semua bagian kode yang butuh konfigurasi cukup terima satu parameter `*Config`, tidak perlu `os.Getenv` berulang-ulang.

**Fungsi `Load()`** bekerja dalam dua langkah:
1. Baca `APP_ENV` dari OS environment. Jika tidak ada, default ke `"development"`.
2. Load file `.env.development` (atau `.env.staging`, dll.) menggunakan `godotenv`. File ini akan di-inject ke OS environment, sehingga `os.Getenv` berikutnya bisa membacanya.

**Kenapa dua langkah?** Karena di server production, env vars biasanya di-set langsung di OS (bukan file), sehingga `godotenv.Load` akan gagal — itu normal dan tidak masalah (sudah di-handle dengan `log.Printf`, bukan `log.Fatalf`).

**Fungsi `getEnv(key, fallback)`** — helper sederhana: ambil env var, jika kosong pakai nilai default. Ini memastikan aplikasi bisa jalan meski env tidak lengkap (terutama saat development).

---

### 5.3 `database/database.go`

```go
package database

import (
    "fmt"
    "log"

    "admin-api/config"

    "github.com/jmoiron/sqlx"
    _ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sqlx.DB

func Connect(cfg *config.Config) {
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName,
    )

    db, err := sqlx.Connect("pgx", dsn)
    if err != nil {
        log.Fatalf("Gagal koneksi database: %v", err)
    }

    if err := migrate(db); err != nil {
        log.Fatalf("Gagal migrate: %v", err)
    }

    DB = db
    log.Println("Database terhubung dan migration selesai")
}

func migrate(db *sqlx.DB) error {
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id         SERIAL      PRIMARY KEY,
            name       TEXT        NOT NULL,
            email      TEXT        NOT NULL UNIQUE,
            password   TEXT        NOT NULL,
            role       TEXT        NOT NULL DEFAULT 'admin',
            created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
            updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
            deleted_at TIMESTAMPTZ
        );
        CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users (deleted_at);
    `)
    return err
}
```

**Penjelasan:**

**`var DB *sqlx.DB`** — variabel global yang menyimpan koneksi database. Bisa diakses dari package manapun dengan `database.DB`.

**`_ "github.com/jackc/pgx/v5/stdlib"`** — import dengan underscore artinya hanya jalankan `init()` dari package ini. Ini mendaftarkan driver `"pgx"` ke `database/sql` agar `sqlx.Connect("pgx", dsn)` bisa bekerja.

**DSN (Data Source Name)** — string koneksi standar PostgreSQL:
- `sslmode=disable` — untuk development lokal. Di Supabase, ganti ke `sslmode=require`.
- `TimeZone=Asia/Jakarta` — semua timestamp tersimpan dalam WIB.

**`migrate(db)`** — menjalankan SQL schema secara langsung. `CREATE TABLE IF NOT EXISTS` aman dijalankan berulang kali — jika tabel sudah ada, tidak ada yang berubah. Untuk menambah kolom baru di masa depan, tambahkan `ALTER TABLE` di dalam fungsi ini.

---

### 5.4 `models/user.go`

```go
package models

import (
    "time"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    ID        uint       `db:"id"         json:"id"`
    Name      string     `db:"name"       json:"name"`
    Email     string     `db:"email"      json:"email"`
    Password  string     `db:"password"   json:"-"`
    Role      string     `db:"role"       json:"role"`
    CreatedAt time.Time  `db:"created_at" json:"created_at"`
    UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
    DeletedAt *time.Time `db:"deleted_at" json:"-"`
}

func (u *User) HashPassword() error {
    hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashed)
    return nil
}

func (u *User) CheckPassword(plain string) bool {
    return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain)) == nil
}
```

**Penjelasan:**

**Struct tag** punya dua bagian yang bekerja independen:
- `db:"..."` — instruksi ke sqlx, memetakan nama kolom database ke field struct
- `json:"..."` — instruksi ke `encoding/json` saat encode/decode JSON

**`db` tag** harus sama persis dengan nama kolom di tabel PostgreSQL. Jika tidak ada tag `db`, sqlx akan coba cocokkan berdasarkan nama field (case-insensitive), tapi lebih baik eksplisit.

**`json:"-"`** pada field `Password` dan `DeletedAt` — tanda minus berarti field ini **tidak akan muncul** di JSON response. Penting untuk keamanan: password hash tidak boleh dikirim ke client.

**`DeletedAt *time.Time`** — pointer ke `time.Time` (bukan value). Kenapa pointer? Karena nilainya bisa `NULL` di database. Kalau pakai `time.Time` biasa (bukan pointer), Go tidak bisa membedakan antara "tidak ada nilai" dan "zero time". Dengan pointer, `nil` = kolom NULL di database.

Soft delete diimplementasi manual: query selalu tambahkan `AND deleted_at IS NULL`, dan hapus data cukup dengan `UPDATE users SET deleted_at = NOW() WHERE id = $1`.

**`HashPassword()`** — menggunakan bcrypt untuk hash password. Bcrypt adalah algoritma one-way: tidak bisa di-decrypt. `bcrypt.DefaultCost` = cost 10, artinya komputasi hash membutuhkan waktu ~100ms — cukup lambat untuk mencegah brute-force attack.

**`CheckPassword(plain string)`** — saat login, password yang diketik user di-hash ulang dan dibandingkan dengan hash yang tersimpan. Bcrypt sudah handle perbandingan ini dengan `CompareHashAndPassword`.

---

### 5.5 `middleware/auth.go`

```go
package middleware

import (
    "context"
    "net/http"
    "strings"

    "github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "userID"

func Auth(secret string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            header := r.Header.Get("Authorization")
            if !strings.HasPrefix(header, "Bearer ") {
                http.Error(w, "unauthorized", http.StatusUnauthorized)
                return
            }

            tokenStr := strings.TrimPrefix(header, "Bearer ")
            token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
                return []byte(secret), nil
            }, jwt.WithValidMethods([]string{"HS256"}))

            if err != nil || !token.Valid {
                http.Error(w, "unauthorized", http.StatusUnauthorized)
                return
            }

            claims, ok := token.Claims.(jwt.MapClaims)
            if !ok {
                http.Error(w, "unauthorized", http.StatusUnauthorized)
                return
            }

            ctx := context.WithValue(r.Context(), UserIDKey, claims["sub"])
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
```

**Penjelasan:**

**Pola middleware Go** — middleware adalah fungsi yang menerima `http.Handler` dan mengembalikan `http.Handler` baru. Ini memungkinkan middleware "membungkus" handler asli:

```
Auth(next) → kalau valid → next.ServeHTTP() → handler asli berjalan
           → kalau invalid → http.Error() → handler asli TIDAK berjalan
```

**`Auth(secret string) func(http.Handler) http.Handler`** — ini *middleware factory*: fungsi yang menghasilkan fungsi. Kenapa? Karena middleware perlu tahu `secret` JWT, tapi signature middleware standar Go tidak punya parameter tambahan. Dengan factory ini, `secret` di-capture oleh closure.

**`type contextKey string` dan `const UserIDKey`** — ini pola idiomatik Go untuk menghindari konflik key di context. Jika pakai string biasa (`"userID"`), package lain yang kebetulan pakai key sama bisa bentrok. Dengan custom type `contextKey`, key ini unik untuk package ini.

**`jwt.Parse()`** — verifikasi token. Di dalamnya:
- Decode base64 header dan payload
- Verifikasi signature menggunakan `secret`
- Cek apakah token sudah expired (`exp` claim)

**`jwt.WithValidMethods([]string{"HS256"})`** — ini penting untuk keamanan. Tanpa ini, attacker bisa mengirim token dengan algorithm `none` (tanpa signature) dan lolos verifikasi. Dengan whitelist ini, hanya token ber-algoritma HS256 yang diterima.

**`context.WithValue(r.Context(), UserIDKey, claims["sub"])`** — menyimpan user ID ke dalam context request. Context adalah cara Go meneruskan data antar fungsi dalam satu request lifecycle tanpa variabel global.

---

### 5.6 `handlers/auth.go`

```go
package handlers

import (
    "encoding/json"
    "net/http"
    "time"

    "admin-api/database"
    "admin-api/models"

    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret string

func SetJWTSecret(s string) { jwtSecret = s }

type loginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type loginResponse struct {
    Token string      `json:"token"`
    User  models.User `json:"user"`
}

func Login(w http.ResponseWriter, r *http.Request) {
    var req loginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "request tidak valid", http.StatusBadRequest)
        return
    }

    var user models.User
    err := database.DB.Get(&user,
        "SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL",
        req.Email,
    )
    if err != nil {
        http.Error(w, "email atau password salah", http.StatusUnauthorized)
        return
    }

    if !user.CheckPassword(req.Password) {
        http.Error(w, "email atau password salah", http.StatusUnauthorized)
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": user.ID,
        "exp": time.Now().Add(24 * time.Hour).Unix(),
    })

    tokenStr, err := token.SignedString([]byte(jwtSecret))
    if err != nil {
        http.Error(w, "gagal generate token", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(loginResponse{Token: tokenStr, User: user})
}

func Me(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value(middleware.UserIDKey)

    var user models.User
    err := database.DB.Get(&user,
        "SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL",
        userID,
    )
    if err != nil {
        http.Error(w, "user tidak ditemukan", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
```

**Penjelasan:**

**`loginRequest` dan `loginResponse`** — struct lokal yang hanya dipakai di file ini. Ini best practice: jangan gunakan struct model (`User`) langsung untuk decode input — karena user bisa mengirim field yang tidak diinginkan.

**Pola error di handler** — perhatikan setiap kemungkinan error selalu diikuti `return`. Ini penting agar kode di bawahnya tidak jalan. Tanpa `return`, Go akan terus eksekusi baris berikutnya meski sudah kirim response error.

**`database.DB.Get(&user, "SELECT ...", arg)`** — sqlx helper yang menjalankan query dan scan hasilnya langsung ke struct `user`. Menggunakan `$1` sebagai placeholder PostgreSQL (parameterized query — aman dari SQL injection). Jika tidak ada baris yang cocok, mengembalikan `sql.ErrNoRows`.

**Kenapa error login selalu "email atau password salah"?** — Sengaja. Jika pesan dibedakan ("email tidak ditemukan" vs "password salah"), attacker bisa tahu apakah suatu email terdaftar atau tidak (*user enumeration attack*).

**JWT Claims:**
- `"sub"` (subject) — standar JWT untuk menyimpan ID entitas (user ID)
- `"exp"` (expiration) — Unix timestamp kapan token kadaluarsa. `time.Now().Add(24 * time.Hour).Unix()` = 24 jam dari sekarang

**`token.SignedString([]byte(jwtSecret))`** — sign token dengan secret key menghasilkan string akhir yang dikirim ke client. Client menyimpan token ini (biasanya di `localStorage` atau cookie) dan mengirimnya kembali di header `Authorization`.

---

### 5.7 `routes/routes.go`

```go
package routes

import (
    "net/http"

    "admin-api/handlers"
    "admin-api/middleware"

    "github.com/go-chi/chi/v5"
    chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func Setup(secret string) *chi.Mux {
    r := chi.NewRouter()

    r.Use(chiMiddleware.Logger)
    r.Use(chiMiddleware.Recoverer)
    r.Use(chiMiddleware.RequestID)

    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("ok"))
    })

    r.Route("/api", func(r chi.Router) {
        r.Post("/login", handlers.Login)

        r.Group(func(r chi.Router) {
            r.Use(middleware.Auth(secret))
            r.Get("/me", handlers.Me)
        })
    })

    return r
}
```

**Penjelasan:**

**Global middleware** (`r.Use(...)`) — berlaku untuk semua route:
- `Logger` — log setiap request ke terminal: method, path, status code, durasi
- `Recoverer` — kalau ada `panic` di handler, tangkap otomatis dan kembalikan 500 (server tidak crash)
- `RequestID` — sisipkan ID unik di setiap request, berguna untuk tracing di log

**`r.Route("/api", ...)`** — membuat sub-router dengan prefix `/api`. Semua route di dalamnya otomatis punya prefix ini, jadi `/login` menjadi `/api/login`.

**`r.Group(...)`** — mengelompokkan route tanpa menambah prefix URL. Berguna untuk menerapkan middleware hanya ke sebagian route. Di sini, `middleware.Auth` hanya berlaku untuk route di dalam group ini.

**Peta route lengkap:**

```
GET  /health       → inline handler (tidak butuh auth)
POST /api/login    → handlers.Login (tidak butuh auth)
GET  /api/me       → middleware.Auth → handlers.Me (butuh JWT)
```

---

## 6. API Endpoints

### `GET /health`

Cek apakah server berjalan.

**Request:** tidak ada body, tidak ada header khusus

**Response:**
```
200 OK
ok
```

---

### `POST /api/login`

Login dan dapatkan JWT token.

**Request:**
```http
POST /api/login
Content-Type: application/json

{
  "email": "admin@vowtera.com",
  "password": "secret123"
}
```

**Response sukses (200):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "name": "Admin",
    "email": "admin@vowtera.com",
    "role": "admin",
    "created_at": "2026-05-11T10:00:00Z",
    "updated_at": "2026-05-11T10:00:00Z"
  }
}
```

**Response gagal (401):**
```
email atau password salah
```

---

### `GET /api/me`

Ambil data user yang sedang login. Membutuhkan JWT token.

**Request:**
```http
GET /api/me
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**Response sukses (200):**
```json
{
  "id": 1,
  "name": "Admin",
  "email": "admin@vowtera.com",
  "role": "admin",
  "created_at": "2026-05-11T10:00:00Z",
  "updated_at": "2026-05-11T10:00:00Z"
}
```

**Response gagal (401):**
```
unauthorized
```

---

## 7. Sistem Environment

Backend mendukung tiga environment. Pemilihan dilakukan via variabel `APP_ENV`.

### Cara Kerja

```
APP_ENV=staging → Load .env.staging
APP_ENV=production → Load .env.production
APP_ENV=(kosong) → Load .env.development
```

### Setup

```bash
cp .env.example .env.development
cp .env.example .env.staging
cp .env.example .env.production
# Edit masing-masing file sesuai environment
```

### Variabel Lengkap

| Variabel | Default | Keterangan |
|----------|---------|------------|
| `APP_ENV` | `development` | Nama environment aktif |
| `APP_PORT` | `8080` | Port server |
| `APP_NAME` | `admin-api` | Nama aplikasi (muncul di log) |
| `DB_HOST` | `localhost` | Host PostgreSQL |
| `DB_PORT` | `5432` | Port PostgreSQL |
| `DB_NAME` | `vowtera_dev` | Nama database |
| `DB_USER` | `postgres` | Username database |
| `DB_PASSWORD` | _(kosong)_ | Password database |
| `JWT_SECRET` | `changeme` | Secret untuk sign/verify JWT — **wajib diganti di production** |

### Menjalankan per Environment

```bash
# Development (default)
go run main.go

# Staging
APP_ENV=staging go run main.go

# Production
APP_ENV=production go run main.go
```

---

## 8. Konsep Penting

### JWT (JSON Web Token)

Token berbentuk tiga bagian dipisah titik:
```
eyJhbGciOiJIUzI1NiJ9   ← Header (algoritma)
.eyJzdWIiOjF9           ← Payload (data: user ID, expiry)
.SflKxwRJSMeKKF2QT4fw  ← Signature (verifikasi keaslian)
```

Server tidak menyimpan token di database. Cukup verifikasi signature — jika valid, server percaya isi payload-nya.

### Bcrypt

Hash satu arah untuk password. Tidak bisa di-decrypt. Proses login:
```
User ketik: "secret123"
    ↓ bcrypt.CompareHashAndPassword
Tersimpan:  "$2a$10$abc..." (hash)
    ↓
Cocok? → login berhasil
```

### Context Request

Cara meneruskan data dari middleware ke handler tanpa variabel global:
```
Middleware: context.WithValue(ctx, "userID", 1)
    ↓
Handler:    r.Context().Value("userID") → 1
```

### Soft Delete

GORM tidak benar-benar hapus baris. Hanya isi `deleted_at`:
```sql
-- Saat db.Delete(&user)
UPDATE users SET deleted_at = NOW() WHERE id = 1

-- Saat db.First(&user)
SELECT * FROM users WHERE id = 1 AND deleted_at IS NULL
```

---

## 9. Cara Menambah Fitur Baru

### Menambah Endpoint Baru

Contoh: tambah `GET /api/users` untuk list semua user.

**Langkah 1** — buat handler di `handlers/user.go`:
```go
package handlers

import (
    "encoding/json"
    "net/http"
    "admin-api/database"
    "admin-api/models"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
    var users []models.User
    database.DB.Find(&users)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}
```

**Langkah 2** — daftarkan di `routes/routes.go`:
```go
r.Group(func(r chi.Router) {
    r.Use(middleware.Auth(secret))
    r.Get("/me", handlers.Me)
    r.Get("/users", handlers.ListUsers) // tambahkan di sini
})
```

### Menambah Model Baru

Contoh: tambah model `Article`.

**Langkah 1** — buat `models/article.go`:
```go
package models

import (
    "time"
    "gorm.io/gorm"
)

type Article struct {
    ID        uint           `gorm:"primarykey" json:"id"`
    Title     string         `gorm:"not null" json:"title"`
    Body      string         `json:"body"`
    UserID    uint           `gorm:"not null" json:"user_id"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
```

**Langkah 2** — daftarkan di `database/database.go`:
```go
db.AutoMigrate(&models.User{}, &models.Article{}) // tambah &models.Article{}
```

GORM otomatis buat tabel `articles` saat server dijalankan berikutnya.
