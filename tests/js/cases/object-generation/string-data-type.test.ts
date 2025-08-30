/**
 * Tests the string and string? data type.
 * Emi supports both string and nullable string
 * definition, and default value forced by compiler.
 * We need to test all cases.
 */
import { describe, it, expect } from "vitest";
import { runEmiActionTs } from "../../common";

describe("Generating different data types from fields need to be working fine.", () => {
  const fieldsMap = {
    stringField: { type: "string", name: "stringField" },
    stringFieldWithValue: {
      type: "string",
      default: "testvalue",
      name: "stringFieldWithValue",
    },
    nullableStringField: { type: "string?", name: "nullableStringField" },
    nullableStringFieldWithValue: {
      type: "string?",
      default: "stringvalue",
      name: "nullableStringFieldWithValue",
    },
  };

  const fields = Object.keys(fieldsMap).map((key) => ({
    name: key,
    ...fieldsMap[key],
  }));

  const { source } = runEmiActionTs("jsGenObject", fields, {
    Flags: "Anonymouse",
    Tags: "react,typescript",
  });

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
      expect(field.getInitializer()?.getText()).toBe(undefined);
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
  });
});
