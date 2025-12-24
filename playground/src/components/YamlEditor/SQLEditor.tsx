import Editor, { type BeforeMount } from "@monaco-editor/react";
import { configureMonacoYaml } from "monaco-yaml";
import * as monaco from "monaco-editor";
import { useSystemTheme } from "../../helpers/useSystemTheme";

export default function SQLEditor({
  onChange,
  value,
}: {
  onChange: (value: string) => void;
  value?: string;
}) {
  const theme = useSystemTheme();

  const beforeMount: BeforeMount = () => {
    configureMonacoYaml(monaco, {
      enableSchemaRequest: true,
      validate: true,
      completion: true,
      hover: true,
      schemas: [
        {
          uri: "/emi-module-spec.json",
          fileMatch: ["config.yaml"], // must match Editor path
        },
      ],
    });
  };

  return (
    <Editor
      path="query.sql"
      options={{
        quickSuggestions: true, // show on typing
        suggestOnTriggerCharacters: true, // e.g. after ":" etc.
        wordBasedSuggestions: "allDocuments", // don’t suggest random words
        fontSize: 13,
      }}
      height="calc(100vh - 125px)"
      onChange={(value) => {
        onChange(value as string);
      }}
      defaultLanguage="sql"
      defaultValue={value}
      theme={theme}
      beforeMount={beforeMount}
    />
  );
}
