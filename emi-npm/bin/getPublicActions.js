import "./wasm_exec.js";

import fs from "fs";
import { fileURLToPath } from "url";
import path from "path";

const __dirname = path.dirname(fileURLToPath(import.meta.url));

const createInstance = async () => {
  const go = new Go(); // global from wasm_exec.js
  const wasmPath = path.join(__dirname, "emi-compiler.wasm");
  const wasmBuffer = fs.readFileSync(wasmPath);

  const { instance } = await WebAssembly.instantiate(
    wasmBuffer,
    go.importObject
  );
  go.run(instance);
};

await createInstance();

const { TextActions, FileActions } = globalThis.getPublicActions();

export { TextActions, FileActions, createInstance };
