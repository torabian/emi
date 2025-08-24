import { useState } from "react";
import "./App.css";
import FeatureSelector from "./components/FeatureSelector";
import LanguageSelector from "./components/LanguageSelector";
import TypescriptEditor from "./components/YamlEditor/TypescriptEditor";
import YamlEditor from "./components/YamlEditor/YamlEditor";
import { usePlaygroundPresenter } from "./helpers/usePlaygroundPresenter";
import { downloadZip } from "./helpers/zipTools";

function App() {
  const {
    files,
    setValue,
    value,
    setFeatures,
    features,
    ready,
    setAssemblyFunction,
    assemblyFunction,
  } = usePlaygroundPresenter();
  const [activeTab, setActiveTab] = useState(files?.[0]?.Name || "");

  // Update active tab if files change
  if (files?.length && !files.find((f) => f.Name === activeTab)) {
    setActiveTab(files[0].Name);
  }

  const activeFile = files?.find((f) => f.Name === activeTab);

  return (
    <>
      <header className="intro">
        <h1>🔥 EMI JavaScript Compiler</h1>
        <p>
          A WebAssembly-powered playground for generating code dynamically.
          Write input on the yaml editor, see generated files below.
        </p>
        <span className="wasm-status">
          WASM Runtime: {ready ? "✅ Ready" : "⏳ Initializing..."} (
          {assemblyFunction}) [{features.join(",")}]
        </span>
        <LanguageSelector onChange={setAssemblyFunction} />
        <FeatureSelector
          options={["nestjs", "typescript", "axios", "react"]}
          setSelected={(value) => setFeatures(value)}
          selected={features}
        />
        <button onClick={() => void downloadZip(files)}>
          Download the sdk ({files?.length || 0})
        </button>
      </header>

      <div className="app-container">
        <div className="left-pane">
          <YamlEditor value={value} onChange={setValue} />
        </div>

        {files?.length > 0 && (
          <div className="right-pane">
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

            <div className="tab-content">
              {activeFile && (
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
              )}
            </div>
          </div>
        )}
      </div>
    </>
  );
}

export default App;
