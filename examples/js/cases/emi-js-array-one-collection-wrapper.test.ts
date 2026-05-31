import { writeFileSync } from "fs";
import path from "path";
import { describe, expect, it } from "vitest";
import { runEmiActionTs, runEmiActionTsNoCheck } from "../common";

describe("Data types with wrappers, one, collection, array need to work as intended.", () => {
  const { resp: userResp } = runEmiActionTs("jsGenObject", [], {
    Flags: JSON.stringify({ name: "User" }),
    Tags: "react,typescript",
  });

  writeFileSync(path.join(__dirname + "/User.ts"), userResp);

  const fieldsMap = {
    arrayField: {
      type: "array",
      name: "arrayField",
      description: "array field, non-nullable",
      $jstype: "array",
      $initializerKind: "ArrayLiteral",
    },
    arrayNullableField: {
      type: "arrayNullable",
      name: "arrayNullableField",
      description: "arrayNullable field, non-nullable",
      $jstype: "array?",
      $initializerKind: "ArrayNullableLiteral",
    },
    collectionField: {
      type: "collection",
      name: "collectionField",
      target: "User",
      description: "collection field, non-nullable",
      $jstype: "collection",
      $initializerKind: "collectionLiteral",
    },
    collectionNullableField: {
      type: "collection?",
      name: "collectionNullableField",
      description: "collectionNullable field, non-nullable",
      target: "User",
      $jstype: "collection?",
      $initializerKind: "collectionNullableLiteral",
    },
    oneField: {
      type: "one",
      name: "oneField",
      target: "User",
      description: "one field, non-nullable",
      $jstype: "one",
      $initializerKind: "oneLiteral",
    },
    oneNullableField: {
      type: "one?",
      name: "oneNullableField",
      description: "oneNullable field, non-nullable",
      target: "User",
      $jstype: "one?",
      $initializerKind: "oneNullableLiteral",
    },
  };
  const fields = Object.keys(fieldsMap).map((key) => ({
    name: key,
    ...fieldsMap[key],
  }));

  /// Now how to import in vite test the created typescript file, I want import Anonymouse class.
  // But I do not want to do it in head of document, because file is generated after we wrtite it not before

  it("should be able to create an instance of it on typescript generation", async () => {
    const { resp } = runEmiActionTsNoCheck("jsGenObject", fields, {
      Flags: JSON.stringify({ name: "Anonymouse" }),
      Tags: "react,typescript",
    });

    const output = path.join(__filename.replace(".test.ts", ".output.ts"));
    writeFileSync(output, resp);

    const mod = await import(output);

    const { Anonymouse } = mod;

    expect(Anonymouse).toBeDefined();

    const instance = new Anonymouse({
      arrayField: {
        __operation: "append",
        items: [new Anonymouse.ArrayField(), new Anonymouse.ArrayField()],
      },
    });

    expect(instance).toBeInstanceOf(Anonymouse);
  });

  it("should be able to create an instance of it on javascript generation", async () => {
    const { resp } = runEmiActionTsNoCheck("jsGenObject", fields, {
      Flags: JSON.stringify({ name: "Anonymouse" }),
      Tags: "react",
    });

    const output = path.join(__filename.replace(".test.ts", ".output.js"));
    writeFileSync(output, resp);

    const mod = await import(output);

    const { Anonymouse } = mod;

    expect(Anonymouse).toBeDefined();

    const instance = new Anonymouse({
      arrayField: {
        __operation: "append",
        items: [new Anonymouse.ArrayField(), new Anonymouse.ArrayField()],
      },
    });

    expect(instance).toBeInstanceOf(Anonymouse);
  });

  // Generate the runtime (plain JS) module once, at collection time, so every
  // scenario below can import the very same class. We deliberately do NOT reuse
  // the `.output.js` written inside the test above — that one is produced at run
  // time, so depending on it here would couple us to test execution order.
  const scenarioOutput = path.join(__dirname, "scenarios.output.js");
  writeFileSync(
    scenarioOutput,
    runEmiActionTsNoCheck("jsGenObject", fields, {
      Flags: JSON.stringify({ name: "Anonymouse" }),
      Tags: "react",
    }).resp,
  );

  // Lazily pull in the freshly generated class together with the runtime wrappers
  // and the referenced `User` entity. Importing the wrappers/User through the same
  // specifiers the generated module uses keeps the module instances identical, so
  // `instanceof MArray` / `instanceof User` checks line up with the generated code.
  const loadScenario = async () => {
    const [{ Anonymouse }, ops, { User }] = await Promise.all([
      import(scenarioOutput),
      import("./sdk/common/operators"),
      import("./User"),
    ]);

    return { Anonymouse, User, ...ops };
  };

  describe("string <-> instance conversion and M* operator scenarios", () => {
    it("hydrates every wrapped field from a JSON string", async () => {
      const { Anonymouse, MArray, MCollection, MOne, User } =
        await loadScenario();

      const json = JSON.stringify({
        arrayField: [{}, {}],
        collectionField: [{}, {}, {}],
        oneField: {},
      });

      const instance = new Anonymouse(json);

      expect(instance).toBeInstanceOf(Anonymouse);

      // array -> MArray of inline ArrayField DTOs
      expect(instance.arrayField).toBeInstanceOf(MArray);
      expect(instance.arrayField.isReplace()).toBe(true);
      expect(instance.arrayField.len()).toBe(2);
      expect(instance.arrayField.get()[0]).toBeInstanceOf(Anonymouse.ArrayField);

      // collection -> MCollection of the target entity (User)
      expect(instance.collectionField).toBeInstanceOf(MCollection);
      expect(instance.collectionField.len()).toBe(3);
      expect(
        instance.collectionField.get().every((u: unknown) => u instanceof User),
      ).toBe(true);

      // one -> MOne wrapping a User
      expect(instance.oneField).toBeInstanceOf(MOne);
      expect(instance.oneField.get()).toBeInstanceOf(User);
    });

    it("round-trips object -> string -> object without drift", async () => {
      const { Anonymouse } = await loadScenario();

      const original = new Anonymouse({
        arrayField: [{}, {}],
        collectionField: [{}],
        oneField: {},
      });

      const serialised = original.toString();
      expect(typeof serialised).toBe("string");

      const revived = new Anonymouse(serialised);
      expect(revived).toBeInstanceOf(Anonymouse);

      // A second hop must produce byte-identical JSON — the wire form is stable.
      expect(JSON.parse(revived.toString())).toEqual(JSON.parse(serialised));
    });

    it("serialises a replace array as a bare list and an append array as a tagged object", async () => {
      const { Anonymouse, MArray } = await loadScenario();

      // replace is implicit on the wire — a bare array.
      const replace = new Anonymouse();
      replace.arrayField = [
        new Anonymouse.ArrayField(),
        new Anonymouse.ArrayField(),
      ];
      expect(replace.arrayField.isReplace()).toBe(true);

      const replaceWire = JSON.parse(replace.toString()).arrayField;
      expect(Array.isArray(replaceWire)).toBe(true);
      expect(replaceWire).toHaveLength(2);

      // append is tagged so the reader keeps the existing rows.
      const append = new Anonymouse();
      append.arrayField = MArray.append([new Anonymouse.ArrayField()]);
      expect(append.arrayField.isAppend()).toBe(true);

      const appendWire = JSON.parse(append.toString()).arrayField;
      expect(Array.isArray(appendWire)).toBe(false);
      expect(appendWire).toMatchObject({ __operation: "append" });
      expect(appendWire.items).toHaveLength(1);
    });

    it("casts plain objects in a collection into User instances and preserves a tagged append", async () => {
      const { Anonymouse, MCollection, User } = await loadScenario();

      const replace = new Anonymouse({ collectionField: [{}, {}] });
      expect(replace.collectionField.isReplace()).toBe(true);
      expect(
        replace.collectionField.get().every((u: unknown) => u instanceof User),
      ).toBe(true);

      // A tagged append object handed straight to the constructor is recognised
      // through MCollection.cast and keeps its "append" operation.
      const appended = new Anonymouse({
        collectionField: { __operation: "append", items: [{}, {}, {}] },
      });
      expect(appended.collectionField).toBeInstanceOf(MCollection);
      expect(appended.collectionField.isAppend()).toBe(true);
      expect(appended.collectionField.len()).toBe(3);
    });

    it("wraps a one field regardless of the input shape", async () => {
      const { Anonymouse, MOne, User } = await loadScenario();

      // plain object
      const fromPlain = new Anonymouse({ oneField: {} });
      expect(fromPlain.oneField).toBeInstanceOf(MOne);
      expect(fromPlain.oneField.get()).toBeInstanceOf(User);

      // already a User instance
      const fromUser = new Anonymouse();
      fromUser.oneField = new User({});
      expect(fromUser.oneField).toBeInstanceOf(MOne);
      expect(fromUser.oneField.get()).toBeInstanceOf(User);

      // already an MOne wrapper — assigned through untouched
      const wrapper = MOne.of(new User({}));
      const fromWrapper = new Anonymouse();
      fromWrapper.oneField = wrapper;
      expect(fromWrapper.oneField).toBe(wrapper);
    });

    it("emits a selector wire form for MOne.select", async () => {
      const { MOne } = await loadScenario();

      const selector = MOne.select({ id: 42 });
      expect(selector.isSelector()).toBe(true);

      // A selector serialises to the explicit replace form rather than content.
      expect(JSON.parse(JSON.stringify(selector))).toEqual({
        __operation: "replace",
        selector: { id: 42 },
      });
    });

    it("setters accept both the M* wrapper form and the pure value form", async () => {
      const { Anonymouse, MArray, MCollection, MOne, User } =
        await loadScenario();

      const instance = new Anonymouse();

      // pure value forms — setters return `this` for chaining.
      expect(instance.setArrayField([new Anonymouse.ArrayField()])).toBe(
        instance,
      );
      expect(instance.arrayField).toBeInstanceOf(MArray);
      expect(instance.arrayField.isReplace()).toBe(true);

      expect(instance.setCollectionField([{}])).toBe(instance);
      expect(instance.collectionField.get()[0]).toBeInstanceOf(User);

      expect(instance.setOneField(new User({}))).toBe(instance);
      expect(instance.oneField.get()).toBeInstanceOf(User);

      // wrapper forms — assigned through while keeping their operation.
      instance.setArrayField(MArray.append([new Anonymouse.ArrayField()]));
      expect(instance.arrayField.isAppend()).toBe(true);

      instance.setCollectionField(MCollection.append([new User({})]));
      expect(instance.collectionField.isAppend()).toBe(true);

      instance.setOneField(MOne.of(new User({})));
      expect(instance.oneField).toBeInstanceOf(MOne);
    });

    it("leaves nullable fields undefined when never supplied", async () => {
      const { Anonymouse } = await loadScenario();

      const instance = new Anonymouse({});
      expect(instance.collectionNullableField).toBeUndefined();
      expect(instance.oneNullableField).toBeUndefined();

      // ...and they are simply omitted from the serialised wire form.
      const wire = JSON.parse(instance.toString());
      expect("collectionNullableField" in wire).toBe(false);
      expect("oneNullableField" in wire).toBe(false);
    });

    it("accepts an explicit null for a nullable collection", async () => {
      const { Anonymouse } = await loadScenario();

      const instance = new Anonymouse();
      instance.collectionNullableField = null;
      expect(instance.collectionNullableField).toBeNull();
    });
  });
});
