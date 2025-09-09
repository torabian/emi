import { writeFileSync } from "fs";
import yaml from "js-yaml";
import prettier from "prettier";
import { PropertyDeclaration } from "ts-morph";
import { describe, expect, it } from "vitest";
import { createInstance } from "../../../emi-npm/bin/getPublicActions";
import { parseGenerated } from "../common";

const sample1 = {
  name: "emiComplexFields",
  complexes: [
    {
      name: "Money",
      location: "../some/directory/Money",
    },
  ],
  actions: [
    {
      name: "sample",
      url: "http://localhost:8081 (for test we use override)",
      method: "post",
      description:
        "Demo action that shows how complex types can be used in an action",
      out: {
        envelope: "GResponse",
        fields: [
          {
            name: "date",
            complex: "+Date",
            description: "Use javascript Date class as complex type",
          },
          {
            name: "money",
            complex: "+Money",
            description: "External money type",
          },
        ],
      },
    },
  ],
};

describe("Generate the complex fields within an action, both for built in and user defined.", () => {
  const content: string[] = [];
  let renderedClass = "";
  let files = [];

  it("should init the wasm", () => {
    createInstance();
  });

  it("should generate and write the Emi JS module", async () => {
    files = globalThis.jsGenModule(JSON.stringify(sample1), {
      Tags: "typescript",
    }) as any;
    renderedClass =
      (files.find((item: any) => item.Name === "SampleAction") as any)
        ?.ActualScript || "";
  });

  it("should have been generated a valid typescript file", () => {
    expect(typeof renderedClass).toBe("string");
  });

  it("should have been formatted with prettier", async () => {
    const formatted = await prettier.format(renderedClass, {
      parser: "typescript",
    });

    // add markdown documentation
    content.push(`# Complex Types in Emi`);
    content.push(
      `In Emi modules, you can define **complexes** as reusable types. Each complex requires a class name and a location where it should be imported from.`
    );
    content.push(
      `Instead of using primitive types like \`number\` or \`string\`, you can specify \`complex: ClassName\` inside the action field definition.`
    );
    content.push(
      `- If you prefix the complex with **"+"**, like \`+ClassName\`, Emi will automatically instantiate the value when parsing the response.`
    );
    content.push(
      `- Built-in types such as \`Date\` can be used directly as complexes.`
    );
    content.push(
      `- User-defined types such as \`Money\` need to be registered in the **complexes** array of your module.`
    );

    content.push("\n## Sample module\n");
    content.push(
      "Here you can see based on defining 2 complex fields, we are generating them as class instances"
    );
    content.push(`
Example schema:
\`\`\`yaml
${yaml.dump(sample1)}
\`\`\``);
    content.push("\n## Example Generated Code\n");
    content.push("```ts\r\n" + formatted + "\r\n```");

    content.push("\n## Notes\n");
    content.push(
      `- \`Date\` is treated as a built-in complex and is not imported.`
    );
    content.push(
      `- \`Money\` is imported from the given location and instantiated when assigned.`
    );
    content.push(
      `- The generated setters (\`setDate\`, \`setMoney\`) return \`this\` so you can chain assignments.`
    );
  });

  it("should have date setter with Date type", () => {
    const sf = parseGenerated(renderedClass);
    const setter = sf
      .getClassOrThrow("SampleActionRes")
      .getSetAccessorOrThrow("date");
    const paramType = setter.getParameters()[0].getType().getText();
    expect(paramType).toBe("Date");
  });

  it("should have #money field private", () => {
    const sf = parseGenerated(renderedClass);
    const clazz = sf.getClassOrThrow("SampleActionRes");

    const moneyProp = clazz
      .getProperties()
      .find((p) => p.getName() === "#money") as PropertyDeclaration | undefined;

    expect(moneyProp).toBeTruthy();
  });

  it("should import Money", () => {
    const sf = parseGenerated(renderedClass);
    const imports = sf
      .getImportDeclarations()
      .map((i) => i.getModuleSpecifierValue());
    expect(imports).toContain("../some/directory/Money");
  });

  it("should not import Date", () => {
    const sf = parseGenerated(renderedClass);
    const imports = sf.getImportDeclarations().map((i) => i.getText());
    expect(imports.join("\n")).not.toMatch(/Date/);
  });

  it("should have setDate method returning this", () => {
    const sf = parseGenerated(renderedClass);
    const method = sf
      .getClassOrThrow("SampleActionRes")
      .getMethodOrThrow("setDate");
    expect(method.getReturnType().getText()).toContain("this");
  });

  it("should have setMoney method returning this", () => {
    const sf = parseGenerated(renderedClass);
    const method = sf
      .getClassOrThrow("SampleActionRes")
      .getMethodOrThrow("setMoney");
    expect(method.getReturnType().getText()).toContain("this");
  });

  it("should write the final doc", () => {
    writeFileSync(
      "../../emi-web/docs/js/emi-complex-types.mdx",
      content.join("\r\n").trim()
    );
  });

  it("should have SampleAction class with static URL property", () => {
    const sf = parseGenerated(renderedClass);
    const clazz = sf.getClassOrThrow("SampleAction");
    const urlProp = clazz.getPropertyOrThrow("URL");
    expect(urlProp.isStatic()).toBe(true);
  });

  it("should have Fetch$ static arrow function returning a Promise", () => {
    const sf = parseGenerated(renderedClass);
    const clazz = sf.getClassOrThrow("SampleAction");
    const fetchProp = clazz.getPropertyOrThrow("Fetch$");
    expect(fetchProp.isStatic()).toBe(true);

    const init = fetchProp.getInitializerOrThrow();
    expect(init.getText()).toMatch(/async/);
  });

  it("should import GResponse from envelopes", () => {
    const sf = parseGenerated(renderedClass);
    const imports = sf
      .getImportDeclarations()
      .map((i) => i.getModuleSpecifierValue());
    expect(imports).toContain("./sdk/envelopes/index");
  });

  it("should declare abstract SampleActionResFactory", () => {
    const sf = parseGenerated(renderedClass);
    const clazz = sf.getClassOrThrow("SampleActionResFactory");
    expect(clazz.isAbstract()).toBe(true);
  });
});
