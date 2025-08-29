/// vite.config.ts
import { defineConfig } from 'vite'
import { defineConfig as defineVitestConfig } from 'vitest/config'

export default defineVitestConfig({
  test: {
    globals: true,   // so you can use "describe", "it", etc. without imports
    environment: 'node', // or 'jsdom' if testing browser code
  },
})

