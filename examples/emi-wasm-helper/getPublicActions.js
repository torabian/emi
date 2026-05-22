import "./wasm_exec.js";

import fs, { mkdirSync, writeFileSync } from "fs";
import { fileURLToPath } from "url";
import path, { dirname, join } from "path";
import prettier from "prettier";

const __dirname = path.dirname(fileURLToPath(import.meta.url));

export const fileWriter = async (files, output) => {
  for (const f of files) {
    const fullPath = join(output, f.Location, f.Name + f.Extension);
    mkdirSync(dirname(fullPath), { recursive: true });

    let result = f.ActualScript;
    try {
      result = await prettier.format(f.ActualScript, {
        parser: "typescript",
      });
    } catch (err) {}
    writeFileSync(fullPath, result, "utf8");
  }
};
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
