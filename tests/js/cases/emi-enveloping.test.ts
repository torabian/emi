import http from "http";
import { afterAll, beforeAll, describe, expect, it } from "vitest";
import { fileWriter } from "../../../emi-npm/bin/getPublicActions";
import { execSync } from "child_process";
import {
  HttpActionAction,
  HttpActionActionRes,
} from "../test-artifacts/http-envelop-test-output/HttpActionAction";
import { GResponse } from "../test-artifacts/http-envelop-test-output/sdk/envelopes";

const sample1 = {
  name: "emiHttpEnvelopTest",
  actions: [
    {
      name: "httpAction",
      url: "http://localhost:8081 (for test we use override)",
      method: "post",
      description:
        "A post request which would return an array, but enveloped in google json styleguide.",
      out: {
        envelope: "GResponse",
        fields: [
          {
            name: "recordNumber",
            type: "int",
            description:
              "Fake record number to simulate a id from database table.",
          },
        ],
      },
    },
  ],
};

let server;
let port;
let sampleVersion = "1.0.0";

async function startServer() {
  server = http.createServer((req, res) => {
    const url = req.url || "";

    if (url.includes("/json")) {
      res.setHeader("Content-Type", "application/json");
      const seq = new GResponse({ apiVersion: sampleVersion });

      seq.setCreator((data) => new HttpActionActionRes(data));
      seq.inject({
        data: {
          item: {
            recordNumber: 102,
          },
        },
      });

      res.end(seq.toString());
    } else {
      res.statusCode = 404;
      res.end("not found");
    }
  });

  await new Promise((resolve) => {
    server.listen(0, () => {
      const addr = server.address();
      port = addr.port;
      console.log(`Server running on http://localhost:${port}`);
      resolve(true);
    });
  });
}

beforeAll(async () => {
  return startServer();
});

afterAll(async () => {
  await new Promise<void>((resolve) => server.close(() => resolve()));
});

describe("Generate the http sdk for it", () => {
  const content: string[] = [];
  const genOutput = "./test-artifacts/http-envelop-test-output";
  let files = [];

  it("should generate and write the Emi JS module", async () => {
    files = globalThis.jsGenModule(JSON.stringify(sample1), {
      Tags: "typescript",
    });
    await fileWriter(files, genOutput);
  });

  it("install dependencies", () => {
    try {
      execSync("npm install", { cwd: genOutput });
    } catch (err) {
      console.error(err);
    }
  });

  it("expect the version field is set correctly when read from envelope", async () => {
    const { response: res } = await HttpActionAction.Fetch(
      undefined,
      undefined,
      undefined,
      undefined,
      `http://localhost:${port}/json`
    );

    expect(res.result instanceof GResponse).toBeTruthy();

    if (res.result instanceof GResponse) {
      const envelope = res.result;
      console.log(1000, envelope.apiVersion);
      expect(envelope.apiVersion).toBe(sampleVersion);
      expect(envelope.data.item.recordNumber).toBe(102);
    }
  });
});
