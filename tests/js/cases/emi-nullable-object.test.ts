import { describe, it } from "vitest";
import prettier from "prettier";
import { createInstance } from "../../../emi-npm/bin/getPublicActions";
import { writeFileSync } from "fs";
import yaml from "js-yaml";

const sample1 = {
  name: "nullableResponseAction",
  fields: [
    {
      name: "mother",
      type: "object",
      description: "Mother info should be present",
      fields: [
        {
          name: "firstName",
          type: "string",
        },
      ],
    },
    {
      name: "father",
      type: "object?",
      description: "Father name is not essential for some goverment papers.",
      fields: [
        {
          name: "firstName",
          type: "string",
        },
      ],
    },
    {
      name: "firstUncle",
      type: "one",
      target: "UncleDto",
      description: "Uncle is a separate dto, therefor we use that entity",
    },
    {
      name: "secondUncle",
      type: "one?",
      target: "UncleDto",
      description: "Second uncle is optional",
    },
  ],
};

describe("Generting the nullable values", () => {
  let file;
  const content: string[] = [];

  it("should init the wasm", () => {
    createInstance();
  });

  it("write the topic a little bit", () => {
    content.push(`---
sidebar_position: 4
---

# Nullable object vs non-nullable

Emi definition allows for nullable object vs non-nullable objects. In case an object is not nullable, should always be present and initialised
by the parent class upon instantiation. For types, simply it would indicate that it needs to be present.

 
\`\`\`yaml
${yaml.dump(sample1)}
\`\`\`

`);
  });

  it("should generate and write the Emi JS module", async () => {
    const classResult = globalThis.jsGenDtoClass(JSON.stringify(sample1), {
      Tags: "typescript",
      Flags: "Anonymouse",
    });

    const formatted = await prettier.format(classResult, {
      parser: "typescript",
    });

    content.push("```ts\r\n" + formatted + "\r\n```");
  });

  it("should write the final doc", () => {
    writeFileSync(
      "../../emi-web/docs/js/emi-nullable-object.md",
      content.join("\r\n").trim()
    );
  });
});
