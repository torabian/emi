import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  base: "/emi/playground/",
  // shiki (used by codemirror-json-schema hovers) reads `process.env[...]`,
  // which doesn't exist in the browser.
  // monaco-yaml's worker runs unbundled in dev and imports CJS/UMD-only packages;
  // pre-bundle them so `module`/`exports` exist in the browser.
  optimizeDeps: {
    include: [
      "monaco-yaml/yaml.worker",
      "vscode-languageserver-types",
      "vscode-languageserver-textdocument",
      "vscode-uri",
      "yaml",
      "path-browserify",
      "prettier/standalone",
      "prettier/plugins/yaml",
    ],
  },
  define: {
    "process.env": {},
  },
});
