import { describe, it, expect } from "vitest";
import { getJsDoc, runEmiActionTs } from "../../common";
import path from "path";
import { writeFileSync } from "fs";

describe("Generating numeric data types", () => {
  const fieldsMap = {
    boolField: {
      type: "bool",
      name: "boolField",
      description: "bool field, non-nullable",
      $jstype: "boolean",
      $initializerKind: "TrueKeyword",
    },
    boolFieldWithValue: {
      type: "bool",
      name: "boolFieldWithValue",
      description: "bool field with default",
      default: true,
      $jstype: "boolean",
      $initializerKind: "TrueKeyword",
    },
    nullableboolField: {
      type: "bool?",
      name: "nullableboolField",
      description: "nullable bool",
      default: undefined,
      $jstype: "boolean",
      $initializerKind: "TrueKeyword",
    },
    nullableboolFieldWithValue: {
      type: "bool?",
      name: "nullableboolFieldWithValue",
      description: "nullable bool with default",
      default: true,
      $jstype: "boolean",
      $initializerKind: "TrueKeyword",
    },

    intField: {
      type: "int",
      name: "intField",
      description: "int field, non-nullable",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
    },
    intFieldWithValue: {
      type: "int",
      name: "intFieldWithValue",
      description: "int field with default",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: 42,
    },
    nullableIntField: {
      type: "int?",
      name: "nullableIntField",
      description: "nullable int",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: undefined,
    },
    nullableIntFieldWithValue: {
      type: "int?",
      name: "nullableIntFieldWithValue",
      description: "nullable int with default",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: 7,
    },

    int32Field: {
      type: "int32",
      name: "int32Field",
      description: "int32 field, non-nullable",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: undefined,
    },
    int32FieldWithValue: {
      type: "int32",
      name: "int32FieldWithValue",
      description: "int32 with default",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: 100,
    },
    nullableInt32Field: {
      type: "int32?",
      name: "nullableInt32Field",
      description: "nullable int32",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: undefined,
    },
    nullableInt32FieldWithValue: {
      type: "int32?",
      name: "nullableInt32FieldWithValue",
      description: "nullable int32 with default",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: 200,
    },

    int64Field: {
      type: "int64",
      name: "int64Field",
      description: "int64 field",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: undefined,
    },
    int64FieldWithValue: {
      type: "int64?",
      name: "int64FieldWithValue",
      description: "int64 with default",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: 123,
    },
    nullableInt64Field: {
      type: "int64?",
      name: "nullableInt64Field",
      description: "nullable int64",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: undefined,
    },
    nullableInt64FieldWithValue: {
      type: "int64?",
      name: "nullableInt64FieldWithValue",
      description: "nullable int64 with default",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: 456,
    },

    float32Field: {
      type: "float32",
      name: "float32Field",
      description: "float32 field",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: undefined,
    },
    float32FieldWithValue: {
      type: "float32",
      name: "float32FieldWithValue",
      description: "float32 with default",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: 1.23,
    },
    nullableFloat32Field: {
      type: "float32?",
      name: "nullableFloat32Field",
      description: "nullable float32",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: undefined,
    },
    nullableFloat32FieldWithValue: {
      type: "float32?",
      name: "nullableFloat32FieldWithValue",
      description: "nullable float32 with default",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: 4.56,
    },

    float64Field: {
      type: "float64",
      name: "float64Field",
      description: "float64 field",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: undefined,
    },
    float64FieldWithValue: {
      type: "float64",
      name: "float64FieldWithValue",
      description: "float64 with default",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: 7.89,
    },
    nullableFloat64Field: {
      type: "float64?",
      name: "nullableFloat64Field",
      description: "nullable float64",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: undefined,
    },
    nullableFloat64FieldWithValue: {
      type: "float64?",
      name: "nullableFloat64FieldWithValue",
      description: "nullable float64 with default",
      $jstype: "number",
      $initializerKind: "NumericLiteral",
      default: 0.12,
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

  writeFileSync(path.join(__filename.replace(".test.ts", ".output.ts")), resp);

  it("should generate class and type", () => {
    expect(source.getClasses().map((c) => c.getName())).toContain("Anonymouse");
    expect(source.getTypeAliases().map((t) => t.getName())).toContain(
      "AnonymouseType"
    );
  });

  Object.keys(fieldsMap).forEach((fieldKey) => {
    const f = fieldsMap[fieldKey];
    describe(`data type: ${fieldKey}`, () => {
      it("should have correct TypeScript type (everything needs to be number in this test)", () => {
        const field = source.getClass("Anonymouse")!.getProperty(f.name)!;
        expect(field.getType().getText()).toBe(f.$jstype);
      });

      if (f.default) {
        it(`should initialize to ${f.default}`, () => {
          const field = source.getClass("Anonymouse")!.getProperty(f.name)!;
          const initializer = field.getInitializer()?.getText() ?? "undefined";

          // Important test for number.
          expect(field.getInitializer()?.getKindName()).toBe(
            f.$initializerKind
          );

          if (f.$initializerKind == "NumericLiteral") {
            expect(+initializer).toBe(f.default);
          } else if (f.$initializerKind == "TrueKeyword") {
            expect(initializer === "true" ? true : false).toBe(f.default);
          }
        });
      }

      it(`${f.name}: should ${
        f.type.endsWith("?") ? "be optional" : "not be optional"
      }`, () => {
        const field = source.getClass("Anonymouse")!.getProperty(f.name)!;
        expect(field.hasQuestionToken()).toBe(f.type.endsWith("?"));
      });

      it("should have correct JSDoc", () => {
        const field = source.getClass("Anonymouse")!.getProperty(f.name)!;
        expect(getJsDoc(field).trim()).toBe(f.description);
      });
      it("setter should take single argument of correct type", () => {
        const setter = source
          .getClass("Anonymouse")!
          .getMethod(`set${f.name.charAt(0).toUpperCase()}${f.name.slice(1)}`)!;
        expect(setter).toBeDefined();
        const paramType = setter.getParameters()[0].getType().getText();
        expect(paramType).toBe(f.$jstype);
      });
    });
  });
});
