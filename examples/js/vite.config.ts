/// vite.config.ts
import path from "path";

import WebSocket from "ws";

(globalThis as any).WebSocket = WebSocket;

import { defineConfig as defineVitestConfig } from "vitest/config";

export default defineVitestConfig({
  server: {
    watch: {
      // The cases write their generated SDK code (User.ts, *.output.ts/js) back
      // into the cases/ dir so they can be dynamically imported. Those files are
      // part of the module graph, so in `vitest watch` every write would retrigger
      // the run, which writes again — an endless refresh loop. Ignore the
      // generated artifacts here so the watcher never reacts to them. (These globs
      // are merged with Vite's own defaults, e.g. node_modules/.git.)
      ignored: [
        "**/cases/User.ts",
        "**/cases/*.output.ts",
        "**/cases/*.output.js",
      ],
    },
  },
  test: {
    globals: true, // so you can use "describe", "it", etc. without imports
    environment: "jsdom", // or 'jsdom' if testing browser code
  },
});
