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

  /// Last step is to write the document down
  it("should write the final doc", () => {
    writeFileSync(
      "../../emi-web/docs/js/emi-static-fields.mdx",
      content.join("\r\n").trim()
    );
  });
});
