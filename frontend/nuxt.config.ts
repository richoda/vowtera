// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: process.env.NUXT_PUBLIC_APP_ENV !== 'production' },
  modules: ['@nuxtjs/tailwindcss', '@nuxt/icon'],

  runtimeConfig: {
    public: {
      apiBaseUrl: process.env.NUXT_PUBLIC_API_BASE_URL || 'http://localhost:8080',
      appName: process.env.NUXT_PUBLIC_APP_NAME || 'Vowtera',
      appEnv: process.env.NUXT_PUBLIC_APP_ENV || 'development',
    },
  },
})
