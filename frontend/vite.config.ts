import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  build: {
    outDir: '../backend/web',
    emptyOutDir: false,
  },
  server: {
    port: 7702,
    proxy: {
      '/api': {
        target: 'http://localhost:7701',
        changeOrigin: true,
      },
    },
  },
})
