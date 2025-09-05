import { writeFileSync } from "fs";
import yaml from "js-yaml";
import { describe, expect, it } from "vitest";
import prettier from "prettier";
import {
  createInstance,
  fileWriter,
} from "../../../../emi-npm/bin/getPublicActions";

import { execSync } from "child_process";
import { afterAll, test } from "vitest";
import { createSocketServer, randomBetween } from "../../common";
import {
  UserStreamAction,
  UserStreamActionReq,
  UserStreamActionRes,
} from "../../web-socket-test-output/UserStreamAction";

let wss;
let port = 8081;

const sample1 = {
  name: "socketTestingModule",
  actions: [
    {
      name: "userStream",
      url: "ws://localhost:8081",
      method: "reactive",
      description:
        "A socket connection which would generate random numbers, based on min, max, and count.",
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

describe("Generate the emi javascript web socket", () => {
  const content: string[] = [];
  const genOutput = "./web-socket-test-output";
  let files = [];
  afterAll(() => wss?.close());

  it("should generate and write the Emi JS module", async () => {
    files = globalThis.jsGenModule(JSON.stringify(sample1), {
      Tags: "typescript",
    });
    fileWriter(files, genOutput);
  });

  it("introduction", () => {
    content.push(`---
sidebar_position: 4
---
      
      `);
    content.push("# WebSocket generation in Emi");
    content.push(`
      Websockets are one of the most important aspect of reactive web programming now adays, which allow sending full-duplex messages between server and front-end.
      You can define an action method: 'reactive', and it would consider the endpoint is a websocket.
      Emi extends the standard WebSocket class in JavaScript, and in case of other generators (react for example), they would wrap the instance.

      In this document, we will generate, document and test different type of the websocket messages for javascript.

      **Important to note** that the javascript generated code, can work both in node.js and browser environment, but main focus of this
      generator is client side. Still, you can use the request and response classes create.

      **WebSocketX** class is the base class which all code for js/ts will be generated on top of it. It's very similar to standard WebSocket class,
      with exception that message and send are overriden, to accept types specificed as generic. Also the constructor, will get factory for message creation,
      in case of working with json messages.

      `);
  });

  it("introduction", () => {
    content.push("## Sample 1: Json requests and json responses");
    content.push(`
      Standard websocket allows for bytes, strings, arraybuffers to be sent on both ways. On top of that, Emi adds the option to send
      typesafe class structure over, in following example, we are working on a socket endpoint, which would expect a message with min,max,count
      variables, and would generate random number messages (as a json class, not primitive), and ui will print it.
      
      \`\`\`yaml
${yaml.dump(sample1)}
\`\`\`
      
      `);
  });

  it("should send a first request and get 5 results", () => {
    wss = createSocketServer(async (ws: WebSocket, data) => {
      const req = new UserStreamActionReq(data);
      for (let i = 0; i < req.count; i++) {
        // Let's send a json object as string
        ws.send(
          JSON.stringify(
            new UserStreamActionRes({}).setNumber(
              randomBetween(req.min, req.max)
            )
          )
        );
      }

      // Let's test the buffer data.
      const buf = Buffer.from([1, 2, 3, 4]);
      ws.send(buf);
    }, port);

    createInstance();
  });

  test("install dependencies", () => {
    try {
      execSync("npm install", { cwd: genOutput });
    } catch (err) {
      console.error(err);
    }
  });

  it("should run WebSocket client test", async () => {
    const ws = UserStreamAction.Create();
    ws.addEventListener("open", () => {
      ws.send(new UserStreamActionReq({}).setCount(5).setMin(60).setMax(80));
    });

    const objectsRecieved: UserStreamActionRes[] = [];
    const arrayBufferRecieved: Blob[] = [];

    const test = new Promise<void>((resolve) => {
      ws.onmessage = (msg) => {
        if (msg.data instanceof UserStreamActionRes) {
          objectsRecieved.push(msg.data);
        } else {
          arrayBufferRecieved.push(msg.data);
        }

        if (objectsRecieved.length === 5) {
          for (const item of objectsRecieved) {
            expect(item instanceof UserStreamActionRes).toBeTruthy();
            expect(!isNaN(Number(item.number))).toBe(true);
            expect(item.number).toBeGreaterThan(50);
          }
        }

        for (const item of arrayBufferRecieved) {
          if (item instanceof Blob) {
            blobToUint8Array(item).then((res) => {
              expect(Array.from(res)).toEqual([1, 2, 3, 4]);
            });
          }
        }

        if (objectsRecieved.length === 5 && arrayBufferRecieved.length === 1) {
          ws.close();
          resolve();
        }
      };
    });

    await test;
  });

  it("write the result class", async () => {
    expect(
      (files as any).find((x: any) => x.Name === "UserStreamAction")
    ).toBeTruthy();
    content.push("## Generated content");
    content.push(`
      Now the result is going to be a complete WebRequestX instance, with proper typings:
       
      `);

    const formatted = await prettier.format(
      (files as any).find((x: any) => x.Name === "UserStreamAction")
        ?.ActualScript,
      {
        parser: "typescript",
      }
    );

    content.push("```ts\r\n" + formatted + "\r\n```");
  });

  it("should write final documentation", () => {
    writeFileSync(
      "../../emi-web/docs/js/emi-javascript-web-socket.mdx",
      content.join("\r\n").trim()
    );
  });
});

function blobToUint8Array(blob: Blob): Promise<Uint8Array> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => resolve(new Uint8Array(reader.result as ArrayBuffer));
    reader.onerror = reject;
    reader.readAsArrayBuffer(blob);
  });
}
