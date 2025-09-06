import { writeFileSync } from "fs";
import prettier from "prettier";
import { describe, it } from "vitest";
import { createInstance } from "../../../emi-npm/bin/getPublicActions";
import yaml from "js-yaml";

describe("Generate documents for array type", () => {
  const content: string[] = [];

  it("should init the wasm", () => {
    createInstance();
  });

  it("generate array doc", async () => {
    content.push(`
---
sidebar_position: 9
---

# Emi array data type
    `);

    const arrayExample = {
      name: "MyArrayClass",
      fields: [
        {
          name: "contacts",
          type: "array",
          fields: [
            { name: "email", type: "string" },
            { name: "phone", type: "string?" },
          ],
        },
      ],
    };

    content.push(`
Example schema:
\`\`\`yaml
${yaml.dump(arrayExample)}
\`\`\`

- \`array\` represents a list of items.  
- You must define an \`items\` field with the child type.  
- Defaults to an empty array \`[]\` if not specified.  
- \`array?\` allows \`null\`/ \`undefined\`.  
    `);

    const classResult = globalThis.jsGenDtoClass(JSON.stringify(arrayExample), {
      Tags: "typescript",
      Flags: "Anonymouse",
    });
    const formatted = await prettier.format(classResult, {
      parser: "typescript",
    });
    content.push("```ts\r\n" + formatted + "\r\n```");

    const arrayObjectExample = {
      name: "MyArrayClass",
      fields: [
        {
          name: "contacts",
          type: "array",
          fields: [
            { name: "email", type: "string" },
            { name: "phone", type: "string?" },
          ],
        },
      ],
    };

    content.push(`
Array of objects:
\`\`\`yaml
${yaml.dump(arrayObjectExample)}
\`\`\`

Arrays can hold primitive types or nested objects.
    `);

    const arrayObjectResult = globalThis.jsGenDtoClass(
      JSON.stringify(arrayObjectExample),
      { Tags: "typescript", Flags: "Anonymouse" }
    );
    const formattedArrayObj = await prettier.format(arrayObjectResult, {
      parser: "typescript",
    });
    content.push("```ts\r\n" + formattedArrayObj + "\r\n```");

    const nullableArrayExample = {
      name: "MyArrayClass",
      fields: [{ name: "nullableTags", type: "array?" }],
    };

    content.push(`
Nullable version:
\`\`\`yaml
${yaml.dump(nullableArrayExample)}
\`\`\`

Defaults to \`undefined\`, but you can assign an array or \`null\`.
    `);

    const nullableResult = globalThis.jsGenDtoClass(
      JSON.stringify(nullableArrayExample),
      { Tags: "typescript", Flags: "Anonymouse" }
    );
    const formattedNullable = await prettier.format(nullableResult, {
      parser: "typescript",
    });
    content.push("```ts\r\n" + formattedNullable + "\r\n```");

    writeFileSync(
      "../../emi-web/docs/js/emi-array-data-type.mdx",
      content.join("\r\n").trim()
    );
  });
});
