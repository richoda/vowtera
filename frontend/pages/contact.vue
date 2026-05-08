<template>
  <div>
    <h1 class="text-2xl font-bold text-[#0F1C33] mb-6">Contact</h1>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Form -->
      <div class="bg-white rounded-xl border border-[#B8ADA1] p-6">
        <h2 class="text-lg font-semibold text-[#0F1C33] mb-5">Kirim Pesan</h2>
        <form @submit.prevent="handleSubmit" class="flex flex-col gap-4">
          <div>
            <label class="block text-sm font-medium text-[#0F1C33] mb-1">Nama</label>
            <input
              v-model="form.name"
              type="text"
              placeholder="Masukkan nama"
              class="w-full border border-[#B8ADA1] rounded-lg px-4 py-2.5 text-sm text-[#0F1C33] placeholder-[#B8ADA1] outline-none focus:border-[#B8873A] transition-colors"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-[#0F1C33] mb-1">Email</label>
            <input
              v-model="form.email"
              type="email"
              placeholder="Masukkan email"
              class="w-full border border-[#B8ADA1] rounded-lg px-4 py-2.5 text-sm text-[#0F1C33] placeholder-[#B8ADA1] outline-none focus:border-[#B8873A] transition-colors"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-[#0F1C33] mb-1">Pesan</label>
            <textarea
              v-model="form.message"
              rows="4"
              placeholder="Tulis pesan..."
              class="w-full border border-[#B8ADA1] rounded-lg px-4 py-2.5 text-sm text-[#0F1C33] placeholder-[#B8ADA1] outline-none focus:border-[#B8873A] transition-colors resize-none"
            />
          </div>
          <button
            type="submit"
            class="bg-[#B8873A] text-white rounded-lg px-6 py-2.5 text-sm font-semibold hover:bg-[#0F1C33] transition-colors duration-200 flex items-center justify-center gap-2"
          >
            <Icon name="mdi:send" class="text-base" />
            Kirim Pesan
          </button>

          <p v-if="sent" class="text-sm text-green-600 flex items-center gap-1">
            <Icon name="mdi:check-circle" />
            Pesan berhasil dikirim!
          </p>
        </form>
      </div>

      <!-- Info Kontak -->
      <div class="flex flex-col gap-4">
        <div v-for="info in contactInfo" :key="info.label"
          class="bg-white rounded-xl border border-[#B8ADA1] p-5 flex items-center gap-4">
          <div class="w-11 h-11 rounded-full bg-[#F7F4EF] flex items-center justify-center shrink-0">
            <Icon :name="info.icon" class="text-xl text-[#B8873A]" />
          </div>
          <div>
            <p class="text-xs text-[#B8ADA1] mb-0.5">{{ info.label }}</p>
            <p class="text-sm font-medium text-[#0F1C33]">{{ info.value }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const form = reactive({ name: '', email: '', message: '' })
const sent = ref(false)

const contactInfo = [
  { label: 'Email', value: 'admin@example.com', icon: 'mdi:email-outline' },
  { label: 'Telepon', value: '+62 812 3456 7890', icon: 'mdi:phone-outline' },
  { label: 'Alamat', value: 'Jakarta, Indonesia', icon: 'mdi:map-marker-outline' },
  { label: 'Jam Kerja', value: 'Senin – Jumat, 08.00 – 17.00', icon: 'mdi:clock-outline' },
]

function handleSubmit() {
  if (!form.name || !form.email || !form.message) return
  sent.value = true
  form.name = ''
  form.email = ''
  form.message = ''
  setTimeout(() => sent.value = false, 3000)
}
</script>
