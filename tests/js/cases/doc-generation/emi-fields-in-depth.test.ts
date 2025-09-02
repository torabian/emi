import { writeFileSync } from "fs";
import prettier from "prettier";
import { describe, it } from "vitest";
import { createInstance } from "../../../../emi-npm/bin/getPublicActions";
import yaml from "js-yaml";

describe("Generate the dto class and explaining different fields created for getter and setter.", () => {
  const content: string[] = [];

  it("should init the wasm", () => {
    createInstance();
  });

  it("should add the introduction", () => {
    content.push(`
---
sidebar_position: 4
---

# Class fields in depth

Emi generates fields into classes, but each field is getting setters and getters to ensure 100% type safety upon usage.
We achive this by closing all options to modify the object directly, and setter functions which would strictly check the values being passed.

Field types are categorized into the following groups, despite their Emi types:

* Strings (nullable?)
* Numbers (nullable?)
* Class instances (object, one, etc.) (always nullable)
* Boolean (nullable?)
* Array of instances (not nullable)
* Array of primitives (not nullable)


In this document, we are going to discuss how such fields would become available in javascript counter part, and how codegen enforces
the type-safety to a next level.
        `);
  });

  /// Last step is to write the document down
  it("should write the final doc", () => {
    writeFileSync(
      "../../emi-web/docs/js/emi-fields-in-depth.mdx",
      content.join("\r\n").trim()
    );
  });
});
