import { writeFileSync } from "fs";
import prettier from "prettier";
import { describe, it } from "vitest";
import { createInstance } from "../../../emi-npm/bin/getPublicActions";
import yaml from "js-yaml";

describe("Generate documents for boolean and boolean?", () => {
  const content: string[] = [];

  it("should init the wasm", () => {
    createInstance();
  });

  it("should add the introduction", () => {
    content.push(`
---
sidebar_position: 5
---

# Emi boolean data type

        `);
  });

  it("explains about the pure boolean, and the default value", async () => {
    const pureBoolExample = {
      name: "MyBoolClass",
      fields: [
        {
          name: "boolWithDefault",
          default: true,
          type: "bool",
        },
        {
          name: "plainBool",
          type: "bool",
        },
      ],
    };

    content.push(`
      Given the following example:
\`\`\`yaml
${yaml.dump(pureBoolExample)}
\`\`\`
      
      `);
    content.push(`
      \`bool\` and \`bool?\` map directly to JavaScript's \`boolean\` type.
      - \`bool\` fields are always required and default to \`false\` unless another default is provided.
      - \`bool?\` fields allow both \`null\` and \`undefined\`, with \`undefined\` being the default if not specified.

      Just like with strings, Emi generates a private backing field (\`#field\`) and a pair of getters and setters.
      This ensures type safety and avoids accidental runtime errors.
      
      Example with a non-nullable \`bool\`:
      `);

    const classResult = globalThis.jsGenDtoClass(
      JSON.stringify(pureBoolExample),
      {
        Tags: "typescript",
        Flags: "Anonymouse",
      }
    );

    const formatted = await prettier.format(classResult, {
      parser: "typescript",
    });

    content.push("```ts\r\n" + formatted + "\r\n```");
  });

  it("explains about the nullable booleans as well.", async () => {
    const nullableBoolExample = {
      name: "MyBoolClass",
      fields: [
        {
          name: "nullableBoolWithDefault",
          default: true,
          type: "bool?",
        },
        {
          name: "nullableBoolWithoutDefault",
          type: "bool?",
        },
      ],
    };

    content.push(`
      Given the following example:
\`\`\`yaml
${yaml.dump(nullableBoolExample)}
\`\`\`
      
      `);
    content.push(`
      With \`bool?\`, the backing field is initialised with \`undefined\`.
      You can explicitly set it to \`true\`, \`false\`, or \`null\`.
      If a default is provided, it will be respected.

      `);

    const classResult = globalThis.jsGenDtoClass(
      JSON.stringify(nullableBoolExample),
      {
        Tags: "typescript",
        Flags: "Anonymouse",
      }
    );

    const formatted = await prettier.format(classResult, {
      parser: "typescript",
    });

    content.push("```ts\r\n" + formatted + "\r\n```");
  });

  it("should write the final doc", () => {
    writeFileSync(
      "../../emi-web/docs/js/emi-boolean-data-type.mdx",
      content.join("\r\n").trim()
    );
  });
});
