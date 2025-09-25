// GetSinglePostAction.test.ts
import http from "http";
import { afterAll, beforeAll, describe, expect, it } from "vitest";
import { fileWriter } from "../../../emi-npm/bin/getPublicActions";
import {
  HttpActionAction,
  HttpActionActionRes,
} from "../test-artifacts/http-action-test-output/HttpActionAction";
import { execSync } from "child_process";

const sample1 = {
  name: "httpTestingModule",
  actions: [
    {
      name: "httpAction",
      url: "http://localhost:8081 (for test we use override)",
      method: "post",
      description:
        "A post connection which would generate random numbers, based on min, max, and count. For normal response, it would return an object, for SSE would be streaming.",
      in: {
        fields: [
          {
            name: "min",
            type: "int",
            description: "Minimum number which can be generated",
          },
          {
            name: "max",
            type: "int",
            description: "Maximum number which can be generated",
          },
          {
            name: "count",
            type: "int",
            description:
              "How many numbers you want to be generated based on maximum and minimum",
          },
        ],
      },
      out: {
        fields: [
          {
            name: "number",
            type: "int",
          },
        ],
      },
    },
  ],
};

let server;
let port;

async function startServer() {
  server = http.createServer((req, res) => {
    const url = req.url || "";

    if (url.includes("/json")) {
      res.setHeader("Content-Type", "application/json");
      res.end(JSON.stringify({ number: 42 }));
    } else if (url.includes("/text")) {
      res.setHeader("Content-Type", "text/plain");
      res.end("plain text response");
    } else if (url.includes("/attachment")) {
      res.setHeader("Content-Disposition", "attachment; filename=a.txt");
      res.end("BINARY_CONTENT");
    } else if (url.includes("/sse")) {
      res.setHeader("Content-Type", "text/event-stream");
      res.write(`data: {"number":10}\n\n`);
      res.end();
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

describe("Generates the http action codes and classes", () => {
  const content: string[] = [];
  const genOutput = "./test-artifacts/http-action-test-output";
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

  it("handles JSON response", async () => {
    const res = await HttpActionAction.Fetch(undefined, {
      overrideUrl: `http://localhost:${port}/json`,
    });
    expect(res.response.result).toBeInstanceOf(HttpActionActionRes);
    expect((res.response.result as HttpActionActionRes)?.number).toBe(42);
  });

  it("handles plain text", async () => {
    const res = await HttpActionAction.Fetch(undefined, {
      overrideUrl: `http://localhost:${port}/text`,
    });
    expect(res.response.result).toBe("plain text response");
  });

  it("handles attachment", async () => {
    const res = await HttpActionAction.Fetch(undefined, {
      overrideUrl: `http://localhost:${port}/attachment`,
    });

    const bodyText = await streamToString(res.response.result);
    expect(bodyText).toBe("BINARY_CONTENT");
  });

  it("handles SSE", async () => {
    let received = "";
    await HttpActionAction.Fetch(undefined, {
      overrideUrl: `http://localhost:${port}/sse`,
      onMessage: (ev) => {
        received = ev.data;
      },
    });
    expect(received).toBe('{"number":10}');
  });
});

async function streamToString(stream: ReadableStream | any) {
  if (!stream) return "";
  if (typeof stream.getReader === "function") {
    // Web streams
    const reader = stream.getReader();
    let result = "";
    while (true) {
      const { done, value } = await reader.read();
      if (done) break;
      result += new TextDecoder().decode(value);
    }
    return result;
  }
  // Node.js streams fallback
  if (typeof stream.arrayBuffer === "function") {
    const buf = await stream.arrayBuffer();
    return new TextDecoder().decode(buf);
  }
  return stream.toString();
}
