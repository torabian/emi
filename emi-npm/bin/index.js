#!/usr/bin/env node
import { Command } from "commander";
const program = new Command();
import { applyFlags } from "./cliutils.js";
import { FileActions, fileWriter, TextActions } from "./getPublicActions.js";
import { readFileSync, mkdirSync, writeFileSync } from "fs";
import { dirname, join } from "path";

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
        content: readFileSync(options.path, "utf8"),
      };

      if (!globalThis[a.WasmFunctionName]) {
        console.log("Wasm function not found:", a.WasmFunctionName);
      }

      const result = await globalThis[a.WasmFunctionName](
        readFileSync(options.path, "utf8"),
        ctx
      );

      if (!ctx.output) {
        console.log(result);
      } else {
        writeFileSync(ctx.output, result, "utf8");
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
        content: readFileSync(options.path, "utf8"),
        Flags: JSON.stringify(options),
      };

      if (!globalThis[a.WasmFunctionName]) {
        console.log("Wasm function not found:", a.WasmFunctionName);
      }

      const files = await globalThis[a.WasmFunctionName](
        readFileSync(options.path, "utf8"),
        ctx
      );

      if (!ctx.output) {
        console.log(JSON.stringify(files, null, 2));
      } else {
        fileWriter(files, ctx.output);
      }
    });

  applyFlags(cmd, a.Flags);
}

program
  .name("Emi compiler")
  .description("Module3 definitions code generator")
  .version("0.8.0");

program.parse();
