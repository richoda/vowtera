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

<template>
  <div class="w-full h-screen flex">

    <!-- Kiri: Side Design -->
    <div class="side-design relative overflow-hidden w-1/2 h-full flex items-center justify-center">
      <div class="area">
        <ul class="circles">
          <li></li><li></li><li></li><li></li><li></li>
          <li></li><li></li><li></li><li></li><li></li>
        </ul>
      </div>
      <img src="~/assets/img/logo1.png" alt="Logo" class="relative z-10 h-[450px] w-auto" />
    </div>

    <!-- Kanan: Side Login -->
    <div class="w-1/2 h-full bg-[#F7F4EF] flex items-center justify-center">
      <div class="bg-white rounded-xl border border-[#B8ADA1] p-8 w-full max-w-sm mx-8">
        <h1 class="text-2xl font-bold text-[#0F1C33] mb-1">Selamat Datang</h1>
        <p class="text-sm text-[#B8ADA1] mb-6">Masuk ke akun Anda</p>

        <form @submit.prevent="handleLogin" class="flex flex-col gap-4">
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
            <label class="block text-sm font-medium text-[#0F1C33] mb-1">Password</label>
            <input
              v-model="form.password"
              type="password"
              placeholder="Masukkan password"
              class="w-full border border-[#B8ADA1] rounded-lg px-4 py-2.5 text-sm text-[#0F1C33] placeholder-[#B8ADA1] outline-none focus:border-[#B8873A] transition-colors"
            />
          </div>

          <p v-if="error" class="text-sm text-red-500 flex items-center gap-1">
            <Icon name="mdi:alert-circle-outline" />
            {{ error }}
          </p>

          <button
            type="submit"
            class="w-full bg-[#B8873A] text-white rounded-lg px-6 py-2.5 text-sm font-semibold hover:bg-[#0F1C33] transition-colors duration-200"
          >
            Masuk
          </button>
        </form>
      </div>
    </div>

  </div>
</template>

<style scoped>
.area {
  position: absolute;
  inset: 0;
  background: linear-gradient(to left, #0F1C33, #B8ADA1);
  z-index: 0;
}

.circles {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  overflow: hidden;
  padding: 0;
  margin: 0;
  list-style: none;
}

.circles li {
  position: absolute;
  display: block;
  width: 20px;
  height: 20px;
  background: rgba(184, 173, 161, 0.685);
  animation: animate 25s linear infinite;
  bottom: -150px;
  clip-path: polygon(
    50%  0%,    57%  33%,   /* atas (panjang)        → dalam */
    77%  23%,   67%  43%,   /* kanan atas (pendek)   → dalam */
    100% 50%,   67%  57%,   /* kanan (panjang)       → dalam */
    77%  77%,   57%  67%,   /* kanan bawah (pendek)  → dalam */
    50%  100%,  43%  67%,   /* bawah (panjang)       → dalam */
    23%  77%,   33%  57%,   /* kiri bawah (pendek)   → dalam */
    0%   50%,   33%  43%,   /* kiri (panjang)        → dalam */
    23%  23%,   43%  33%    /* kiri atas (pendek)    → dalam */
  );
}

.circles li:nth-child(1)  { left: 25%; width: 80px;  height: 80px;  animation-delay: 0s; }
.circles li:nth-child(2)  { left: 10%; width: 20px;  height: 20px;  animation-delay: 2s;  animation-duration: 12s; }
.circles li:nth-child(3)  { left: 70%; width: 20px;  height: 20px;  animation-delay: 4s; }
.circles li:nth-child(4)  { left: 40%; width: 60px;  height: 60px;  animation-delay: 0s;  animation-duration: 18s; }
.circles li:nth-child(5)  { left: 65%; width: 20px;  height: 20px;  animation-delay: 0s; }
.circles li:nth-child(6)  { left: 75%; width: 110px; height: 110px; animation-delay: 3s; }
.circles li:nth-child(7)  { left: 35%; width: 150px; height: 150px; animation-delay: 7s; }
.circles li:nth-child(8)  { left: 50%; width: 25px;  height: 25px;  animation-delay: 15s; animation-duration: 45s; }
.circles li:nth-child(9)  { left: 20%; width: 15px;  height: 15px;  animation-delay: 2s;  animation-duration: 35s; }
.circles li:nth-child(10) { left: 85%; width: 150px; height: 150px; animation-delay: 0s;  animation-duration: 11s; }

@keyframes animate {
  0%   { transform: translateY(0) rotate(0deg);       opacity: 1; border-radius: 0; }
  100% { transform: translateY(-1000px) rotate(720deg); opacity: 0; border-radius: 50%; }
}
</style>
