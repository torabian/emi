import { useState } from "react";
import Dropdown from "react-dropdown";
import "react-dropdown/style.css";
import { Panel, PanelGroup, PanelResizeHandle } from "react-resizable-panels";
import "./App.css";
import FeatureSelector from "./components/FeatureSelector";
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
const options = [
  { value: "goGen", label: "Golang" },
  { value: "kotlinGen", label: "Kotlin" },
  { value: "jsGenModule", label: "JavaScript" },
  { value: "pythonGenModule", label: "Python" },
  { value: "dartGenModule", label: "Dart" },
  { value: "csharpGenModule", label: "C#" },
  { value: "javaGenModule", label: "Java" },
  { value: "phpGenModule", label: "PHP" },
  { value: "cGenModule", label: "C" },
  { value: "cppGenModule", label: "C++" },
  { value: "sqlQueryPredict", label: "QueryPredict(SQL)" },
  { value: "preprocessorGen", label: "Preprocessor" },
  { value: "postmanGen", label: "Postman" },
  { value: "openapiGen", label: "OpenApi" },
  { value: "mdGen", label: "Markdown" },
  { value: "swiftGen", label: "Swift(All)" },
];

function App() {
  const {
    files: generatedFiles,
    setValue,
    value,
    definitionId,
    selectDefinition,
    setFeatures,
    features,
    ready,
    setAssemblyFunction,
    assemblyFunction,
  } = usePlaygroundPresenter();

  const files = generatedFiles || [];
  // DEFINITION_KEY shows the Monaco editor, anything else a generated file.
  const [selectedKey, setSelectedKey] = useState(DEFINITION_KEY);
  const activeFile = files.find((f) => fileKey(f) === selectedKey);
  const showEditor = !activeFile;
  const isSql = assemblyFunction === "sqlQueryPredict";
  const definitions = examplesForTarget(assemblyFunction);

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
              options={options}
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
          <div style={{ display: "flex" }}>
            {assemblyFunction === "goGen" ? (
              <FeatureSelector
                options={[
                  "no-client",
                  "skip-gin",
                  "split-gin",
                  "skip-cli",
                  "split-cli",
                  "skip-http",
                  "split-http",
                ]}
                setSelected={(value) => setFeatures(value)}
                selected={features}
              />
            ) : null}
            {assemblyFunction === "jsGenModule" ? (
              <FeatureSelector
                options={[
                  "typescript",
                  "react",
                  "nestjs",
                  "no-class",
                  "no-jsdoc",
                  "no-definition",
                  "no-sdk",
                  "no-package",
                  "no-envelope",
                  "include-ext",
                  "json-schema",
                ]}
                setSelected={(value) => setFeatures(value)}
                selected={features}
              />
            ) : null}
            {assemblyFunction === "pythonGenModule" ? (
              <FeatureSelector
                options={["async", "no-sdk"]}
                setSelected={(value) => setFeatures(value)}
                selected={features}
              />
            ) : null}
            {[
              "dartGenModule",
              "csharpGenModule",
              "javaGenModule",
              "phpGenModule",
              "cGenModule",
            ].includes(assemblyFunction) ? (
              <FeatureSelector
                options={["no-sdk", "no-pkg"]}
                setSelected={(value) => setFeatures(value)}
                selected={features}
              />
            ) : null}
            {assemblyFunction === "cppGenModule" ? (
              <FeatureSelector
                options={["unreal", "no-sdk", "no-pkg"]}
                setSelected={(value) => setFeatures(value)}
                selected={features}
              />
            ) : null}
          </div>
        </div>
      </header>
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
