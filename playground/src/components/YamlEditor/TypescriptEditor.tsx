import * as monaco from "monaco-editor";

import Editor, { type OnMount } from "@monaco-editor/react";
import type { VirtualFile } from "../../definitions";
import { useSystemTheme } from "../../helpers/useSystemTheme";

const uriCheck = (uri: string) => uri.replaceAll("////", "///");

export default function TypescriptEditor({
  onChange,
  value,
  allFiles,
  file,
}: {
  onChange: (value: string) => void;
  value?: string;
  file: VirtualFile;
  allFiles: VirtualFile[]; // 👈 pass all virtual files here
}) {
  const theme = useSystemTheme();

  const handleMount: OnMount = () => {
    registerVirtualFiles(allFiles);
  };

  return (
    <Editor
      path={uriCheck(`file:///${file.Location}/${file.Name}.${file.Extension}`)}
      options={{
        quickSuggestions: true, // show on typing
        suggestOnTriggerCharacters: true, // e.g. after ":" etc.
        wordBasedSuggestions: "allDocuments", // don’t suggest random words
        fontSize: 13,
      }}
      onMount={handleMount}
      height="calc(100vh - 160px)"
      onChange={(value) => {
        onChange(value as string);
      }}
      theme={theme}
      defaultLanguage={
        file.Name.includes(".ts")
          ? "typescript"
          : file.Name.includes("json")
            ? "json"
            : file.Name.includes(".js")
              ? "javascript"
              : ""
      }
      defaultValue={value}
    />
  );
}

export const registerVirtualFiles = (files: VirtualFile[]) => {
  files.forEach((file) => {
    const x = uriCheck(
      `file:///${file.Location}/${file.Name}${file.Extension ? "." + file.Extension : ""}`
    );
    const uri = monaco.Uri.parse(x);
    let model = monaco.editor.getModel(uri);

    if (!model) {
      model = monaco.editor.createModel(
        file.ActualScript,
        guessLanguage(file.Extension),
        uri
      );
    } else {
      model.setValue(file.ActualScript);
    }
  });
};

// crude but works for now
function guessLanguage(ext: string) {
  switch (ext) {
    case "ts":
    case "tsx":
      return "typescript";
    case "js":
      return "javascript";
    case "json":
      return "json";
    case "yaml":
    case "yml":
      return "yaml";
    default:
      return "plaintext";
  }
}
