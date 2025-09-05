import { describe, it, expect } from "vitest";
import { runEmiActionTs } from "../../common";
import { writeFileSync } from "fs";
import yaml from "js-yaml";

describe("Action headers class needs to be generated correctly.", () => {
  const content: string[] = [];

  const action = [
    { name: "accept-language", type: "string" },
    { name: "authroization", type: "string" },
    { name: "cache-time", type: "int64" },
  ];

  const { source, resp } = runEmiActionTs("jsGenActionHeaders", action, {
    Flags: "SampleHeaders",
  });

  it("generate int doc", async () => {
    content.push(`
---
sidebar_position: 3
---

# Emi Headers class generator

In the context of http requests, there is always request and response headers, which is a key pair
definition. Modern javascript standard is using Headers class, and Emi generates the header classes,
by extending that class, and adding extra functions.

    `);
  });

  it("should document the initial yaml", () => {
    content.push(`
          Example schema:
\`\`\`yaml
${yaml.dump(action)}
\`\`\`

Would generate the following code:

\`\`\`ts
${resp}
\`\`\`

          `);
  });

  it("should have created only a single class representing the header", () => {
    expect(source.getClasses()).toHaveLength(1);
  });

  it("should have action name uppercase, and ReqHeader affix", () => {
    const class$ = source.getClasses()[0].getName();
    expect(class$).toBe(`SampleHeaders`);
  });

  it("should extend the standard Headers class", () => {
    const base = source.getClasses()[0].getExtends();
    expect(base).not.toBeNull();
    expect(base?.getText()).toBe("Headers");
  });

  it("the string field, needs to have getter, and setter, with appropirate types", () => {
    const getter = source.getClasses()[0].getMethod("getAcceptLanguage");
    const setter = source.getClasses()[0].getMethod("setAcceptLanguage");
    expect(getter).toBeDefined();
    expect(setter).toBeDefined();

    // jsdoc check
    const getterDocs = getter!
      .getJsDocs()
      .map((d) => d.getInnerText())
      .join("\n");
    expect(getterDocs).toContain("@returns { string | null }");

    const setterDocs = setter!
      .getJsDocs()
      .map((d) => d.getInnerText())
      .join("\n");
    expect(setterDocs).toContain("@param { string | null } value");

    // @todo Test cases need to be here, to test the return type of the typescript getter,
    // and the value type of setter are correct

    // // return type / param type
    // expect(getter!.getReturnType().getText()).toBe("any"); // compiler sees `any` because of #getTyped, so rely on jsdoc instead
    // expect(setter!.getParameters()[0].getName()).toBe("value");

    // body checks
    expect(getter!.getBodyText()).toContain(
      "this.#getTyped('accept-language', 'string | null')"
    );
    expect(setter!.getBodyText()).toContain(
      "this.set('accept-language', value)"
    );
  });

  it("should define #getTyped and toObject methods", () => {
    const cls = source.getClasses()[0];
    const members = cls.getMembers().map((m) => (m as any).getName());
    expect(members).toContain("#getTyped");
    expect(cls.getMethod("toObject")).toBeDefined();
  });

  it("should write the documentation", () => {
    writeFileSync(
      "../../emi-web/docs/js/emi-headers-class.mdx",
      content.join("\r\n").trim()
    );
  });
});
