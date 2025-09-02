import { writeFileSync } from "fs";
import prettier from "prettier";
import { describe, it } from "vitest";
import { createInstance } from "../../../../emi-npm/bin/getPublicActions";
import yaml from "js-yaml";

describe("Generate documents for object type", () => {
  const content: string[] = [];

  it("should init the wasm", () => {
    createInstance();
  });

  it("generate object doc", async () => {
    content.push(`
---
sidebar_position: 8
---

# Emi object data type
    `);

    const objectExample = {
      name: "MyObjectClass",
      fields: [
        {
          name: "profile",
          type: "object",
          fields: [
            { name: "firstName", type: "string" },
            { name: "age", type: "int" },
          ],
        },
      ],
    };

    content.push(`
Example schema:
\`\`\`yaml
${yaml.dump(objectExample)}
\`\`\`

- \`object\` lets you embed nested fields inside your class.  
- Works like inline type definition in TypeScript.  
- Defaults to an empty object \`{}\` if not specified.  
- \`object?\` allows \`null\`/ \`undefined\`.  
    `);

    const classResult = globalThis.jsGenDtoClass(
      JSON.stringify(objectExample),
      { Tags: "typescript", Flags: "Anonymouse" }
    );
    const formatted = await prettier.format(classResult, {
      parser: "typescript",
    });
    content.push("```ts\r\n" + formatted + "\r\n```");

    const nullableObjectExample = {
      name: "MyObjectClass",
      fields: [
        {
          name: "optionalProfile",
          type: "object?",
          fields: [{ name: "lastName", type: "string" }],
        },
      ],
    };

    content.push(`
Nullable version:
\`\`\`yaml
${yaml.dump(nullableObjectExample)}
\`\`\`

Defaults to \`undefined\`, but you can assign an object with required child fields or \`null\`.
    `);

    const nullableResult = globalThis.jsGenDtoClass(
      JSON.stringify(nullableObjectExample),
      { Tags: "typescript", Flags: "Anonymouse" }
    );
    const formattedNullable = await prettier.format(nullableResult, {
      parser: "typescript",
    });
    content.push("```ts\r\n" + formattedNullable + "\r\n```");

    writeFileSync(
      "../../emi-web/docs/js/emi-object-data-type.mdx",
      content.join("\r\n").trim()
    );
  });
});
