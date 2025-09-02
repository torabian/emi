import { writeFileSync } from "fs";
import prettier from "prettier";
import { describe, it } from "vitest";
import { createInstance } from "../../../../emi-npm/bin/getPublicActions";
import yaml from "js-yaml";

describe("Generate documents for int / float64 types", () => {
  const contentInt: string[] = [];
  const contentFloat: string[] = [];

  it("should init the wasm", () => {
    createInstance();
  });

  //
  // INT DOCUMENT
  //
  it("generate int doc", async () => {
    contentInt.push(`
---
sidebar_position: 6
---

# Emi integer data type
    `);

    const pureIntExample = {
      name: "MyIntClass",
      fields: [
        { name: "intWithDefault", default: 42, type: "int" },
        { name: "plainInt", type: "int" },
      ],
    };

    contentInt.push(`
      Example schema:
\`\`\`yaml
${yaml.dump(pureIntExample)}
\`\`\`
      
- \`int\` maps to JavaScript \`number\`, limited to safe 32-bit values.  
- If no default provided, it defaults to \`0\`.  
- \`int?\` allows \`null\` and \`undefined\`, with \`undefined\` as default.  
    `);

    const classResult = globalThis.jsGenDtoClass(
      JSON.stringify(pureIntExample),
      { Tags: "typescript", Flags: "Anonymouse" }
    );
    const formatted = await prettier.format(classResult, {
      parser: "typescript",
    });
    contentInt.push("```ts\r\n" + formatted + "\r\n```");

    const nullableIntExample = {
      name: "MyIntClass",
      fields: [
        { name: "nullableIntWithDefault", default: 7, type: "int?" },
        { name: "nullableIntWithoutDefault", type: "int?" },
      ],
    };

    contentInt.push(`
      Nullable version:
\`\`\`yaml
${yaml.dump(nullableIntExample)}
\`\`\`

Defaults to \`undefined\`, but you can assign \`0\`, any number, or \`null\`.
    `);

    const nullableResult = globalThis.jsGenDtoClass(
      JSON.stringify(nullableIntExample),
      { Tags: "typescript", Flags: "Anonymouse" }
    );
    const formattedNullable = await prettier.format(nullableResult, {
      parser: "typescript",
    });
    contentInt.push("```ts\r\n" + formattedNullable + "\r\n```");

    contentInt.push(`
> Besides \`int\` and \`int?\`, Emi also supports \`int32\`, \`int32?\`, \`int64\`, and \`int64?\` types.
    `);

    writeFileSync(
      "../../emi-web/docs/js/emi-int-data-type.mdx",
      contentInt.join("\r\n").trim()
    );
  });

  //
  // FLOAT64 DOCUMENT
  //
  it("generate float64 doc", async () => {
    contentFloat.push(`
---
sidebar_position: 7
---

# Emi float64 data type
    `);

    const pureFloatExample = {
      name: "MyFloatClass",
      fields: [
        { name: "floatWithDefault", default: 3.14, type: "float64" },
        { name: "plainFloat", type: "float64" },
      ],
    };

    contentFloat.push(`
      Example schema:
\`\`\`yaml
${yaml.dump(pureFloatExample)}
\`\`\`
      
- \`float64\` maps to JavaScript \`number\` with double precision.  
- Default is \`0.0\` if not specified.  
- \`float64?\` allows \`null\` and \`undefined\`, defaulting to \`undefined\`.  
    `);

    const classResult = globalThis.jsGenDtoClass(
      JSON.stringify(pureFloatExample),
      { Tags: "typescript", Flags: "Anonymouse" }
    );
    const formatted = await prettier.format(classResult, {
      parser: "typescript",
    });
    contentFloat.push("```ts\r\n" + formatted + "\r\n```");

    const nullableFloatExample = {
      name: "MyFloatClass",
      fields: [
        { name: "nullableFloatWithDefault", default: 1.23, type: "float64?" },
        { name: "nullableFloatWithoutDefault", type: "float64?" },
      ],
    };

    contentFloat.push(`
      Nullable version:
\`\`\`yaml
${yaml.dump(nullableFloatExample)}
\`\`\`

Defaults to \`undefined\`, but you can assign any float value or \`null\`.
    `);

    const nullableResult = globalThis.jsGenDtoClass(
      JSON.stringify(nullableFloatExample),
      { Tags: "typescript", Flags: "Anonymouse" }
    );
    const formattedNullable = await prettier.format(nullableResult, {
      parser: "typescript",
    });
    contentFloat.push("```ts\r\n" + formattedNullable + "\r\n```");

    contentFloat.push(`
> Emi also supports \`float32\`, but \`float64\` is the default for decimals.
    `);

    writeFileSync(
      "../../emi-web/docs/js/emi-float64-data-type.mdx",
      contentFloat.join("\r\n").trim()
    );
  });
});
