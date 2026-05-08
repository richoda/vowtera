# Admin Config

Template admin panel berbasis **Nuxt 4** dengan sidebar collapsible, dibangun menggunakan Tailwind CSS dan Nuxt Icon.

---

## Tech Stack

| Teknologi | Versi | Keterangan |
|-----------|-------|------------|
| Nuxt | ^4.4.4 | Framework utama (SSR/SPA) |
| Vue | ^3.5.33 | UI framework |
| Vue Router | ^5.0.6 | Client-side routing |
| Tailwind CSS | ^6.14.0 | Utility-first CSS |
| Nuxt Icon | ^2.2.2 | Ikon berbasis Iconify |
| MDI Icons | ^1.2.3 | Paket ikon Material Design |

---

## Struktur Project

```
admin-config/
├── app.vue                    # Entry point — hanya <NuxtLayout><NuxtPage />
├── layouts/
│   ├── default.vue            # Layout utama (SideBar + Header + <slot />)
│   └── plain.vue              # Layout kosong untuk halaman tanpa sidebar (login)
├── components/
│   ├── SideBar.vue            # Sidebar collapsible, fixed kiri, active state
│   └── Header.vue             # Header fixed atas, judul & ikon dinamis
├── composables/
│   └── useSidebar.ts          # Shared state open/close sidebar
├── pages/
│   ├── index.vue              # Home
│   ├── login.vue              # Halaman login (layout plain)
│   ├── overview.vue           # Dashboard stats & aktivitas
│   ├── about.vue              # Info aplikasi, tech stack, tim
│   └── contact.vue            # Form pesan & info kontak
├── assets/
│   ├── css/global.css
│   └── img/
│       ├── logo.png           # Logo penuh (sidebar expanded)
│       ├── logo1.png          # Logo di halaman login
│       └── icon.png           # Ikon kecil (sidebar collapsed)
├── public/
├── nuxt.config.ts
├── tailwind.config.js         # content paths wajib diisi agar class ter-generate
└── package.json
```

---

## Design System

### Warna

| Token | Hex | Digunakan untuk |
|-------|-----|-----------------|
| Navy | `#0F1C33` | Teks utama, background gelap, hover button |
| Amber | `#B8873A` | Border aktif sidebar, tombol primary, aksen |
| Krem | `#F7F4EF` | Background sidebar, input, halaman login |
| Abu | `#B8ADA1` | Border, placeholder, teks sekunder |
| Putih | `#FFFFFF` | Background card, header |

### Ikon

