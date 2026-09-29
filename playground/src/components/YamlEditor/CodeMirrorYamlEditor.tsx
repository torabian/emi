import { yaml } from "@codemirror/lang-yaml";
import { lintGutter } from "@codemirror/lint";
import CodeMirror from "@uiw/react-codemirror";
import { yamlSchema } from "codemirror-json-schema/yaml";
import { useEffect, useMemo, useState } from "react";
import type { JSONSchema7 } from "json-schema";

const SCHEMA_URL = `${import.meta.env.BASE_URL}emi-module-spec.json`;

const INITIAL = `namespace: 
targets:
  - 
`;

export default function CodeMirrorYamlEditor({
  value = INITIAL,
  onChange,
}: {
  value?: string;
  onChange?: (value: string) => void;
}) {
  const [schema, setSchema] = useState<JSONSchema7>();

  useEffect(() => {
    fetch(SCHEMA_URL)
      .then((r) => r.json())
      .then(setSchema)
      .catch((e) => console.error("Failed to load emi schema", e));
  }, []);

  // yamlSchema bundles yaml(), completion, validation (lint) and hover tooltips.
  const extensions = useMemo(
    () => [lintGutter(), schema ? yamlSchema(schema) : yaml()],
    [schema],
  );

  return (
    <CodeMirror
      value={value}
      height="calc(100vh - 125px)"
      extensions={extensions}
      onChange={onChange}
      basicSetup={{ autocompletion: true }}
    />
  );
}
