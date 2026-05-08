<template>
  <div>
    <h1 class="text-2xl font-bold text-[#0F1C33] mb-6">Overview</h1>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
      <div v-for="stat in stats" :key="stat.label"
        class="bg-white rounded-xl p-5 border border-[#B8ADA1] flex items-center gap-4">
        <div class="w-12 h-12 rounded-full flex items-center justify-center" :class="stat.bg">
          <Icon :name="stat.icon" class="text-2xl text-white" />
        </div>
        <div>
          <p class="text-sm text-[#B8ADA1]">{{ stat.label }}</p>
          <p class="text-xl font-bold text-[#0F1C33]">{{ stat.value }}</p>
        </div>
      </div>
    </div>

    <!-- Recent Activity -->
    <div class="bg-white rounded-xl border border-[#B8ADA1] p-6">
      <h2 class="text-lg font-semibold text-[#0F1C33] mb-4">Aktivitas Terbaru</h2>
      <ul class="divide-y divide-[#F7F4EF]">
        <li v-for="item in activities" :key="item.id" class="py-3 flex items-center gap-3">
          <Icon name="mdi:circle-small" class="text-[#B8873A] text-2xl shrink-0" />
          <div>
            <p class="text-sm text-[#0F1C33] font-medium">{{ item.title }}</p>
            <p class="text-xs text-[#B8ADA1]">{{ item.time }}</p>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
const stats = [
  { label: 'Total Pengguna', value: '1,284', icon: 'mdi:account-group', bg: 'bg-[#B8873A]' },
  { label: 'Pendapatan', value: 'Rp 48 jt', icon: 'mdi:cash-multiple', bg: 'bg-[#0F1C33]' },
  { label: 'Transaksi', value: '320', icon: 'mdi:swap-horizontal', bg: 'bg-[#B8873A]' },
  { label: 'Laporan Baru', value: '7', icon: 'mdi:file-chart', bg: 'bg-[#0F1C33]' },
]

const activities = [
  { id: 1, title: 'Pengguna baru mendaftar: Budi Santoso', time: '5 menit lalu' },
  { id: 2, title: 'Transaksi #4821 berhasil diproses', time: '23 menit lalu' },
  { id: 3, title: 'Laporan bulanan telah diunggah', time: '1 jam lalu' },
  { id: 4, title: 'Pengaturan sistem diperbarui', time: '3 jam lalu' },
  { id: 5, title: 'Backup database selesai', time: 'Kemarin, 22:00' },
]
</script>
