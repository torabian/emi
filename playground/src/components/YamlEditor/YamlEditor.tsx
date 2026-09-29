import Editor, { loader } from "@monaco-editor/react";
import * as monaco from "monaco-editor";
import editorWorker from "monaco-editor/esm/vs/editor/editor.worker?worker";
import { configureMonacoYaml } from "monaco-yaml";
import yamlWorker from "monaco-yaml/yaml.worker?worker";
import {
  CODE_FONT_FAMILY,
  CODE_FONT_SIZE,
  CODE_LINE_HEIGHT,
} from "../../helpers/codeFont";
import { useSystemTheme } from "../../helpers/useSystemTheme";

// Use the bundled monaco instead of @monaco-editor/react's CDN copy; otherwise
// monaco-yaml would configure a different instance than the one rendered.
loader.config({ monaco });

self.MonacoEnvironment = {
  getWorker: (_id, label) => {
    switch (label) {
      case "yaml":
        return new yamlWorker();
      default:
        return new editorWorker();
    }
  },
};

const monacoYaml = configureMonacoYaml(monaco, {
  enableSchemaRequest: true,
  validate: true,
  completion: true,
  hover: true,
  schemas: [
    {
      // Served from /public, so it lives under Vite's base path.
      uri: new URL(
        `${import.meta.env.BASE_URL}emi-module-spec.json`,
        window.location.origin,
      ).href,
      fileMatch: ["*"], // every yaml model in the playground uses the emi schema
    },
  ],
});

// Without this every hot reload registers another set of providers, which
// shows up as duplicated suggestions.
import.meta.hot?.dispose(() => monacoYaml.dispose());

export default function YamlEditor({
  onChange,
  value,
  path,
}: {
  onChange: (value: string) => void;
  value?: string;
  /** Unique per definition; Monaco keeps a separate model (and undo history) per path. */
  path: string;
}) {
  const theme = useSystemTheme();

  return (
    <Editor
      path={`${path}.yml`}
      options={{
        quickSuggestions: true, // show on typing
        suggestOnTriggerCharacters: true, // e.g. after ":" etc.
        wordBasedSuggestions: "off", // schema suggestions only; words duplicate them
        automaticLayout: true, // re-fit when shown again after being hidden
        fontFamily: CODE_FONT_FAMILY,
        fontSize: CODE_FONT_SIZE,
        lineHeight: CODE_LINE_HEIGHT,
      }}
      height="calc(100vh - 125px)"
      onChange={(value) => {
        onChange(value as string);
      }}
      defaultLanguage="yaml"
      defaultValue={value}
      theme={theme}
      onMount={(editor) => {
        // Monaco only auto-opens suggestions on word characters; also open
        // them after Enter (new line + auto-indent) so keys are offered.
        editor.onDidChangeModelContent((e) => {
          if (e.changes.some((c) => /^\r?\n[ \t]*$/.test(c.text))) {
            setTimeout(
              () => editor.trigger("yaml", "editor.action.triggerSuggest", {}),
              0,
            );
          }
        });
      }}
    />
  );
}
