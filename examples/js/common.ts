import { createInstance } from "../emi-wasm-helper/getPublicActions";
import yaml from "js-yaml";
import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";
import {
  MethodDeclaration,
  ModuleKind,
  ModuleResolutionKind,
  Project,
  PropertyDeclaration,
  ts,
} from "ts-morph";
import { WebSocketServer } from "ws";

const __dirname = path.dirname(fileURLToPath(import.meta.url));

// Generated SDK code imports its runtime helpers from "./sdk/common/*". To
// type-check that code in isolation, we mirror the real SDK-kit sources into
// the in-memory project under "/sdk/common" so every import resolves against
// the actual exports instead of a hand-written stub that can drift.
const SDK_COMMON_SRC = path.join(__dirname, "../js-sdk-kit/src/common");

const sdkCommonModules = fs
  .readdirSync(SDK_COMMON_SRC)
  .filter((file) => file.endsWith(".ts"))
  .map((file) => {
    const name = file.replace(/\.ts$/, "");
    let content = fs.readFileSync(path.join(SDK_COMMON_SRC, file), "utf8");

    if (name === "operators") {
      // The Go generator imports a single `operators` symbol from this module
      // (see lib/js/js-common-object.go). The source file exposes the operator
      // classes individually, so surface them under an `operators` namespace
      // here to keep the generated import resolvable.
      content +=
        "\nexport const operators = { MOne, MOneNullable, MArray, MArrayNullable, MCollection, MCollectionNullable };\n";
    }

    return { filePath: `/sdk/common/${name}.ts`, content };
  });

function toYaml(obj: unknown): string {
  return yaml.dump(obj, { noRefs: true });
}

// Initialize the wasm
createInstance();

function parseGenerated(code: string) {
  const project = new Project({
    compilerOptions: {
      moduleResolution: ModuleResolutionKind.NodeNext,
    },
  });

  const sf = project.createSourceFile("temp.ts", code);

  return sf;
}

function checkTsCode(code: string) {
  const project = new Project({
    useInMemoryFileSystem: true,
    skipAddingFilesFromTsConfig: true,
    skipFileDependencyResolution: true,
    compilerOptions: {
      target: ts.ScriptTarget.ES2019, // or higher, e.g. ES2020,\
    },
  });

  for (const module of sdkCommonModules) {
    project.createSourceFile(module.filePath, module.content);
  }

  const sourceFile = project.createSourceFile("temp.ts", code);

  // Get syntax & type errors
  const diagnostics = sourceFile.getPreEmitDiagnostics();

  if (diagnostics.length === 0) {
    return { valid: true };
  }

  const errors = diagnostics.map((d) => ({
    message: d.getMessageText(),
    line: d.getLineNumber(),
    column: d.getStart(),
  }));

  return { valid: false, errors };
}

function runEmiActionTs(
  actionWasmFunctionName,
  emiActionDefinition,
  context = {},
) {
  let resp = globalThis[actionWasmFunctionName](
    toYaml(emiActionDefinition),
    context,
  );

  // resp = resp.replace(/^import .*$/gm, "");

  resp += `
  
  function isPlausibleObject(v: any) { return false }
  `;

  const validation = checkTsCode(resp);
  if (!validation.valid) {
    console.error(validation.errors);
    console.log(resp);
    throw validation.errors;
  }

  return { source: parseGenerated(resp), resp };
}

function getJsDoc(node: PropertyDeclaration | MethodDeclaration): string {
  const jsDocs = node.getJsDocs();

  if (jsDocs.length === 0) return "no jsdoc";

  return jsDocs.map((d) => d.getDescription() || "").join("\n");
}

// --- helper: start a WebSocket server ---
function createSocketServer(handler: (ws: WebSocket, msg: any) => void, port) {
  const server = new WebSocketServer({ port });
  server.on("connection", (ws) => {
    ws.on("message", async (msg) => {
      const data = JSON.parse(msg.toString());
      await handler(ws, data);
    });
  });
  return server;
}

// --- helper: run client test against a server ---
function runSocketClientTest(
  createMessage: () => any,
  verify: (received: any[]) => void,
  ws: WebSocket,
) {
  return new Promise<void>((resolve) => {
    const received: any[] = [];

    ws.addEventListener("message", (msg) => {
      received.push(msg.data);
      verify(received);
      if (received.length >= 5) ws.close();
    });

    ws.addEventListener("open", () => ws.send(JSON.stringify(createMessage())));
    ws.addEventListener("close", resolve);
  });
}

function randomBetween(min, max) {
  // inclusive min and max
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

export {
  parseGenerated,
  randomBetween,
  runEmiActionTs,
  getJsDoc,
  createSocketServer,
  runSocketClientTest,
};
