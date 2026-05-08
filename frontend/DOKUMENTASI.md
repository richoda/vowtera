# Dokumentasi Kode — Admin Config

Dokumen ini menjelaskan setiap bagian kode yang dibangun dalam project ini: cara kerjanya, kenapa ditulis seperti itu, dan konsep Vue/Nuxt yang digunakan.

---

## Daftar Isi

1. [Alur Kerja Keseluruhan](#1-alur-kerja-keseluruhan)
2. [app.vue — Entry Point](#2-appvue--entry-point)
3. [useSidebar.ts — Shared State](#3-usesidebarts--shared-state)
4. [layouts/default.vue — Layout Utama](#4-layoutsdefaultvue--layout-utama)
5. [layouts/plain.vue — Layout Kosong](#5-layoutsplainvue--layout-kosong)
6. [SideBar.vue — Sidebar Collapsible](#6-sidebarvue--sidebar-collapsible)
7. [Header.vue — Header Dinamis](#7-headervue--header-dinamis)
8. [pages/login.vue — Halaman Login](#8-pagesloginvue--halaman-login)
9. [pages/overview.vue](#9-pagesoverviewvue)
10. [pages/about.vue](#10-pagesaboutvue)
11. [pages/contact.vue](#11-pagescontactvue)
12. [tailwind.config.js](#12-tailwindconfigjs)
13. [Konsep Vue & Nuxt yang Digunakan](#13-konsep-vue--nuxt-yang-digunakan)

---

## 1. Alur Kerja Keseluruhan

```
Browser membuka URL
       │
       ├─── /login ──────────────────────────────────────────────────────────►  layouts/plain.vue
       │                                                                         └── <slot /> = login.vue
       │
       └─── /overview, /about, /contact, / ────────────────────────────────►  layouts/default.vue
                                                                                 ├── <SideBar />
                                                                                 ├── <Header />
                                                                                 └── <main>
                                                                                       └── <slot /> = halaman aktif
                                                                                             │
                                                                              useSidebar.ts  │
                                                                           (state open/close │
                                                                            dibagi ke semua  │
                                                                            komponen)        │
```

---

## 2. `app.vue` — Entry Point

```vue
<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>
```

- `<NuxtLayout>` — membungkus seluruh aplikasi. Nuxt otomatis memilih layout berdasarkan `definePageMeta` di masing-masing halaman.
- `<NuxtPage />` — merender halaman sesuai URL, hasilnya masuk ke `<slot />` layout.
- File ini sengaja sesederhana mungkin — semua logic ada di layout dan halaman masing-masing.

---

## 3. `composables/useSidebar.ts` — Shared State

```ts
export const useSidebar = () => {
  const open = useState('sidebar', () => true)
  return { open }
}
```

| Kode | Artinya |
|------|---------|
| `useState('sidebar', () => true)` | State global bernama `'sidebar'`, nilai awal `true` (expanded). Bawaan Nuxt, bukan `ref` biasa. |
| `export const useSidebar` | Dipanggil dari komponen mana saja: `const { open } = useSidebar()` |
| `return { open }` | Mengekspos state agar bisa dibaca dan diubah dari luar |

**Kenapa `useState` bukan `ref`?**

`ref()` bersifat lokal — tiap komponen yang memanggil `ref(true)` mendapat salinan terpisah. `useState` dari Nuxt menjamin satu instance yang sama dibagikan ke semua komponen, sehingga toggle di `SideBar.vue` langsung dirasakan oleh `Header.vue` dan `layouts/default.vue`.

---

## 4. `layouts/default.vue` — Layout Utama

```vue
<script setup>
const { open } = useSidebar()
</script>

<template>
  <div>
    <SideBar />
    <Header />
    <main class="mt-20 p-6 transition-all duration-300" :class="open ? 'ml-64' : 'ml-24'">
      <slot />
    </main>
  </div>
</template>
```

- `<SideBar />` dan `<Header />` muncul di semua halaman tanpa perlu ditambahkan satu per satu.
- `mt-20` — mendorong konten ke bawah setinggi header (`h-20 = 80px`) yang bersifat `fixed`.
- `:class="open ? 'ml-64' : 'ml-24'"` — konten bergeser mengikuti lebar sidebar.
- `transition-all duration-300` — animasi margin sinkron dengan animasi lebar sidebar.

### Apa itu `<slot />`?

`<slot />` adalah lubang yang disediakan layout untuk diisi konten halaman dari luar.

```
layouts/default.vue (bingkai)
┌──────────────────────────────────────┐
│  SideBar  │  Header                  │
│           │──────────────────────────│
│           │                          │
│           │   [ <slot /> ]           │ ← diisi oleh NuxtPage
│           │                          │
└──────────────────────────────────────┘
```

Konten yang mengisi slot ditentukan di `app.vue`:
```vue
<NuxtLayout>
  <NuxtPage />  ← masuk ke <slot />
</NuxtLayout>
```

Saat user buka `/overview`:
1. Nuxt render `layouts/default.vue` → sidebar + header + main
2. `pages/overview.vue` dirender di dalam `<slot />`
3. Hasilnya: kerangka tetap, hanya isi main yang berganti

---

## 5. `layouts/plain.vue` — Layout Kosong

```vue
<template>
  <div>
    <slot />
  </div>
</template>
```

Layout minimal tanpa sidebar dan header. Digunakan oleh halaman yang tidak memerlukan navigasi, seperti halaman login.

Cara menggunakannya di `pages/*.vue`:
```js
definePageMeta({ layout: 'plain' })
```

Nuxt akan menggunakan `layouts/plain.vue` sebagai pembungkus halaman tersebut, bukan `layouts/default.vue`.

---

## 6. `SideBar.vue` — Sidebar Collapsible

```vue
<script setup>
const { open } = useSidebar()
</script>
```

State `open` diambil dari composable global — bukan `ref` lokal.

### Struktur Elemen

```
<header fixed top-0 left-0 h-screen>
  ├── Area Logo (h-20)
  │     ├── logo.png   — muncul saat open = true  (cross-fade)
  │     └── icon.png   — muncul saat open = false (cross-fade)
  ├── <nav> Menu Navigasi
  │     ├── Overview  → /overview
  │     ├── About     → /about
  │     └── Contact   → /contact
  └── Tombol Toggle Hamburger
```

### Animasi Logo (Cross-fade)

```vue
<img src="~/assets/img/logo.png"
  class="absolute transition-all duration-300"
  :class="open ? 'opacity-100 scale-100' : 'opacity-0 scale-90 pointer-events-none'"
/>
<img src="~/assets/img/icon.png"
  class="absolute transition-all duration-300"
  :class="!open ? 'opacity-100 scale-100' : 'opacity-0 scale-90 pointer-events-none'"
/>
```

Kedua gambar selalu ada di DOM. Yang berubah hanya `opacity` dan `scale` — tidak ada `v-if` agar animasi bisa berjalan mulus. `pointer-events-none` mencegah gambar tersembunyi diklik.

### Active State Menu

```vue
<NuxtLink to="/overview"
  class="... border-r-4 border-transparent hover:bg-[#B8873A]/10 hover:text-[#B8873A]"
  exact-active-class="!border-[#B8873A] !text-[#B8873A] bg-[#B8873A]/10">
```

| Class | Fungsi |
|-------|--------|
| `border-r-4 border-transparent` | Border kanan selalu ada (transparan) agar posisi konten tidak geser saat aktif |
| `exact-active-class` | Vue Router otomatis menerapkan class ini saat URL cocok persis |
| `!border-[#B8873A]` | `!` = `!important` — memaksa border amber menggantikan `border-transparent` |
| `bg-[#B8873A]/10` | Background amber 10% opacity — subtle highlight |

**`exact-active-class` vs `active-class`:**
- `active-class` — aktif jika URL *mengandung* path (misal `/overview/detail` juga highlight `/overview`)
- `exact-active-class` — aktif *hanya* jika URL *persis sama* → lebih tepat untuk sidebar

### Tombol Toggle Hamburger

```vue
<button @click="open = !open">
  <span :class="{ 'rotate-45 translate-y-2': !open }"></span>
  <span :class="{ 'opacity-0': !open }"></span>
  <span :class="{ '-rotate-45 -translate-y-2': !open }"></span>
</button>
```

Tiga `<span>` membentuk ≡. Saat `open = false`: baris atas rotasi +45°, tengah hilang, bawah rotasi -45° → membentuk ✕.

---

## 7. `Header.vue` — Header Dinamis

```vue
<script setup>
const { open } = useSidebar()
const route = useRoute()

const pages = {
  '/': { title: 'Home', icon: 'mdi:home' },
  '/overview': { title: 'Overview', icon: 'mdi:chart-bubble' },
  '/about': { title: 'About', icon: 'mdi:information' },
  '/contact': { title: 'Contact', icon: 'mdi:phone' },
}

const currentPage = computed(() =>
  pages[route.path] ?? { title: 'Dashboard', icon: 'mdi:view-dashboard' }
)
</script>
```

| Kode | Artinya |
|------|---------|
| `useRoute()` | Membaca route aktif termasuk `route.path` (URL saat ini) |
| `pages[route.path]` | Lookup judul & ikon berdasarkan URL |
| `??` | Nullish coalescing — fallback ke `'Dashboard'` jika URL tidak ada di object |
| `computed()` | Otomatis hitung ulang setiap kali `route.path` berubah |

### Posisi Mengikuti Sidebar

```vue
<header class="fixed top-0 right-0 ... transition-all duration-300"
  :class="open ? 'left-64' : 'left-24'">
```

`left-64` / `left-24` menggeser batas kiri header sinkron dengan lebar sidebar.

---

## 8. `pages/login.vue` — Halaman Login

```vue
<script setup>
definePageMeta({ layout: 'plain' })

const form = reactive({ email: '', password: '' })
const error = ref('')

function handleLogin() {
  if (form.email === 'admin@vowtera.com' && form.password === 'admin123') {
    navigateTo('/overview')
  } else {
    error.value = 'Email atau password salah'
  }
}
</script>
```

| Kode | Artinya |
|------|---------|
| `definePageMeta({ layout: 'plain' })` | Memberitahu Nuxt agar halaman ini pakai `layouts/plain.vue`, bukan `default.vue` |
| `reactive({ email, password })` | Object reaktif untuk form — akses langsung `form.email` tanpa `.value` |
| `ref('')` | State string untuk pesan error |
| `navigateTo('/overview')` | Redirect programatik bawaan Nuxt setelah login berhasil |

### Layout Dua Kolom

```
┌─────────────────────┬─────────────────────┐
│   side-design       │   side-login        │
│   (gradient bg +    │   (form login)      │
│    animasi bintang) │                     │
└─────────────────────┴─────────────────────┘
```

```vue
<div class="w-full h-screen flex">
  <div class="relative overflow-hidden w-1/2 h-full ...">  <!-- side-design -->
  <div class="w-1/2 h-full bg-[#F7F4EF] ...">             <!-- side-login -->
```

- `relative overflow-hidden` pada `side-design` — mengurung animasi agar tidak bocor keluar area
- `w-1/2` pada keduanya — membagi layar tepat 50:50

### Animasi Bintang (`clip-path`)

```css
.circles li {
  clip-path: polygon(
    50%  0%,    57%  33%,
    77%  23%,   67%  43%,
    100% 50%,   67%  57%,
    77%  77%,   57%  67%,
    50%  100%,  43%  67%,
    23%  77%,   33%  57%,
    0%   50%,   33%  43%,
    23%  23%,   43%  33%
  );
}
```

`clip-path: polygon()` memotong elemen menjadi bentuk bintang 8 sudut. 16 titik koordinat bergantian antara sudut luar (radius `50`) dan lekukan dalam (radius `18`). 4 sudut diagonal menggunakan radius `38` agar lebih pendek dari 4 sudut utama.

```
         (50,0)          ← panjang
    (23,23)   (77,23)    ← pendek
  (0,50)         (100,50) ← panjang
    (23,77)   (77,77)    ← pendek
         (50,100)        ← panjang
```

Elemen naik dengan `animation: animate` yang menggerakkan dari bawah ke atas sambil berputar 720°.

---

## 9. `pages/overview.vue`

- Data `stats` dan `activities` ditulis sebagai `const` biasa (bukan `ref`/`reactive`) karena nilainya tidak berubah.
- `v-for` merender kartu dan list dari array. `:key` wajib untuk tracking Vue.
- `:class="stat.bg"` — warna avatar diambil dari properti object, bukan ditulis ulang.

---

## 10. `pages/about.vue`

- Data statis `techStack` cukup `const` biasa.
- Grid tech stack dirender dengan `v-for`.

---

## 11. `pages/contact.vue`

```vue
const form = reactive({ name: '', email: '', message: '' })
const sent = ref(false)
```

| Kode | Artinya |
|------|---------|
| `reactive()` | Untuk object form — akses `form.name` langsung |
| `ref(false)` | Boolean untuk tampilkan/sembunyi pesan sukses |
| `v-model` | Two-way binding input ↔ state |
| `@submit.prevent` | Cegah reload halaman saat form disubmit |
| `setTimeout(..., 3000)` | Sembunyikan pesan sukses otomatis setelah 3 detik |

---

## 12. `tailwind.config.js`

```js
content: [
  './components/**/*.{vue,js,ts}',
  './layouts/**/*.vue',
  './pages/**/*.vue',
  './composables/**/*.{js,ts}',
  './app.vue',
],
```

`content` memberitahu Tailwind file mana yang harus di-scan untuk mencari class yang digunakan. Jika `content: []` (kosong), Tailwind tidak men-generate class apapun dan semua styling tidak akan tampil. **Wajib diisi** dan **wajib restart dev server** setelah mengubah file ini.

---

## 13. Konsep Vue & Nuxt yang Digunakan

### Vue 3

| Konsep | Digunakan di | Penjelasan |
|--------|-------------|------------|
| `ref()` | `contact.vue`, `login.vue` | State reaktif untuk nilai tunggal |
| `reactive()` | `contact.vue`, `login.vue` | State reaktif untuk object/form |
| `computed()` | `Header.vue` | Nilai turunan yang otomatis update |
| `<script setup>` | Semua file | Syntax modern Vue 3 |
| `v-for` | `overview`, `about`, `contact` | Render list dari array |
| `v-if` | `SideBar`, `contact`, `login` | Render kondisional |
| `v-model` | `contact.vue`, `login.vue` | Two-way binding input ↔ state |
| `:class` | Semua komponen | Class CSS dinamis sesuai kondisi |
| `@click`, `@submit.prevent` | `SideBar`, `contact`, `login` | Event listener |

### Nuxt 4

| Konsep | Digunakan di | Penjelasan |
|--------|-------------|------------|
| `useState()` | `useSidebar.ts` | State global lintas komponen, SSR-safe |
| `useRoute()` | `Header.vue` | Membaca URL/route aktif |
| `navigateTo()` | `login.vue` | Redirect programatik |
| `definePageMeta()` | `login.vue` | Konfigurasi per halaman (layout, middleware, dll.) |
| `<NuxtLayout>` | `app.vue` | Membungkus app dengan sistem layout |
| `<NuxtPage>` | `app.vue` | Merender halaman sesuai URL |
| `<NuxtLink>` | `SideBar.vue` | Navigasi internal tanpa reload |
| `exact-active-class` | `SideBar.vue` | Class otomatis saat route aktif persis |
| `layouts/default.vue` | Auto | Layout default semua halaman |
| `layouts/plain.vue` | `login.vue` | Layout alternatif tanpa sidebar |
| `composables/` | `useSidebar.ts` | Auto-import tanpa perlu `import` manual |
| `pages/` | Semua halaman | File-based routing — nama file = URL |
