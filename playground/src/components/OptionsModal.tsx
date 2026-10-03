import { useEffect } from "react";
import "./OptionsModal.css";

export type CompilerTag = { Tag: string; Description: string };

type OptionsModalProps = {
  /** Display name of the compiler the options belong to. */
  language: string;
  /** Options the compiler reports (from wasm). */
  tags: CompilerTag[];
  selected: string[];
  onChange: (selected: string[]) => void;
  onClose: () => void;
};

export default function OptionsModal({
  language,
  tags,
  selected,
  onChange,
  onClose,
}: OptionsModalProps) {
  useEffect(() => {
    const onKey = (e: KeyboardEvent) => e.key === "Escape" && onClose();
    window.addEventListener("keydown", onKey);
    return () => window.removeEventListener("keydown", onKey);
  }, [onClose]);

  const toggle = (tag: string) =>
    onChange(
      selected.includes(tag)
        ? selected.filter((t) => t !== tag)
        : [...selected, tag],
    );

  return (
    <div className="om-backdrop" onClick={onClose}>
      <div
        className="om-dialog"
        role="dialog"
        aria-label={`${language} compiler tags`}
        onClick={(e) => e.stopPropagation()}
      >
        <div className="om-header">
          <h2>{language} compiler tags</h2>
          <button className="om-close" onClick={onClose} aria-label="Close">
            ×
          </button>
        </div>
        <div className="om-body">
          {tags.length === 0 ? (
            <p className="om-empty">This compiler has no options.</p>
          ) : (
            tags.map((t) => (
              <label key={t.Tag} className="om-option">
                <input
                  type="checkbox"
                  checked={selected.includes(t.Tag)}
                  onChange={() => toggle(t.Tag)}
                />
                <span>
                  <code className="om-name">{t.Tag}</code>
                  <span className="om-desc">{t.Description}</span>
                </span>
              </label>
            ))
          )}
        </div>
      </div>
    </div>
  );
}
