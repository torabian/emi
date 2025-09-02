import { createInstance } from "../../emi-npm/bin/getPublicActions";
import yaml from "js-yaml";
import { MethodDeclaration, Project, PropertyDeclaration, ts } from "ts-morph";

function toYaml(obj: unknown): string {
  return yaml.dump(obj, { noRefs: true });
}

// Initialize the wasm
createInstance();

function parseGenerated(code: string) {
  const project = new Project();
  const sf = project.createSourceFile("temp.ts", code);
  return sf;
}

function checkTsCode(code: string) {
  const project = new Project({
    useInMemoryFileSystem: true,
    skipAddingFilesFromTsConfig: true,
    compilerOptions: {
      target: ts.ScriptTarget.ES2015, // or higher, e.g. ES2020
    },
  });

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
  context = {}
) {
  const resp = globalThis[actionWasmFunctionName](
    toYaml(emiActionDefinition),
    context
  );

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

export { parseGenerated, runEmiActionTs, getJsDoc };
