import { useState } from "react";
import "./App.css";
import FeatureSelector from "./components/FeatureSelector";
import TypescriptEditor from "./components/YamlEditor/TypescriptEditor";
import YamlEditor from "./components/YamlEditor/YamlEditor";
import type { VirtualFile } from "./definitions";
import { usePlaygroundPresenter } from "./helpers/usePlaygroundPresenter";
import { downloadZip } from "./helpers/zipTools";

function App() {
  const {
    files: generatedFiles,
    setValue,
    value,
    setFeatures,
    features,
    ready,
    // setAssemblyFunction,
    // assemblyFunction,
  } = usePlaygroundPresenter();

  const yamlDoc: VirtualFile = {
    ActualScript: "",
    Extension: "",
    Location: "",
    MimeType: "yaml",
    Name: "definition",
  };

  const files = [yamlDoc, ...(generatedFiles || [])];
  const [activeTab, setActiveTab] = useState(files?.[0]?.Name || "");

  // Update active tab if files change
  if (files?.length && !files.find((f) => f.Name === activeTab)) {
    setActiveTab(files[0].Name);
  }

  const activeFile = files?.find((f) => f.Name === activeTab);

  return (
    <>
      <header className="intro">
        <div
          style={{
            display: "flex",
            flexDirection: "row",
            justifyContent: "space-between",
          }}
        >
          <h1>🔥 EMI Compiler</h1>
          <p>
            A WebAssembly-powered playground for generating code dynamically.
          </p>
        </div>
        <div
          style={{
            display: "flex",
            justifyItems: "center",
            alignItems: "center",
            marginBottom: "20px",
            justifyContent: "space-between",
          }}
        >
          <div style={{ display: "flex" }}>
            <span className="wasm-status">
              WASM Runtime: {ready ? "✅ Ready" : "⏳ Initializing..."}
            </span>
            <FeatureSelector
              options={[
                "nestjs",
                "typescript",
                "axios",
                "react",
                "axiosbundle",
              ]}
              setSelected={(value) => setFeatures(value)}
              selected={features}
            />
          </div>
          <button
            style={{ marginBottom: "5px" }}
            onClick={() => void downloadZip(files)}
          >
            Download the sdk ({files?.length || 0})
          </button>
        </div>
        {/* <LanguageSelector onChange={setAssemblyFunction} /> */}
      </header>

      <div className="app-container">
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

            {activeFile ? (
              <div className="tab-content">
                {activeFile && activeFile.Name == "definition" ? (
                  <YamlEditor value={value} onChange={setValue} />
                ) : (
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
            ) : null}
          </div>
        )}
      </div>
    </>
  );
}

export default App;
