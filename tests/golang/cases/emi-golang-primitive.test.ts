import { writeFileSync } from "fs";
import yaml from "js-yaml";
import { describe, it } from "vitest";
import { createInstance } from "../../../emi-npm/bin/getPublicActions";
import { formatGoCode } from "../common";

describe("Generates DTO for all possible data types in a single go.", () => {
  let content: string[] = [];

  it("should init the wasm", () => {
    createInstance();
  });

  it("should add the introduction", () => {
    content.push(`
---
sidebar_position: 5
---

# Golang primitive data types.

Emi golang compiler converts the Emi data types into appropriate keywords and structs 
in Golang. Emi types are fundamentally inspired by golang data types anyway,
but still there might be some mismatches added over time. Besides for nested objects,
new structs needed to be created.

## Default values in Go version

Golang pure structs (unlike classes in js) cannot have a custom default value without initialization.
Emi might generate helper function to solve this issue, but at the written time 
of this document is not available.

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
          type: "bool?",
        },
        {
          name: "intField",
          type: "int",
        },
        {
          name: "int32Field",
          type: "int32",
        },
        {
          name: "int64Field",
          type: "int64",
        },
        {
          name: "intFieldNullable",
          type: "int?",
        },
        {
          name: "int32FieldNullable",
          type: "int32?",
        },
        {
          name: "int64FieldNullable",
          type: "int64?",
        },
        {
          name: "stringField",
          type: "string",
        },
        {
          name: "stringFieldNullable",
          type: "string?",
        },
        {
          name: "anyField",
          type: "any",
        },
        {
          name: "inner",
          type: "object",
          fields: [
            {
              name: "sliceField",
              type: "slice",
              primitive: "string",
            },
            {
              name: "sliceFieldOptional",
              type: "slice?",
              primitive: "int",
            },
          ],
        },
      ],
    };

    content.push(`
      Primitive support in Golang is standard as \`bool\` and \`bool?\` emi types. \`bool\` resolves the primitive 
      counterpart in golang without modification, whearas the nullable data types such as
      \`bool?\` will be wrapped around Nullable[bool]
      similar to all other primitives and structs which are nullable.

      Emi doesn't generate pointer elements in Golang, and it's support has been removed many years ago in Fireback.
\`\`\`yaml
${yaml.dump(pureBoolExample)}
\`\`\`
    
Possible result:


      `);

    let classResult = globalThis.goGenObject(JSON.stringify(pureBoolExample), {
      Flags: JSON.stringify({ pkg: "datatypes" }),
    });

    classResult = await formatGoCode(classResult);
    content.push("```go\r\n" + classResult + "\r\n```");
  });

  it("should write the final doc", () => {
    writeFileSync(
      "../../emi-web/docs/golang/emi-primitive.mdx",
      content.join("\r\n").trim(),
    );
  });
});
