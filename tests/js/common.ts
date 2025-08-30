import { createInstance } from "../../emi-npm/bin/getPublicActions";
import yaml from "js-yaml";
import { Project } from "ts-morph";

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

function runEmiActionTs(
  actionWasmFunctionName,
  emiActionDefinition,
  context = {}
) {
  const resp = globalThis[actionWasmFunctionName](
    toYaml(emiActionDefinition),
    context
  );
  return { source: parseGenerated(resp), resp };
}

export { parseGenerated, runEmiActionTs };
