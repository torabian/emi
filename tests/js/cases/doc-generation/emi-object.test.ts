import { describe, it, expect } from "vitest";
import { createInstance } from "../../../../emi-npm/bin/getPublicActions";
import { existsSync, writeFileSync } from "fs";
import path from "path";

var schema: any = {};

const coreDefinition = "lib/core/definitions.go";

describe("Generate the emi object documentation", () => {
  const content: string[] = [];

  it("should init the wasm", () => {
    createInstance();
  });

  it("should add the introduction", () => {
    content.push(`
---
sidebar_position: 3
---

# Javascript object generation

Every request body, request response, message on web socket, entity, dto in Emi are called 'object'. The object is generated based on
fields array, which can be nested, of course. Emi uses the same function for generting different type of classes.

## Javascript classes vs objects

When working with APIs, it’s tempting to just define a TypeScript type and fetch data. For example:

\`\`\`ts
type Post = { userId: number; id: number; title: string; body: string };

async function fetchPost(id: number): Promise<Post> {
  const res = await fetch(\`https://jsonplaceholder.typicode.com/posts/\${id}\`);
  return res.json();
}

// Example usage
fetchPost(1).then(post => {
  console.log("Title:", post.title);
});

\`\`\`

This looks fine at first glance, but it has some important problems:

### Problems with plain types

1. Backend drift
The backend may evolve over time. Fields might be renamed, removed, or added. Your TypeScript types won’t catch this at runtime — they only exist at compile time.

2. No defaults
If the backend omits a field, you’ll end up with \`undefined\`. There’s no way to guarantee initialized values, which makes your code brittle.

3. Runtime unsafety
Even if TypeScript says everything is fine, the actual JSON from the server might be malformed or missing required fields. TypeScript does not validate data at runtime.


### Solution: Use Classes

By wrapping responses into **classes**, we get:

- Defaults for missing fields  
- Runtime-safe construction of objects  
- A clean place for helper methods  
- Extending them and embed the extra logic in single place rather than functions scattered across the app.

Take a look at the rewritten mechanism, this time using classes:


\`\`\`ts
class Post {
  userId: number;
  id: number;
  title: string;
  body: string;

  constructor(data: Partial<Post>) {
    this.userId = data.userId ?? 0;
    this.id = data.id ?? 0;
    this.title = data.title ?? "Untitled";
    this.body = data.body ?? "";
  }

  static async fetch(id: number): Promise<Post> {
    const res = await fetch(\`https://jsonplaceholder.typicode.com/posts/\${id}\`);
    const json = await res.json();
    return new Post(json);
  }

  shortBody(): string {
    return this.body.slice(0, 50) + "...";
  }
}

// Usage
(async () => {
  const post = await Post.fetch(1);
  console.log("Title:", post.title);
  console.log("Preview:", post.shortBody());
})();

\`\`\`

Now, this is much more reliable code instead of just typing them, you almost never get undefined. But still, writing that manually
is a hassle, and type checking and other tools might be lost. Now, let's take a look how this could be generated using Emi js compiler:

 
\`\`\`yaml
- name: userId
  type: int64
- name: id
  type: int64
- name: title
  type: string
- name: body
  type: string
\`\`\`

And Emi compiler would generate such shopisticated details:



        `);
  });

  it("compile the Emi fields array and show it as json", () => {
    const classResult = globalThis.jsGenObject(
      JSON.stringify([
        {
          name: "userId",
          type: "int64",
        },
      ]),
      {
        Flags: "Anonymouse",
      }
    );

    console.log(classResult);
    content.push("```ts\r\n" + classResult + "\r\n```");
  });

  it("should generate the schema directly from wasm", () => {
    schema = JSON.parse(globalThis.genEmiSpec("", {}));
  });

  /// Last step is to write the document down
  it("should write the final doc", () => {
    writeFileSync(
      "../../emi-web/docs/js/emi-object.mdx",
      content.join("\r\n").trim()
    );
  });
});
