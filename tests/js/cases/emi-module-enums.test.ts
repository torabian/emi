import { describe, it } from "vitest";
import prettier from "prettier";
import { createInstance } from "../../../emi-npm/bin/getPublicActions";
import { writeFileSync } from "fs";
import yaml from "js-yaml";

const sample1 = {
  name: "emiEnums",
  enums: [
    {
      name: "enum1",
      fields: [
        {
          k: "Key1",
          value: "Value1",
        },
        {
          k: "Key2",
          value: "Value2",
        },
      ],
    },
  ],
};

describe("Generting the enum values", () => {
  const content: string[] = [];

  it("should init the wasm", () => {
    createInstance();
  });

  it("write the topic a little bit", () => {
    content.push(`---
sidebar_position: 4
---

# Enums in modules

Besides defining inline enums, emi definition allows for standalone enums which is really useful,
for sharing enums between multiple files.

 
\`\`\`yaml
${yaml.dump(sample1)}
\`\`\`

`);
  });

  it("should generate and write the Emi JS module", async () => {
    const catalog = globalThis.jsGen(JSON.stringify(sample1), {
      Tags: "typescript",
      Flags: "Anonymouse",
    });

    const classResult = catalog.find((x) => x.Name === "Enum1").ActualScript;

    const formatted = await prettier.format(classResult, {
      parser: "typescript",
    });

    content.push("```ts\r\n" + formatted + "\r\n```");
  });

  it("should write the final doc", () => {
    writeFileSync(
      "../../emi-web/docs/js/emi-enum.md",
      content.join("\r\n").trim()
    );
  });
});
