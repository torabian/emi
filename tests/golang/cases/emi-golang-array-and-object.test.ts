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
sidebar_position: 4
---

# Golang object and array

        `);
  });
 
  it("array an object", async () => {
 
    const objectAndArrays = {
      name: "Catalog",
      fields: [
        {
          name: "person",
          type: "object",
          fields: [
            {
              name: "firstName",
              type: "string"
            },
            {
              name: "lastName",
              type: "string"
            }
          ]
        },
        {
          name: "people",
          type: "array",
          fields: [
            {
              name: "firstName",
              type: "string"
            },
            {
              name: "lastName",
              type: "string"
            },
            {
              name: "parent",
              type: "one",
              target: "PersonDto"
            }
          ]
        },   
        {
          name: "personOptional",
          type: "object?",
          fields: [
            {
              name: "firstName",
              type: "string"
            },
            {
              name: "lastName",
              type: "string"
            }
          ]
        },
        {
          name: "peopleOptional",
          type: "array?",
          fields: [
            {
              name: "firstName",
              type: "string"
            },
            {
              name: "lastName",
              type: "string"
            },
            {
              name: "parent",
              type: "one",
              target: "PersonDto"
            }
          ]
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
${yaml.dump(objectAndArrays)}
\`\`\`
    
Possible result:


      `);
    
    let classResult = globalThis.goGenObject(
      JSON.stringify(objectAndArrays),
      {
        Flags: JSON.stringify({"pkg": "datatypes"})
      },
    );

    classResult= await formatGoCode(classResult)
    content.push("```go\r\n" + classResult + "\r\n```");
  });
 
it("should write the final doc", () => {
    writeFileSync(
      "../../emi-web/docs/golang/emi-array-and-objects.mdx",
      content.join("\r\n").trim(),
    );
  });

});
