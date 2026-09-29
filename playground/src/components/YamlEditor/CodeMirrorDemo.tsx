import { useState } from "react";
import CodeMirrorYamlEditor from "./CodeMirrorYamlEditor";

export default function CodeMirrorDemo() {
  const [value, setValue] = useState<string | undefined>(undefined);
  return (
    <div style={{ padding: 16 }}>
      <h3>CodeMirror YAML + emi schema (Ctrl+Space for completion)</h3>
      <CodeMirrorYamlEditor value={value} onChange={setValue} />
    </div>
  );
}