Semua ikon menggunakan **MDI via Nuxt Icon**:
```html
<Icon name="mdi:nama-ikon" class="text-xl" />
```
Referensi ikon: [Iconify MDI](https://icon-sets.iconify.design/mdi/)

### Ukuran Sidebar

| State | Lebar | Class |
|-------|-------|-------|
| Expanded | 256px | `w-64` / `ml-64` / `left-64` |
| Collapsed | 96px | `w-24` / `ml-24` / `left-24` |

---

## Komponen & File Utama

### `useSidebar.ts`
Composable global untuk state buka/tutup sidebar. Dipakai di `SideBar.vue`, `Header.vue`, dan `layouts/default.vue`.
```ts
const { open } = useSidebar()
```

### `SideBar.vue`
- `fixed top-0 left-0 h-screen` — menempel di kiri viewport
- Animasi logo cross-fade antara `logo.png` dan `icon.png`
- Active state: border kanan amber + background tipis via `exact-active-class`
- Hover: background `10%` opacity amber + teks amber
- Tombol hamburger di bawah dengan animasi CSS ke ✕

### `Header.vue`
- `fixed top-0 right-0` dengan `:class="open ? 'left-64' : 'left-24'"`
- Judul & ikon halaman berubah otomatis via `useRoute()` + `computed()`
- Berisi: judul halaman, search bar, notifikasi, user avatar

### `layouts/default.vue`
- Berlaku otomatis untuk semua halaman kecuali yang mendefinisikan layout lain
- `<main>` pakai `:class="open ? 'ml-64' : 'ml-24'"` + `mt-20`

### `layouts/plain.vue`
- Layout kosong tanpa sidebar dan header
- Dipakai oleh `pages/login.vue` via `definePageMeta({ layout: 'plain' })`

### `pages/login.vue`
- Split layout: kiri (animasi bintang + logo), kanan (form login)
- Animasi bintang menggunakan `clip-path: polygon()` CSS
- Gradient background `#0F1C33` → `#B8ADA1`
- Login tanpa backend — kredensial di-hardcode untuk prototype

---

## Setup

```bash
npm install
npm run dev      # http://localhost:3000
npm run build
npm run preview
```

---

## Menambah Halaman Baru

1. Buat `pages/nama-halaman.vue`
2. Tambahkan menu di `SideBar.vue`:

```html
<li>
  <NuxtLink to="/nama-halaman"
    class="flex items-center gap-3 py-4 text-[#0F1C33] border-r-4 border-transparent hover:bg-[#B8873A]/10 hover:text-[#B8873A] transition-colors duration-200"
    :class="open ? 'px-6' : 'justify-center px-4'"
    exact-active-class="!border-[#B8873A] !text-[#B8873A] bg-[#B8873A]/10">
    <Icon name="mdi:nama-ikon" class="text-xl shrink-0" />
    <span v-if="open">Label Menu</span>
  </NuxtLink>
</li>
```

3. Tambahkan entry di `Header.vue` object `pages`:
```js
'/nama-halaman': { title: 'Label', icon: 'mdi:nama-ikon' },
```

---

## Do & Don't

### Warna & Styling

| ✅ Do | ❌ Don't |
|-------|---------|
| Gunakan token warna `#0F1C33`, `#B8873A`, `#F7F4EF`, `#B8ADA1` | Jangan pakai warna Tailwind generik (`red-500`, `blue-600`, `slate-600`, dll.) |
| Gunakan `transition-colors duration-200` untuk hover interaktif | Jangan animasi tanpa `transition` |
| Card: `bg-white rounded-xl border border-[#B8ADA1]` | Jangan pakai `shadow-*` sebagai pengganti border |
| Pastikan `tailwind.config.js` punya `content` yang scan semua file `.vue` | Jangan biarkan `content: []` kosong |

### Komponen & State

| ✅ Do | ❌ Don't |
|-------|---------|
| Gunakan `useSidebar()` untuk membaca/mengubah state sidebar | Jangan buat `ref(true)` lokal untuk state sidebar |
| Gunakan `<script setup>` di semua komponen | Jangan pakai Options API |
| Gunakan `reactive()` untuk form, `ref()` untuk nilai tunggal | Jangan campur tanpa alasan jelas |

### Layout & Halaman

| ✅ Do | ❌ Don't |
|-------|---------|
| Letakkan semua halaman di folder `pages/` | Jangan render halaman di luar sistem routing Nuxt |
| Gunakan `definePageMeta({ layout: 'plain' })` untuk halaman tanpa sidebar | Jangan taruh `<SideBar />` langsung di dalam `pages/` |
| Tambahkan entry di object `pages` di `Header.vue` saat buat halaman baru | Jangan biarkan halaman baru tanpa judul di header |
| Gunakan `<NuxtLink>` untuk navigasi internal | Jangan pakai `<a href>` antar halaman |
| Gunakan `navigateTo()` untuk redirect programatik | Jangan pakai `window.location.href` |

### Ikon & Gambar

| ✅ Do | ❌ Don't |
|-------|---------|
| Gunakan `<Icon name="mdi:..." />` | Jangan pakai SVG inline atau `<img>` untuk ikon |
| Simpan gambar di `assets/img/` dan akses dengan `~/assets/img/` | Jangan gunakan URL gambar eksternal untuk aset lokal |
| Gunakan animasi opacity + scale untuk pergantian gambar | Jangan pakai `v-if`/`v-else` jika pergantian perlu animasi |
