import React from "react";
import http from "http";
import { afterAll, beforeAll, describe, expect, it } from "vitest";
import { fileWriter } from "../../../emi-npm/bin/getPublicActions";
import { execSync } from "child_process";
import {
  HttpActionAction,
  HttpActionActionRes,
  useHttpActionAction,
  useHttpActionActionQuery,
} from "../test-artifacts/http-emi-react-output/HttpActionAction";
import { GResponse } from "../test-artifacts/http-emi-react-output/sdk/envelopes";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { renderHook, act, waitFor } from "@testing-library/react";
import { FetchxContext } from "../test-artifacts/http-emi-react-output/sdk/common/fetchx";

function createWrapper() {
  const queryClient = new QueryClient({
    defaultOptions: { queries: { retry: false } },
  });
  return ({ children }) => (
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  );
}

const sample1 = {
  name: "emi-react-test",
  actions: [
    {
      name: "httpAction",
      url: "http://localhost:8081 (for test we use override)",
      method: "get",
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
  const genOutput = "./test-artifacts/http-emi-react-output";
  let files = [];

  it("should generate and write the Emi react module", async () => {
    files = globalThis.jsGenModule(JSON.stringify(sample1), {
      Tags: "typescript,react",
    });
    await fileWriter(files, genOutput);
  });

  // it("install dependencies", () => {
  //   try {
  //     execSync("npm install", { cwd: genOutput });
  //   } catch (err) {
  //     console.error(err);
  //   }
  // });

  it("expect the version field is set correctly when read from envelope", async () => {
    const { response: res } = await HttpActionAction.Fetch(
      undefined,
      undefined,
      undefined,
      undefined,
      undefined,
      `http://localhost:${port}/json`
    );

    expect(res.result instanceof GResponse).toBeTruthy();

    if (res.result instanceof GResponse) {
      const envelope = res.result;
      expect(envelope.apiVersion).toBe(sampleVersion);
      expect(envelope.data.item.recordNumber).toBe(102);
    }
  });
});

describe("React hooks for HttpActionAction", () => {
  it("useHttpActionActionQuery should fetch data successfully", async () => {
    const { result } = renderHook(
      () =>
        useHttpActionActionQuery({
          overrideUrl: `http://localhost:${port}/json`,
          ctx: new FetchxContext(""),
        }),
      { wrapper: createWrapper() }
    );

    // wait until query is settled
    await result.current.refetch();

    // wait until query is marked as success
    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true);
    });

    expect(result.current.isSuccess).toBe(true);
    expect(
      (result.current.data as unknown as GResponse<HttpActionActionRes>).data
        .item.recordNumber
    ).toBe(102);
  });
});
