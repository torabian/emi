import Editor from "@monaco-editor/react";
import {
  CODE_FONT_FAMILY,
  CODE_FONT_SIZE,
  CODE_LINE_HEIGHT,
} from "../../helpers/codeFont";
import { useSystemTheme } from "../../helpers/useSystemTheme";

export default function SQLEditor({
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
      path={path}
      options={{
        quickSuggestions: true, // show on typing
        suggestOnTriggerCharacters: true, // e.g. after ":" etc.
        wordBasedSuggestions: "allDocuments", // don’t suggest random words
        automaticLayout: true, // re-fit when shown again after being hidden
        fontFamily: CODE_FONT_FAMILY,
        fontSize: CODE_FONT_SIZE,
        lineHeight: CODE_LINE_HEIGHT,
      }}
      height="calc(100vh - 125px)"
      onChange={(value) => {
        onChange(value as string);
      }}
      defaultLanguage="sql"
      defaultValue={value}
      theme={theme}
    />
  );
}
