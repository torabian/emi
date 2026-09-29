import hljs from "highlight.js/lib/core";
import bash from "highlight.js/lib/languages/bash";
import c from "highlight.js/lib/languages/c";
import cpp from "highlight.js/lib/languages/cpp";
import csharp from "highlight.js/lib/languages/csharp";
import dart from "highlight.js/lib/languages/dart";
import go from "highlight.js/lib/languages/go";
import java from "highlight.js/lib/languages/java";
import javascript from "highlight.js/lib/languages/javascript";
import json from "highlight.js/lib/languages/json";
import kotlin from "highlight.js/lib/languages/kotlin";
import php from "highlight.js/lib/languages/php";
import python from "highlight.js/lib/languages/python";
import sql from "highlight.js/lib/languages/sql";
import swift from "highlight.js/lib/languages/swift";
import typescript from "highlight.js/lib/languages/typescript";
import xml from "highlight.js/lib/languages/xml";
import yaml from "highlight.js/lib/languages/yaml";
import { useMemo } from "react";
import type { VirtualFile } from "../definitions";
import {
  CODE_FONT_FAMILY,
  CODE_FONT_SIZE,
  CODE_LINE_HEIGHT,
} from "../helpers/codeFont";
import "./CodeViewer.css";

Object.entries({
  bash,
  c,
  cpp,
  csharp,
  dart,
  go,
  java,
  javascript,
  json,
  kotlin,
  php,
  python,
  sql,
  swift,
  typescript,
  xml,
  yaml,
}).forEach(([name, definition]) => hljs.registerLanguage(name, definition));

// Highlighting very large outputs blocks the main thread for little benefit.
const MAX_HIGHLIGHT_LENGTH = 300_000;

const languageByExtension: Record<string, string> = {
  ts: "typescript",
  tsx: "typescript",
  js: "javascript",
  json: "json",
  yaml: "yaml",
  yml: "yaml",
  py: "python",
  dart: "dart",
  cs: "csharp",
  java: "java",
  php: "php",
  c: "c",
  h: "c",
  cpp: "cpp",
  hpp: "cpp",
  go: "go",
  kt: "kotlin",
  swift: "swift",
  sql: "sql",
  sh: "bash",
  xml: "xml",
  html: "xml",
};

// Extension is stored with its leading dot (".py", ".ts"...); some outputs only
// carry the extension in their name instead.
function guessLanguage(file: VirtualFile): string | undefined {
  const ext = (file.Extension || file.Name.split(".").pop() || "").replace(
    /^\./,
    "",
  );
  return languageByExtension[ext.toLowerCase()];
}

export default function CodeViewer({ file }: { file: VirtualFile }) {
  const html = useMemo(() => {
    const code = file.ActualScript ?? "";
    const language = guessLanguage(file);
    if (!language || code.length > MAX_HIGHLIGHT_LENGTH) {
      return null;
    }
    return hljs.highlight(code, { language, ignoreIllegals: true }).value;
  }, [file]);

  return (
    <pre
      className="code-viewer"
      style={{
        fontFamily: CODE_FONT_FAMILY,
        fontSize: CODE_FONT_SIZE,
        lineHeight: `${CODE_LINE_HEIGHT}px`,
      }}
    >
      {html === null ? (
        <code>{file.ActualScript}</code>
      ) : (
        <code className="hljs" dangerouslySetInnerHTML={{ __html: html }} />
      )}
    </pre>
  );
}
