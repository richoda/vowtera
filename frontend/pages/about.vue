<template>
  <div>
    <h1 class="text-2xl font-bold text-[#0F1C33] mb-6">About</h1>

    <!-- Info Aplikasi -->
    <div class="bg-white rounded-xl border border-[#B8ADA1] p-6 mb-6">
      <div class="flex items-center gap-4 mb-4">
        <div class="w-14 h-14 rounded-full bg-[#B8873A] flex items-center justify-center">
          <Icon name="mdi:information" class="text-3xl text-white" />
        </div>
        <div>
          <h2 class="text-xl font-bold text-[#0F1C33]">Admin Config</h2>
          <p class="text-sm text-[#B8ADA1]">Versi 1.0.0</p>
        </div>
      </div>
      <p class="text-sm text-[#0F1C33] leading-relaxed">
        Admin Config adalah panel administrasi berbasis Nuxt 4 dan Tailwind CSS.
        Dibangun untuk memudahkan pengelolaan data, monitoring aktivitas, dan konfigurasi sistem secara terpusat.
      </p>
    </div>

    <!-- Tech Stack -->
    <div class="bg-white rounded-xl border border-[#B8ADA1] p-6 mb-6">
      <h2 class="text-lg font-semibold text-[#0F1C33] mb-4">Tech Stack</h2>
      <div class="grid grid-cols-2 sm:grid-cols-3 gap-3">
        <div v-for="tech in techStack" :key="tech.name"
          class="flex items-center gap-2 bg-[#F7F4EF] rounded-lg px-4 py-3">
          <Icon :name="tech.icon" class="text-xl text-[#B8873A] shrink-0" />
          <span class="text-sm font-medium text-[#0F1C33]">{{ tech.name }}</span>
        </div>
      </div>
    </div>

    <!-- Tim -->
    <div class="bg-white rounded-xl border border-[#B8ADA1] p-6">
      <h2 class="text-lg font-semibold text-[#0F1C33] mb-4">Tim Pengembang</h2>
      <div class="flex items-center gap-4">
        <div class="w-12 h-12 rounded-full bg-[#0F1C33] flex items-center justify-center text-white font-bold text-lg">
          R
        </div>
        <div>
          <p class="font-semibold text-[#0F1C33]">richoda</p>
          <p class="text-sm text-[#B8ADA1]">Lead Developer</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const techStack = [
  { name: 'Nuxt 4', icon: 'mdi:nuxt' },
  { name: 'Vue 3', icon: 'mdi:vuejs' },
  { name: 'Tailwind CSS', icon: 'mdi:tailwind' },
  { name: 'TypeScript', icon: 'mdi:language-typescript' },
  { name: 'Nuxt Icon', icon: 'mdi:emoticon-outline' },
  { name: 'Iconify MDI', icon: 'mdi:shape' },
]
</script>
