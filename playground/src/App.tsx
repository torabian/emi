import { useState } from "react";
import Dropdown from "react-dropdown";
import "react-dropdown/style.css";
import { Panel, PanelGroup, PanelResizeHandle } from "react-resizable-panels";
import "./App.css";
import FeatureSelector from "./components/FeatureSelector";
import TypescriptEditor from "./components/YamlEditor/TypescriptEditor";
import YamlEditor from "./components/YamlEditor/YamlEditor";
import { usePlaygroundPresenter } from "./helpers/usePlaygroundPresenter";
import { downloadZip } from "./helpers/zipTools";
import SQLEditor from "./components/YamlEditor/SQLEditor";
const options = [
  { value: "goGen", label: "Golang" },
  { value: "kotlinGen", label: "Kotlin" },
  { value: "jsGenModule", label: "JavaScript" },
  { value: "sqlQueryPredict", label: "QueryPredict(SQL)" },
  { value: "swiftGen", label: "Swift(All)" },
];

function App() {
  const {
    files: generatedFiles,
    setValue,
    value,
    setFeatures,
    features,
    ready,
    setAssemblyFunction,
    assemblyFunction,
  } = usePlaygroundPresenter();

  const files = generatedFiles || [];
  const [activeTab, setActiveTab] = useState(files?.[0]?.Name || "");
  const [direction, setDirection] = useState<any>("horizontal");

  // Update active tab if files change
  if (files?.length && !files.find((f) => f.Name === activeTab)) {
    setActiveTab(files[0].Name);
  }

  const activeFile = files?.find((f) => f.Name === activeTab);

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
          <button
            style={{ borderRadius: 0, height: "41px", marginLeft: "5px" }}
            onClick={() =>
              setDirection((direction: string) =>
                direction === "horizontal" ? "vertical" : "horizontal"
              )
            }
          >
            {direction === "horizontal" ? "V" : "H"}
          </button>
          <div style={{ display: "flex" }}>
            {assemblyFunction === "jsGenModule" ? (
              <FeatureSelector
                options={["nestjs", "typescript", "react"]}
                setSelected={(value) => setFeatures(value)}
                selected={features}
              />
            ) : null}
          </div>
        </div>
      </header>
      <PanelGroup
        direction={direction}
        style={{ height: "calc(100vh - 100px)" }}
      >
        <Panel defaultSize={50} minSize={10}>
          {"sqlQueryPredict" === assemblyFunction ? (
            <SQLEditor value={value} onChange={setValue} />
          ) : (
            <YamlEditor value={value} onChange={setValue} />
          )}
        </Panel>
        <PanelResizeHandle>
          <div style={{ width: 5, background: "gray", cursor: "col-resize" }} />
        </PanelResizeHandle>
        <Panel defaultSize={50} minSize={10}>
          <div className="tabs">
            {files.map((file) => (
              <div
                key={file.Name}
                className={`tab ${file.Name === activeTab ? "active" : ""}`}
                onClick={() => setActiveTab(file.Name)}
              >
                {file.Location ? file.Location + "/" : ""}
                {file.Name}
                {file.Extension ? "" + file.Extension : ""}
              </div>
            ))}
          </div>
          {activeFile ? (
            <>
              <TypescriptEditor
                key={activeFile.Name + activeFile.ActualScript} // force remount when content changes
                allFiles={files}
                value={activeFile.ActualScript}
                onChange={(newValue) => {
                  // Update the corresponding file content
                  activeFile.ActualScript = newValue;
                }}
                file={activeFile}
              />
            </>
          ) : null}
        </Panel>
      </PanelGroup>
    </>
  );
}

export default App;
