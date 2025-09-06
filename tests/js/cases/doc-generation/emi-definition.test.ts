import { describe, it, expect } from "vitest";
import { createInstance } from "../../../../emi-npm/bin/getPublicActions";
import { existsSync, writeFileSync } from "fs";
import path from "path";

var schema: any = {};

const coreDefinition = "lib/core/Emi.go";

describe("Generate the emi introduction documentation", () => {
  const content: string[] = [];

  it("should init the wasm", () => {
    createInstance();
  });

  it("golang definition file should exists on disk", () => {
    expect(existsSync(path.join("../..", coreDefinition))).toBe(true);
  });

  it("should add the introduction", () => {
    content.push(`
---
sidebar_position: 3
---

# Emi definitions

In this document, we discuss the emi definition, and features it has. Emi definition can be defiend via yaml, or json,
and doesn't have a special language. It can have a different series of features in it, from actions, dtos, entities.
In an absract level, these are just definitions, which different sub compilers, for example js, would create
different target codes (React, Vanilajs, Nest.js) based on the tags provided.

Emi definitions aim to prevent having specific target language code, so it could be compiled into as many as languages over time.
You will find all structs in https://github.com/torabian/emi/tree/main/${coreDefinition} 
        `);
  });

  it("should generate the schema directly from wasm", () => {
    schema = JSON.parse(globalThis.genEmiSpec("", {}));
  });

  it("Should document an example of Emi module", () => {
    {
      content.push("## Simple emi module");
      content.push(
        "Here, you can see a minimal module written in Emi, which defines get method to fetch a json placeholder."
      );
      content.push(
        `\`\`\`yaml
name: sampleModule
actions:
  - name: getSinglePost
    url: https://jsonplaceholder.typicode.com/posts/:id
    cliName: get-single-post
    method: get
    description: Get's an specific post from the endpoint
    out:
      fields:
        - name: userId
          type: int64
        - name: id
          type: int64
        - name: title
          type: string
        - name: body
          type: string
\`\`\`
        `.trim()
      );
    }
  });

  it("should be able to document the Emi main function", () => {
    {
      content.push("## Emi main features.");
      content.push(
        "On the yaml root, you can set different features of a the module, such as entities, dtos, actions, and many more."
      );
    }
    {
      // In this block we make a table in markdown explaining the features of the Emi module.
      const items = schema.definitions.Emi.properties;
      const header = `| Property | Type | Description |\n|----------|------|-------------|`;
      const rows = Object.entries(items)
        .map(([name, def]: any) => {
          return `| \`${name}\` | \`${def.type}\` | ${def.description} |`;
        })
        .join("\n");

      const doc = [header, rows].join("\n");
      content.push(doc);
    }
  });

  /// Last step is to write the document down
  it("should write the final doc", () => {
    writeFileSync(
      "../../emi-web/docs/emi-definitions.mdx",
      content.join("\r\n").trim()
    );
  });
});
