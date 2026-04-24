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

# Referencing other dto, collection and one

        `);
  });
 
  it("Referencing the other dtos", async () => {
 
    const objectAndArrays = {
      name: "Catalog",
      fields: [
        {
          name: "person",
          type: "one",
          target: "PersonDto"
        },
        {
          name: "personOptional",
          type: "one?",
          target: "PersonDto"
        },
        {
          name: "people",
          type: "collection",
          target: "PersonDto"
        },
        {
          name: "peopleOptional",
          type: "collection?",
          target: "PersonDto"
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
      "../../emi-web/docs/golang/emi-golang-referencing-other-dtos.mdx",
      content.join("\r\n").trim(),
    );
  });

});
