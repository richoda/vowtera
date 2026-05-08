<script setup>
const { open } = useSidebar()
const route = useRoute()

const pages = {
  '/': { title: 'Home', icon: 'mdi:home' },
  '/overview': { title: 'Overview', icon: 'mdi:chart-bubble' },
  '/about': { title: 'About', icon: 'mdi:information' },
  '/contact': { title: 'Contact', icon: 'mdi:phone' },
}

const currentPage = computed(() => pages[route.path] ?? { title: 'Dashboard', icon: 'mdi:view-dashboard' })
</script>

<template>
  <header
    class="fixed top-0 right-0 z-10 bg-white border-b border-[#B8ADA1] h-20 flex items-center justify-between px-6 transition-all duration-300"
    :class="open ? 'left-64' : 'left-24'"
  >
    <!-- Kiri: Judul halaman -->
    <div class="flex items-center gap-2 text-[#0F1C33]">
      <Icon :name="currentPage.icon" class="text-[#B8873A] text-xl" />
      <span class="font-semibold text-lg">{{ currentPage.title }}</span>
    </div>

    <!-- Kanan: Search + Notifikasi + User -->
    <div class="flex items-center gap-4">
      <!-- Search -->
      <div class="flex items-center gap-2 bg-[#F7F4EF] border border-[#B8ADA1] rounded-lg px-3 py-1.5">
        <Icon name="mdi:magnify" class="text-[#B8ADA1] text-lg" />
        <input
          type="text"
          placeholder="Cari..."
          class="bg-transparent text-sm text-[#0F1C33] placeholder-[#B8ADA1] outline-none w-40"
        />
      </div>

      <!-- Notifikasi -->
      <button class="relative p-2 rounded-lg hover:bg-[#F7F4EF] text-[#0F1C33] transition-colors duration-200">
        <Icon name="mdi:bell-outline" class="text-xl" />
        <span class="absolute top-1.5 right-1.5 w-2 h-2 rounded-full bg-[#B8873A]"></span>
      </button>

      <!-- User -->
      <div class="flex items-center gap-2 cursor-pointer hover:bg-[#F7F4EF] rounded-lg px-2 py-1.5 transition-colors duration-200">
        <div class="w-8 h-8 rounded-full bg-[#B8873A] flex items-center justify-center text-white text-sm font-semibold">
          A
        </div>
        <span class="text-sm font-medium text-[#0F1C33]">Admin</span>
        <Icon name="mdi:chevron-down" class="text-[#B8ADA1] text-lg" />
      </div>
    </div>
  </header>
</template>
