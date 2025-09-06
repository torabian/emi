import { writeFileSync } from "fs";
import prettier from "prettier";
import { describe, it } from "vitest";
import { createInstance } from "../../../emi-npm/bin/getPublicActions";
import yaml from "js-yaml";

describe("Generate documents for string and strings?", () => {
  const content: string[] = [];

  it("should init the wasm", () => {
    createInstance();
  });

  it("should add the introduction", () => {
    content.push(`
---
sidebar_position: 4
---

# Emi string data type
 
        `);
  });

  it("explains about the pure string, and the default value", async () => {
    const pureStringExample = {
      name: "MyStringClass",
      fields: [
        {
          name: "stringWithDefault",
          default: "With default value",
          type: "string",
        },
        {
          name: "emptyString",
          type: "string",
        },
      ],
    };

    content.push(`
      Given the following example:
\`\`\`yaml
${yaml.dump(pureStringExample)}
\`\`\`
      
      `);
    content.push(`
      \`string\` and \`string?\` types will be converted into the javascript counter part with a difference that ? allows the null value, and by default is \`undefined\`.
      As a general pattern the field private value (#field) would be created, and based on that, there will be getters and setters.

      First let's take a look on the \`string\`, and see how it would prevent the crashes by generating the typesafe code:

      `);

    const classResult = globalThis.jsGenDtoClass(
      JSON.stringify(pureStringExample),
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

  it("explains about the nullable strings as well.", async () => {
    const pureStringExample = {
      name: "MyStringClass",
      fields: [
        {
          name: "nullableStringWithDefault",
          default: "With default value",
          type: "string?",
        },
        {
          name: "nullableWithoutDefault",
          type: "string?",
        },
      ],
    };

    content.push(`
      Given the following example:
\`\`\`yaml
${yaml.dump(pureStringExample)}
\`\`\`
      
      `);
    content.push(`
      In nullable scenario, the field is initialised with 'undefined' and allows null value as well.
      If default provided, the default value will be placed.

      `);

    const classResult = globalThis.jsGenDtoClass(
      JSON.stringify(pureStringExample),
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

  /// Last step is to write the document down
  it("should write the final doc", () => {
    writeFileSync(
      "../../emi-web/docs/js/emi-string-data-type.mdx",
      content.join("\r\n").trim()
    );
  });
});
