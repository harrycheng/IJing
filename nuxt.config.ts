export default defineNuxtConfig({
  sourcemap: {
    server: true,
    client: true
  },
  nitro: {
    plugins: ['~/server/plugins/ijing.server.ts'],
    publicAssets: [
      {
        baseURL: '/public',
        dir: 'public'
      }
    ]
  }
})