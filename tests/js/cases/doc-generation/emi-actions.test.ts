import { describe, it, expect } from "vitest";
import { createInstance } from "../../../../emi-npm/bin/getPublicActions";
import { writeFileSync } from "fs";

var schema: any = {};

describe("Emi generate the web socket (reactive) hooks", () => {
  const content: string[] = [];

  it("should init the wasm", () => {
    createInstance();
  });

  it("should add the introduction", () => {
    content.push(`
---
sidebar_position: 4
---

# Emi Actions

Actions is a crucial part of almost every software. Actions in empty basically functions, which get an input, and return output. They can be correspond to 'controllers'
in traditional API development, each action is basically an endpoint. Actions also can be called from other contexts than http, for example from cli.

For most cases, actions help you to define 'get', 'post' ... actions, as well as 'reactive' which is a websocket, and 'sse', which stands for server side event.

        `);
  });

  it("should generate the schema directly from wasm", () => {
    schema = JSON.parse(globalThis.genEmiSpec("", {}));
  });

  it("Should document an example of Emi module", () => {
    {
      content.push("## Defining actions overview");
      content.push(
        "As mentioned in the general definitions, actions is a part of module, which is an array, and allows developer to define multiple actions, as items. Each" +
          " action must have a name, and that's basically enough for it to be compiled by different sub compilers in Emi project."
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
  - name: sampleSse
    url: http://localhost:3000/stream
    method: sse
    description: SSE Sample
    out:
      fields:
        - name: message
          type: string
  - name: webSocketOrgEcho
    url: "wss://echo.websocket.org/.ws"
    method: reactive
    description: Websocket.org eco server, to send a json and recieve back
    in:
      fields:
        - name: firstName
          type: string
        - name: lastName
          type: string
    out:
      fields:
        - name: lastName
          type: string

\`\`\`
        `.trim()
      );
    }
  });

  it("should be able to document the Emi actions function", () => {
    {
      content.push("## Actions properties");
      content.push(
        "Each action within a module needs to have a camel case name, and needs to be unique. The other properties are optional. Nevertheless, 'in', 'out', 'method' and 'url'" +
          " are the most common properties of an action which you'll set as a developer."
      );

      content.push(
        `Depending on your use case, you might not need all features, or they might not be used by the sub compilers.`
      );
    }
    {
      // In this block we make a table in markdown explaining the features of the Emi module.
      const items = schema.definitions.EmiAction.properties;
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

  it("should be able to mention the allowed different methods of the action which is allowed", () => {
    {
      content.push("## Action methods");
      content.push(
        `Emi allows for standard http methods, as well as some extra which will be converted to different protocols.`
      );
    }
    {
      // In this block we make a table in markdown explaining the features of the Emi module.
      const items = schema.definitions.EmiAction.properties.method.enum;

      expect(items).toContain("reactive");
      expect(items).toContain("post");
      expect(items).toContain("get");
      expect(items).toContain("delete");
      expect(items).toContain("patch");
      expect(items).toContain("put");
      expect(items).toContain("sse");

      const rows = items.map((x) => `\`${x}\``).join(", ");

      content.push(
        `Currently supported the actions and methods are the following: ${rows}. Calling for non-existing method, would resolve to http with the non-standard method.`
      );
    }
  });

  /// Last step is to write the document down
  it("should write the final doc", () => {
    writeFileSync(
      "../../emi-web/docs/emi-actions.mdx",
      content.join("\r\n").trim()
    );
  });
});
