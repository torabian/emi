import { useState } from "react";

type Option = {
  label: string;
  value: string;
};

const options: Option[] = [
  { label: "JavaScript (Common)", value: "javascript" },
];

export default function LanguageSelector({
  onChange,
}: {
  onChange: (value: string) => void;
}) {
  const [selected, setSelected] = useState(options[0].value);

  const handleChange = (value: string) => {
    setSelected(value);
    onChange(value); // propagate up (for sending to wasm/assembly)
  };

  return (
    <div className="radio-group">
      {options.map((opt) => (
        <label key={opt.value} className="radio-option">
          <input
            type="radio"
            name="language"
            value={opt.value}
            checked={selected === opt.value}
            onChange={() => handleChange(opt.value)}
          />
          {opt.label}
        </label>
      ))}
    </div>
  );
}
