<script setup>
const { open } = useSidebar()
</script>

<template>
  <header class="bg-[#F7F4EF] h-screen fixed top-0 left-0 flex flex-col transition-all duration-300 overflow-hidden"
    :class="open ? 'w-64' : 'w-24'">
    <!-- Top Logo -->
    <div class="relative flex items-center justify-center border-b border-[#B8ADA1] px-4 py-5 h-20">
      <img src="~/assets/img/logo.png" alt="Logo" class="absolute h-40 w-auto transition-all duration-300"
        :class="open ? 'opacity-100 scale-100' : 'opacity-0 scale-90 pointer-events-none'" />
      <img src="~/assets/img/icon.png" alt="Icon" class="absolute h-20 w-auto transition-all duration-300"
        :class="!open ? 'opacity-100 scale-100' : 'opacity-0 scale-90 pointer-events-none'" />
    </div>

    <!-- Navigation Menu -->
    <nav class="flex-1">
      <ul class="flex flex-col">
        <li>
          <NuxtLink to="/overview"
            class="flex items-center gap-3 py-4 text-[#0F1C33] border-r-4 border-transparent hover:bg-[#B8873A]/10 hover:text-[#B8873A] transition-colors duration-200"
            :class="open ? 'px-6' : 'justify-center px-4'"
            exact-active-class="!border-[#B8873A] !text-[#B8873A] bg-[#B8873A]/10">
            <Icon name="mdi:chart-bubble" class="text-xl shrink-0" />
            <span v-if="open">Overview</span>
          </NuxtLink>
        </li>
        <li>
          <NuxtLink to="/about"
            class="flex items-center gap-3 py-4 text-[#0F1C33] border-r-4 border-transparent hover:bg-[#B8873A]/10 hover:text-[#B8873A] transition-colors duration-200"
            :class="open ? 'px-6' : 'justify-center px-4'"
            exact-active-class="!border-[#B8873A] !text-[#B8873A] bg-[#B8873A]/10">
            <Icon name="mdi:information" class="text-xl shrink-0" />
            <span v-if="open">About</span>
          </NuxtLink>
        </li>
        <li>
          <NuxtLink to="/contact"
            class="flex items-center gap-3 py-4 text-[#0F1C33] border-r-4 border-transparent hover:bg-[#B8873A]/10 hover:text-[#B8873A] transition-colors duration-200"
            :class="open ? 'px-6' : 'justify-center px-4'"
            exact-active-class="!border-[#B8873A] !text-[#B8873A] bg-[#B8873A]/10">
            <Icon name="mdi:phone" class="text-xl shrink-0" />
            <span v-if="open">Contact</span>
          </NuxtLink>
        </li>
      </ul>
    </nav>

    <!-- Toggle Button (Bottom) -->
    <div class="border-t border-[#B8ADA1] px-4 py-4 flex" :class="open ? 'justify-end' : 'justify-center'">
      <button @click="open = !open" class="relative w-7 h-5 flex flex-col justify-between">
        <span class="block h-0.5 w-7 bg-[#0F1C33] transition duration-500"
          :class="{ 'rotate-45 translate-y-2': !open }"></span>
        <span class="block h-0.5 w-5 bg-[#0F1C33] transition duration-500" :class="{ 'opacity-0': !open }"></span>
        <span class="block h-0.5 w-7 bg-[#0F1C33] transition duration-500"
          :class="{ '-rotate-45 -translate-y-2': !open }"></span>
      </button>
    </div>
  </header>
</template>
