import { useState } from "react";

type FeatureSelectorProps = {
  options: string[]; // e.g., ["NestJS", "TypeScript", "React"]
  onChange: (selected: string[]) => void; // callback with current selected
};

export default function FeatureSelector({
  options,
  onChange,
}: FeatureSelectorProps) {
  const [selected, setSelected] = useState<string[]>([]);

  const handleToggle = (feature: string) => {
    setSelected((prev) => {
      const isChecked = prev.includes(feature);
      const updated = isChecked
        ? prev.filter((f) => f !== feature)
        : [...prev, feature];
      onChange(updated);
      return updated;
    });
  };

  return (
    <div className="feature-selector">
      {options.map((feature) => (
        <label key={feature} style={{ marginRight: 12, cursor: "pointer" }}>
          <input
            type="checkbox"
            checked={selected.includes(feature)}
            onChange={() => handleToggle(feature)}
          />
          {feature}
        </label>
      ))}
    </div>
  );
}
