import { writeFileSync } from "fs";
import yaml from "js-yaml";
import prettier from "prettier";
import { describe, it } from "vitest";
import { createInstance } from "../../../emi-npm/bin/getPublicActions";
var schema: any = {};

describe("Generate the Emi static fields generation", () => {
  const content: string[] = [];

  const exampleModel = [
    { name: "firstName", type: "string" },
    {
      name: "loginHistory",
      type: "array",
      fields: [
        { name: "time", type: "string" },
        {
          name: "data",
          type: "object",
          fields: [{ name: "ipAddress", type: "string" }],
        },
      ],
    },
  ];

  it("should init the wasm", () => {
    createInstance();
  });

  it("should add the introduction", () => {
    content.push(`
---
sidebar_position: 4
---

# Static fields path generation

In Javascript world, often you need to set a variable using a json path, for example using 'lodash' or 'formik' libraries. This is very common way to set value of a variable, deep inside an object, array, or combination.

Let's assume the following code:


\`\`\`javascript
const form = {
  firstName: 'Ali',
  loginHistory: [
    {
      time: new Date(),
      data: {
        ipAddress: '0.0.0.0',
      } 
    }
  ]
}
\`\`\`

Now, let's call the \`lodash.set(field, value)\` function on the following code:

\`\`\`javascript
import { set } from 'lodash';

set(form, 'loginHistory[0].data.ipAddress', '192.168.1.1')
\`\`\`

*A side note is you might ask why not set 'form.loginHistory[0].data.ipAddress = "192.168.1.1' instead,
in javascript world it's very common to use the path selector, in forms context, or when you do not have access
to whole flow.*


## Issue with the code

As we called \`set\` function with string path finder, this path might be changed over time, and 
javascript compiler might silently continue to set that form, and it's hard to debug.
If the \`form.loginHistory\` becomes \`form.loginHistories\`, then still some part of code still
tries to set the address in a wrong location.

## How Emi static generation would solve this?

When defining a model with Emi, compiler generates a \`static Fields = ... \` statement in each sub class,
and automatically connect the path in depth. Let's assume the following schema, this time written in Emi.

 
\`\`\`yaml
${yaml.dump(exampleModel)}
\`\`\`

And as you can see in the generated code there are \`static get Fields()\` statements, which instead
of setting string, now you can:

\`\`\`javascript
import { set } from 'lodash';

set(form, at(Anonymouse.Fields.loginHistory.data.ipAddress, 0), '192.168.1.1')
\`\`\`

Now in case the model changes, compiler would complain, and there is zero chance to modify no existing fields.
Few notes to remember:

- When a field is object or array (and all similar types), the field itself appears twice, first with $ sign,
which is representing the string field, and another time as the name of field, which would reference the other
class representing the sub class, or original reference in case of Dto or Entity.
- In javascript core library, there is a \`at\` function, which would help to resolve the array string paths. You can 
pass number as many as you want, and it would replace [:i] statements to have proper index selected.

          `);
  });

  it("compile the Emi fields array and show it as json", async () => {
    const classResult = globalThis.jsGenObject(JSON.stringify(exampleModel), {
      Flags: "Anonymouse",
    });

    const formatted = await prettier.format(classResult, {
      parser: "typescript",
    });

    content.push("```ts\r\n" + formatted + "\r\n```");
  });

  //   it("should generate the schema directly from wasm", () => {
  //     schema = JSON.parse(globalThis.genEmiSpec("", {}));
  //   });

  //   it("Should document an example of Emi module", () => {
  //     {
  //       content.push("## Defining actions overview");
  //       content.push(
  //         "As mentioned in the general definitions, actions is a part of module, which is an array, and allows developer to define multiple actions, as items. Each" +
  //           " action must have a name, and that's basically enough for it to be compiled by different sub compilers in Emi project."
  //       );
  //       content.push(
  //         `\`\`\`yaml
  // name: sampleModule
  // actions:
  //   - name: getSinglePost
  //     url: https://jsonplaceholder.typicode.com/posts/:id
  //     cliName: get-single-post
  //     method: get
  //     description: Get's an specific post from the endpoint
  //     out:
  //       fields:
  //         - name: userId
  //           type: int64
  //         - name: id
  //           type: int64
  //         - name: title
  //           type: string
  //         - name: body
  //           type: string
  //   - name: sampleSse
  //     url: http://localhost:3000/stream
  //     method: sse
  //     description: SSE Sample
  //     out:
  //       fields:
  //         - name: message
  //           type: string
  //   - name: webSocketOrgEcho
  //     url: "wss://echo.websocket.org/.ws"
  //     method: reactive
  //     description: Websocket.org eco server, to send a json and recieve back
  //     in:
  //       fields:
  //         - name: firstName
  //           type: string
  //         - name: lastName
  //           type: string
  //     out:
  //       fields:
  //         - name: lastName
  //           type: string

  // \`\`\`
  //         `.trim()
  //       );
  //     }
  //   });

  //   it("should be able to document the Emi actions function", () => {
  //     {
  //       content.push("## Actions properties");
  //       content.push(
  //         "Each action within a module needs to have a camel case name, and needs to be unique. The other properties are optional. Nevertheless, 'in', 'out', 'method' and 'url'" +
  //           " are the most common properties of an action which you'll set as a developer."
  //       );

  //       content.push(
  //         `Depending on your use case, you might not need all features, or they might not be used by the sub compilers.`
  //       );
  //     }
  //     {
  //       // In this block we make a table in markdown explaining the features of the Emi module.
  //       const items = schema.definitions.EmiAction.properties;
  //       const header = `| Property | Type | Description |\n|----------|------|-------------|`;
  //       const rows = Object.entries(items)
  //         .map(([name, def]: any) => {
  //           return `| \`${name}\` | \`${def.type}\` | ${def.description} |`;
  //         })
  //         .join("\n");

  //       const doc = [header, rows].join("\n");
  //       content.push(doc);
  //     }
  //   });

  //   it("should be able to mention the allowed different methods of the action which is allowed", () => {
  //     {
  //       content.push("## Action methods");
  //       content.push(
  //         `Emi allows for standard http methods, as well as some extra which will be converted to different protocols.`
  //       );
  //     }
  //     {
  //       // In this block we make a table in markdown explaining the features of the Emi module.
  //       const items = schema.definitions.EmiAction.properties.method.enum;

  //       expect(items).toContain("reactive");
  //       expect(items).toContain("post");
  //       expect(items).toContain("get");
  //       expect(items).toContain("delete");
  //       expect(items).toContain("patch");
  //       expect(items).toContain("put");
  //       expect(items).toContain("sse");

  //       const rows = items.map((x) => `\`${x}\``).join(", ");

  //       content.push(
  //         `Currently supported the actions and methods are the following: ${rows}. Calling for non-existing method, would resolve to http with the non-standard method.`
  //       );
  //     }
  //   });

  /// Last step is to write the document down
  it("should write the final doc", () => {
    writeFileSync(
      "../../emi-web/docs/js/emi-static-fields.mdx",
      content.join("\r\n").trim()
    );
  });
});
