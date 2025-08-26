#!/usr/bin/env node
import { Command } from "commander";
const program = new Command();

import "./wasm_exec.js";

import fs from "fs";
import { fileURLToPath } from "url";
import path from "path";
import { applyFlags } from "./cliutils.js";

const __dirname = path.dirname(fileURLToPath(import.meta.url));

const go = new Go(); // global from wasm_exec.js
const wasmPath = path.join(__dirname, "emi-compiler.wasm");
const wasmBuffer = fs.readFileSync(wasmPath);

const { instance } = await WebAssembly.instantiate(wasmBuffer, go.importObject);
go.run(instance);

const { TextActions, FileActions } = globalThis.getPublicActions();

for (const a of TextActions) {
  const cmd = program
    .command(a.Name)
    .description(a.Description)
    .allowUnknownOption(true) // Commander is picky; better allow extra
    .action(async (options) => {
      const ctx = {
        path: options.path,
        output: options.output || "",
        tags: options.tags || "",
        content: fs.readFileSync(options.path, "utf8"),
      };

      if (!globalThis[a.WasmFunctionName]) {
        console.log("Wasm function not found:", a.WasmFunctionName);
      }

      const result = await globalThis[a.WasmFunctionName](
        fs.readFileSync(options.path, "utf8"),
        ctx
      );

      if (!ctx.output) {
        console.log(result);
      } else {
        fs.writeFileSync(ctx.output, result, "utf8");
      }
    });
  applyFlags(cmd, a.Flags);
}

// add file actions
for (const a of FileActions) {
  const cmd = program
    .command(a.Name)
    .description(a.Description)
    .action(async (options) => {
      const ctx = {
        path: options.path,
        output: options.output || "",
        Tags: options.tags || "",
        content: fs.readFileSync(options.path, "utf8"),
      };

      if (!globalThis[a.WasmFunctionName]) {
        console.log("Wasm function not found:", a.WasmFunctionName);
      }

      const files = await globalThis[a.WasmFunctionName](
        fs.readFileSync(options.path, "utf8"),
        ctx
      );

      if (!ctx.output) {
        console.log(JSON.stringify(files, null, 2));
      } else {
        for (const f of files) {
          const dir = path.join(ctx.output, f.Location);
          fs.mkdirSync(dir, { recursive: true });
          fs.writeFileSync(
            path.join(dir, f.Name + f.Extension),
            f.ActualScript,
            "utf8"
          );
        }
      }
    });

  applyFlags(cmd, a.Flags);
}

program
  .name("Emi compiler")
  .description("Module3 definitions code generator")
  .version("0.8.0");

program.parse();
