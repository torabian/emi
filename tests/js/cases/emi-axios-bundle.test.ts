import { describe, expect, it } from "vitest";
import { createInstance } from "../../../emi-npm/bin/getPublicActions";
import prettier from "prettier";

const sample1 = {
  name: "axiosBundle",
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

describe("Generate the axios bundle, and it should be a valid file", () => {
  const content: string[] = [];
  let axiosBundleFile: string | null = null;
  let files = [];

  it("should init the wasm", () => {
    createInstance();
  });

  it("without 'axiosbundle' tag, it should not contain the axios bundle file.", async () => {
    const invalidFiles = globalThis.jsGenModule(JSON.stringify(sample1), {
      Tags: "typescript",
    }) as any;

    expect(invalidFiles.map((file) => file.Name)).not.toContain(
      "AxiosBundleAxiosClient"
    );
  });

  it("should generate and the axios bundle", async () => {
    files = globalThis.jsGenModule(JSON.stringify(sample1), {
      Tags: "typescript,axiosbundle",
    }) as any;

    expect(files.map((file: any) => file.Name)).toContain(
      "AxiosBundleAxiosClient"
    );
    axiosBundleFile =
      (files.find((item: any) => item.Name === "AxiosBundleAxiosClient") as any)
        ?.ActualScript || "";
  });

  it("should have been generated a valid typescript file", () => {
    expect(typeof axiosBundleFile).toBe("string");
  });

  it("should have been formatted with prettier", async () => {
    const formatted = await prettier.format(axiosBundleFile as string, {
      parser: "typescript",
    });

    axiosBundleFile = formatted;
  });
});
