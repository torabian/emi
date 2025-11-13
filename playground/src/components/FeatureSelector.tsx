type FeatureSelectorProps = {
  options: string[]; // e.g., ["NestJS", "TypeScript", "React"]
  setSelected: React.Dispatch<React.SetStateAction<string[]>>;
  selected: Array<string>;
};

export default function FeatureSelector({
  options,
  selected,
  setSelected,
}: FeatureSelectorProps) {
  const handleToggle = (feature: string) => {
    setSelected((prev) => {
      const isChecked = prev.includes(feature);
      const updated = isChecked
        ? prev.filter((f) => f !== feature)
        : [...prev, feature];

      return updated;
    });
  };

  return (
    <div className="feature-selector">
      {options.map((feature) => (
        <label key={feature}>
          <input
            type="checkbox"
            checked={selected.includes(feature)}
            onChange={() => handleToggle(feature)}
          />
          <span>{feature}</span>
        </label>
      ))}
    </div>
  );
}
