import Editor, { type BeforeMount } from "@monaco-editor/react";
import { configureMonacoYaml } from "monaco-yaml";
import * as monaco from "monaco-editor";

export default function YamlEditor({
  onChange,
  value,
}: {
  onChange: (value: string) => void;
  value?: string;
}) {
  const beforeMount: BeforeMount = () => {
    configureMonacoYaml(monaco, {
      enableSchemaRequest: true,
      validate: true,
      completion: true,
      hover: true,
      schemas: [
        {
          uri: "/emi-definitions.json",
          fileMatch: ["config.yaml"], // must match Editor path
        },
      ],
    });
  };

  return (
    <Editor
      path="config.yaml"
      options={{
        quickSuggestions: true, // show on typing
        suggestOnTriggerCharacters: true, // e.g. after ":" etc.
        wordBasedSuggestions: "allDocuments", // don’t suggest random words
      }}
      height="50vh"
      onChange={(value) => {
        onChange(value as string);
      }}
      width={"calc(100vw - 50px)"}
      defaultLanguage="yaml"
      defaultValue={value}
      beforeMount={beforeMount}
    />
  );
}
