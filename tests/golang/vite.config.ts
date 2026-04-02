/// vite.config.ts
import path from "path";

import WebSocket from "ws";

(globalThis as any).WebSocket = WebSocket;

import { defineConfig as defineVitestConfig } from "vitest/config";

export default defineVitestConfig({
  test: {
    globals: true, // so you can use "describe", "it", etc. without imports
    environment: "jsdom", // or 'jsdom' if testing browser code
  },
});
