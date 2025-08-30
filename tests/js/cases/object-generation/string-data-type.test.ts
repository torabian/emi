/**
 * Tests the string and string? data type.
 * Emi supports both string and nullable string
 * definition, and default value forced by compiler.
 * We need to test all cases.
 */
import { describe, it, expect } from "vitest";
import { runEmiActionTs } from "../../common";
import { MethodDeclaration, PropertyDeclaration } from "ts-morph";
import { writeFileSync } from "fs";
import path from "path";

export function getJsDoc(
  node: PropertyDeclaration | MethodDeclaration
): string {
  const jsDocs = node.getJsDocs();

  if (jsDocs.length === 0) return "no jsdoc";

  return jsDocs.map((d) => d.getDescription() || "").join("\n");
}

describe("Generating different data types from fields need to be working fine.", () => {
  const fieldsMap = {
    stringField: {
      type: "string",
      name: "stringField",
      description:
        "This is a pure string field, there for never can be null, and by default needs to be empty string",
      default: undefined,
    },
    stringFieldWithValue: {
      type: "string",
      default: "testvalue",
      name: "stringFieldWithValue",
      description:
        "Pure string field, but with an intial string value, and never can be undefined or null",
    },
    nullableStringField: {
      type: "string?",
      name: "nullableStringField",
      default: undefined,
      description:
        "Nullable string field. Can be undefined, or set to null to indicate intentional emptiness",
    },
    nullableStringFieldWithValue: {
      type: "string?",
      default: "stringvalue",
      name: "nullableStringFieldWithValue",
      description:
        "Nullable string field, can be undefined or null, but with an initial value",
    },
  };

  const fields = Object.keys(fieldsMap).map((key) => ({
    name: key,
    ...fieldsMap[key],
  }));

  const { source, resp } = runEmiActionTs("jsGenObject", fields, {
    Flags: "Anonymouse",
    Tags: "react,typescript",
  });

  writeFileSync(path.join(__dirname, "string-data-type.ts"), resp);

  it("should have generated a class named Anonymouse", () => {
    expect(source.getClasses().map((c) => c.getName())).toContain("Anonymouse");
  });

  it("should have generated a type named AnonymouseType", () => {
    expect(source.getTypeAliases().map((t) => t.getName())).toContain(
      "AnonymouseType"
    );
  });

  it("should have generated correct methods in Anonymouse", () => {
    const anonClass = source.getClass("Anonymouse")!;

    const expectedFunctions = [
      "getStringField",
      "setStringField",
      "getStringFieldWithValue",
      "setStringFieldWithValue",
      "getNullableStringField",
      "setNullableStringField",
      "getNullableStringFieldWithValue",
      "setNullableStringFieldWithValue",
    ];

    expect(anonClass.getMethods().map((m) => m.getName())).toEqual(
      expect.arrayContaining(expectedFunctions)
    );
  });

  /**
   * string (no default)
   */
  describe("data type: stringField", () => {
    const fieldName = fieldsMap.stringField.name;
    it("should not be optional", () => {
      const field = source.getClass("Anonymouse")!.getProperty(fieldName);
      expect(field).toBeDefined();
      expect(field?.hasQuestionToken()).toBeFalsy();
    });

    it("should be initialized to empty string if no default", () => {
      const field = source.getClass("Anonymouse")!.getProperty(fieldName)!;
      expect(field.getInitializer()?.getText()).toBe('""');
    });

    it("should have TypeScript type string", () => {
      const field = source.getClass("Anonymouse")!.getProperty(fieldName)!;
      expect(field.getType().getText()).toBe(fieldsMap[fieldName].type);
    });

    it("setter should take a string argument", () => {
      const setter = source
        .getClass("Anonymouse")!
        .getMethod("setStringField")!;
      const paramType = setter.getParameters()[0].getType().getText();
      expect(paramType).toBe("string");
    });

    it("should have correct JSDoc", () => {
      const field = source.getClass("Anonymouse")!.getProperty(fieldName)!;
      const jsdoc = getJsDoc(field);
      expect(jsdoc.trim()).toBe(fieldsMap.stringField.description);
    });
  });

  /**
   * string with value
   */
  describe("data type: stringFieldWithValue", () => {
    const fieldName = fieldsMap.stringFieldWithValue.name;

    it("should not be optional", () => {
      const field = source.getClass("Anonymouse")!.getProperty(fieldName)!;
      expect(field.hasQuestionToken()).toBeFalsy();
    });

    it("should initialize to given default", () => {
      const field = source.getClass("Anonymouse")!.getProperty(fieldName)!;
      expect(field.getInitializer()?.getText()).toBe(
        `"${fieldsMap.stringFieldWithValue.default}"`
      );
    });

    it("should have correct JSDoc", () => {
      const field = source.getClass("Anonymouse")!.getProperty(fieldName)!;
      const jsdoc = getJsDoc(field);
      expect(jsdoc.trim()).toBe(fieldsMap.stringFieldWithValue.description);
    });
  });

  /**
   * string? (nullable, no default)
   */
  describe("data type: nullableStringField", () => {
    const fieldName = fieldsMap.nullableStringField.name;

    it("should be optional", () => {
      const field = source.getClass("Anonymouse")!.getProperty(fieldName)!;
      expect(field.hasQuestionToken()).toBeTruthy();
    });

    it("should initialize to empty string", () => {
      const field = source.getClass("Anonymouse")!.getProperty(fieldName)!;
      expect(field.getInitializer()?.getText()).toBe("undefined");
    });

    it("should have correct JSDoc", () => {
      const field = source.getClass("Anonymouse")!.getProperty(fieldName)!;
      const jsdoc = getJsDoc(field);
      expect(jsdoc.trim()).toBe(fieldsMap.nullableStringField.description);
    });
  });

  /**
   * string? with value
   */
  describe("data type: nullableStringFieldWithValue", () => {
    const fieldName = fieldsMap.nullableStringFieldWithValue.name;

    it("should be optional", () => {
      const field = source.getClass("Anonymouse")!.getProperty(fieldName)!;
      expect(field.hasQuestionToken()).toBeTruthy();
    });

    it("should initialize to given default", () => {
      const field = source.getClass("Anonymouse")!.getProperty(fieldName)!;
      expect(field.getInitializer()?.getText()).toBe(
        `"${fieldsMap[fieldName].default}"`
      );
    });

    it("should have correct JSDoc", () => {
      const field = source.getClass("Anonymouse")!.getProperty(fieldName)!;
      const jsdoc = getJsDoc(field);
      expect(jsdoc.trim()).toBe(
        fieldsMap.nullableStringFieldWithValue.description
      );
    });
  });
});
