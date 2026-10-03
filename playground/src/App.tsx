import { useState } from "react";
import Dropdown from "react-dropdown";
import "react-dropdown/style.css";
import { Panel, PanelGroup, PanelResizeHandle } from "react-resizable-panels";
import "./App.css";
import OptionsModal from "./components/OptionsModal";
import { targetLabel, targets } from "./targets";
import CodeViewer from "./components/CodeViewer";
import FileExplorer, {
  DEFINITION_KEY,
  fileKey,
} from "./components/FileExplorer";
import YamlEditor from "./components/YamlEditor/YamlEditor";
import { examplesForTarget } from "./examples";
import { usePlaygroundPresenter } from "./helpers/usePlaygroundPresenter";
import { downloadZip } from "./helpers/zipTools";
import SQLEditor from "./components/YamlEditor/SQLEditor";
import MarkdownPreview from "./components/MarkdownPreview";

function App() {
  const {
    files: generatedFiles,
    setValue,
    value,
    definitionId,
    selectDefinition,
    setFeatures,
    features,
    compilerTags,
    ready,
    setAssemblyFunction,
    assemblyFunction,
  } = usePlaygroundPresenter();

  const [optionsOpen, setOptionsOpen] = useState(false);
  const files = generatedFiles || [];
  // DEFINITION_KEY shows the Monaco editor, anything else a generated file.
  const [selectedKey, setSelectedKey] = useState(DEFINITION_KEY);
  const activeFile = files.find((f) => fileKey(f) === selectedKey);
  const showEditor = !activeFile;
  const isSql = assemblyFunction === "sqlQueryPredict";
  const definitions = examplesForTarget(assemblyFunction).map((e) => ({
    id: e.id,
    label: e.label,
    detail: `${targetLabel(e.target)}${e.tags.length ? " · " + e.tags.join(", ") : ""}`,
  }));

  return (
    <>
      <header className="intro" style={{}}>
        <div
          style={{
            display: "flex",
            flexDirection: "row",
            justifyContent: "space-between",
            padding: "10px 15px",
            alignItems: "center",
          }}
        >
          <h1>EMI Compiler</h1>
          <span>
            <a target="_blank" href="https://github.com/torabian/emi">
              Github
            </a>{" "}
            <a target="_blank" href="https://torabian.github.io/emi">
              Documentation
            </a>
            <span style={{ marginLeft: "20px" }} className="wasm-status">
              {ready ? "✅" : "⏳"}
            </span>
          </span>
        </div>
        <div
          style={{
            display: "flex",
            backgroundColor: "gray",
            padding: "3px 15px",
            alignContent: "center",
            alignItems: "center",
          }}
        >
          <div style={{ width: "170px" }}>
            <Dropdown
              options={targets}
              onChange={(value) => {
                setAssemblyFunction(value.value);
              }}
              value={assemblyFunction}
              placeholder="Select an option"
            />
          </div>
          <button
            style={{ borderRadius: 0, height: "41px", marginLeft: "5px" }}
            onClick={() => void downloadZip(files)}
          >
            Download ({files?.length || 0})
          </button>
          <button
            style={{ borderRadius: 0, height: "41px", marginLeft: "5px" }}
            onClick={() => setOptionsOpen(true)}
            disabled={!ready || Object.keys(compilerTags).length === 0}
          >
            Compiler tags ({features.length})
          </button>
          <span
            style={{
              marginLeft: "10px",
              fontSize: "12px",
              maxWidth: "340px",
              lineHeight: 1.3,
              color: "#f0f0f0",
            }}
          >
            Switches that change what the {targetLabel(assemblyFunction)}{" "}
            compiler generates.
          </span>
        </div>
      </header>
      {optionsOpen ? (
        <OptionsModal
          language={targetLabel(assemblyFunction)}
          tags={compilerTags[assemblyFunction] ?? []}
          selected={features}
          onChange={setFeatures}
          onClose={() => setOptionsOpen(false)}
        />
      ) : null}
      <PanelGroup
        direction="horizontal"
        style={{ height: "calc(100vh - 100px)" }}
      >
        <Panel defaultSize={22} minSize={10}>
          <FileExplorer
            files={files}
            selectedKey={selectedKey}
            onSelect={setSelectedKey}
            definitions={definitions}
            selectedDefinitionId={definitionId}
            onSelectDefinition={(id) => {
              selectDefinition(id);
              setSelectedKey(DEFINITION_KEY);
            }}
          />
        </Panel>
        <PanelResizeHandle>
          <div style={{ width: 3, height: "100%", cursor: "col-resize" }} />
        </PanelResizeHandle>
        <Panel defaultSize={78} minSize={20}>
          {/* Kept mounted (just hidden) so undo history and cursor survive
              switching to a generated file. */}
          <div style={{ display: showEditor ? "block" : "none" }}>
            {isSql ? (
              <SQLEditor
                path={definitionId}
                value={value}
                onChange={setValue}
              />
            ) : (
              <YamlEditor
                path={definitionId}
                value={value}
                onChange={setValue}
              />
            )}
          </div>
          {activeFile ? (
            activeFile.Extension === ".md" ? (
              <MarkdownPreview source={activeFile.ActualScript} />
            ) : (
              <CodeViewer file={activeFile} />
            )
          ) : null}
        </Panel>
      </PanelGroup>
    </>
  );
}

export default App;
