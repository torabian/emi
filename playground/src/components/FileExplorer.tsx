import { useMemo, useState } from "react";
import type { VirtualFile } from "../definitions";
import "./FileExplorer.css";

type FolderNode = {
  kind: "folder";
  name: string;
  path: string;
  children: TreeNode[];
};
type FileNode = { kind: "file"; name: string; key: string };
type TreeNode = FolderNode | FileNode;

// The generators put folders in both Location and Name (e.g. Location
// "sdk/envelopes", Name "common/EnvelopeClass"), so the real path is all of it.
const pathParts = (file: VirtualFile) =>
  `${file.Location ?? ""}/${file.Name}${file.Extension ?? ""}`
    .split("/")
    .filter((p) => p && p !== ".");

// Unique identity of a generated file, used for selection.
export const fileKey = (file: VirtualFile) => pathParts(file).join("/");

function buildTree(files: VirtualFile[]): TreeNode[] {
  const root: FolderNode = { kind: "folder", name: "", path: "", children: [] };

  for (const file of files) {
    const parts = pathParts(file);
    const fileName = parts.pop() as string;
    let folder = root;
    for (const part of parts) {
      const path = folder.path ? `${folder.path}/${part}` : part;
      let next = folder.children.find(
        (c): c is FolderNode => c.kind === "folder" && c.name === part,
      );
      if (!next) {
        next = { kind: "folder", name: part, path, children: [] };
        folder.children.push(next);
      }
      folder = next;
    }
    folder.children.push({
      kind: "file",
      name: fileName,
      key: fileKey(file),
    });
  }

  // Folders first, then files, each alphabetical - like VS Code's explorer.
  const sort = (nodes: TreeNode[]) => {
    nodes.sort((a, b) =>
      a.kind !== b.kind
        ? a.kind === "folder"
          ? -1
          : 1
        : a.name.localeCompare(b.name),
    );
    nodes.forEach((n) => n.kind === "folder" && sort(n.children));
  };
  sort(root.children);
  return root.children;
}

// Selection value of the pinned definition entry (not a generated file).
export const DEFINITION_KEY = "\u0000definition";

export default function FileExplorer({
  files,
  selectedKey,
  onSelect,
  definitions,
  selectedDefinitionId,
  onSelectDefinition,
}: {
  files: VirtualFile[];
  /** DEFINITION_KEY while the editor is shown, otherwise a generated file's key. */
  selectedKey: string;
  onSelect: (key: string) => void;
  definitions: { id: string; label: string }[];
  selectedDefinitionId: string;
  onSelectDefinition: (id: string) => void;
}) {
  const tree = useMemo(() => buildTree(files), [files]);
  const [collapsed, setCollapsed] = useState<Set<string>>(new Set());

  const toggle = (path: string) =>
    setCollapsed((prev) => {
      const next = new Set(prev);
      if (!next.delete(path)) next.add(path);
      return next;
    });

  const renderNodes = (nodes: TreeNode[], depth: number) =>
    nodes.map((node) => {
      const indent = { paddingLeft: 8 + depth * 12 };

      if (node.kind === "folder") {
        const isCollapsed = collapsed.has(node.path);
        return (
          <div key={node.path}>
            <div
              className="fx-row fx-folder"
              style={indent}
              onClick={() => toggle(node.path)}
            >
              <span className="fx-chevron">{isCollapsed ? "▸" : "▾"}</span>
              {node.name}
            </div>
            {isCollapsed ? null : renderNodes(node.children, depth + 1)}
          </div>
        );
      }

      return (
        <div
          key={node.key}
          className={`fx-row fx-file ${node.key === selectedKey ? "active" : ""}`}
          style={indent}
          title={node.key}
          onClick={() => onSelect(node.key)}
        >
          <span className="fx-chevron" />
          {node.name}
        </div>
      );
    });

  return (
    <div className="file-explorer">
      <div className="fx-section">Examples</div>
      {definitions.map((definition) => (
        <div
          key={definition.id}
          className={`fx-row fx-definition ${
            selectedKey === DEFINITION_KEY &&
            definition.id === selectedDefinitionId
              ? "active"
              : ""
          }`}
          style={{ paddingLeft: 8 }}
          onClick={() => onSelectDefinition(definition.id)}
        >
          <span className="fx-chevron" />
          {definition.label}
        </div>
      ))}
      {tree.length ? (
        <>
          <div className="fx-section fx-divider">Generated files</div>
          {renderNodes(tree, 0)}
        </>
      ) : null}
    </div>
  );
}
