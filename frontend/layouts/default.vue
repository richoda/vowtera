<script setup>
const { open } = useSidebar()
</script>

<template>
  <div>
    <SideBar />
    <Header />
    <main
      class="mt-20 p-6 transition-all duration-300"
      :class="open ? 'ml-64' : 'ml-24'"
    >
      <slot />
    </main>
  </div>
</template>
