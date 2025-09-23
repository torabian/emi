import { writeFileSync } from "fs";
import prettier from "prettier";
import { describe, expect, it } from "vitest";
import { createInstance } from "../../../emi-npm/bin/getPublicActions";
import yaml from "js-yaml";
import { AutoInitClassDto } from "../test-artifacts/auto-init.dto";

describe("Objects auto init should work perfectly fine.", () => {
  const content: string[] = [];
  const arrayExample = {
    name: "AutoInitClass",
    fields: [
      {
        name: "object1",
        type: "object",
        description: "This field will be always initialised",
        fields: [
          {
            name: "object2",
            type: "object",
            description: "This field also will be initialised, always.",
            fields: [
              {
                name: "contacts",
                type: "array",
                description: "This field will be always an array",
                fields: [
                  {
                    name: "email",
                    type: "string",
                    default: "emi-compiler@emi-compiler.com",
                  },
                  { name: "phone", type: "string?" },
                ],
              },
            ],
          },
        ],
      },
    ],
  };

  it("should init the wasm", () => {
    createInstance();
  });

  it("should add the introduction", () => {
    content.push(`
---
sidebar_position: 3
---

# Emi Objects Auto-Init Feature

In the DTO world, it often happens that an API promises to return a string, object, or array for a specific property. TypeScript can enforce this at the type level, but the main issue with most codegen tools is that they trust the definitions too much.

Over time, APIs change, and the expected data type may not match what actually comes back. In strongly typed languages, this problem is usually handled automatically. For example, in Go, structs and primitives are always initialized, regardless of the input. Some languages don’t even allow uninitialized values.

TypeScript tries to address this, but only at compile time—it doesn’t validate data at runtime. That means code like \`const data = body.response.substring()\` will break if the response is suddenly a number instead of a string.

And these issues *do* happen. Maybe not on day one, but eventually, as the codebase grows and multiple people contribute, inconsistencies creep in.


## How Emi-Generated DTOs Solve This Problem

When you define an object in Emi, it generates both a TypeScript type and a class. When an action returns an instance of that class, all non-nullable fields are automatically initialized—even if they are arrays, objects, or nested DTOs.

This ensures your code is always reliable. You don’t need to sprinkle \`?\`, \`??\`, or \`typeof x === 'string'\` checks everywhere.

For example, when generating an Emi class with nested objects, Emi will automatically initialize everything in the response, so you can safely work with the data without extra guards.


Now, this is much more reliable code instead of just typing them, you almost never get undefined. But still, writing that manually
is a hassle, and type checking and other tools might be lost. Now, let's take a look how this could be generated using Emi js compiler:

 
\`\`\`yaml
${yaml.dump(arrayExample)}
\`\`\`

And Emi compiler would generate such shopisticated details (content is a bit long but worth to take a look at it.)
Emi compiler generates ts type and, class, with full getter, setters and validators for each field.


        `);
  });

  it("compile the Emi fields array and show it as json", async () => {
    const classResult = globalThis.jsGenDtoClass(JSON.stringify(arrayExample), {
      Tags: "typescript",
      Flags: "AutoInitDto",
    });

    const formatted = await prettier.format(classResult, {
      parser: "typescript",
    });

    content.push("```ts\r\n" + formatted + "\r\n```");

    writeFileSync(__dirname + "/../test-artifacts/auto-init.dto.ts", formatted);
  });

  it("Test 1, creating an empty class out of it, needs to have all the path", async () => {
    let m = new AutoInitClassDto();

    // By default, the contents need to be an array, regardless.
    expect(m.object1.object2.contacts).to.be.an("array");

    m = AutoInitClassDto.from({ object1: { object2: { contacts: [] } } });
    expect(m.object1.object2.contacts).to.be.an("array");

    // Should not be able to set 22 as an array!
    m = AutoInitClassDto.from({
      object1: { object2: { contacts: 22 as any } },
    });
    expect(m.object1.object2.contacts).to.be.an("array");

    m = AutoInitClassDto.from({
      object1: { object2: { contacts: [{ email: "hass" }] } },
    });

    expect(m.object1.object2.contacts[0].email).toEqual("hass");

    m = AutoInitClassDto.from({
      object1: { object2: { contacts: [{} as any] } },
    });

    expect(m.object1.object2.contacts[0].email).toEqual(
      "emi-compiler@emi-compiler.com"
    );

    // Phone needs to be exactly as undefined.
    expect(typeof m.object1.object2.contacts[0].phone).toEqual("undefined");
  });

  /// Last step is to write the document down
  it("should write the final doc", () => {
    writeFileSync(
      "../../emi-web/docs/js/emi-fields-auto-init.mdx",
      content.join("\r\n").trim()
    );
  });
});
