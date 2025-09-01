import { writeFileSync } from "fs";
import prettier from "prettier";
import { describe, it } from "vitest";
import { createInstance } from "../../../../emi-npm/bin/getPublicActions";
import yaml from "js-yaml";

describe("Generate the emi object documentation", () => {
  const content: string[] = [];
  const exampleFields = [
    { name: "userId", type: "int64" },
    { name: "id", type: "int64" },
    { name: "title", type: "string" },
    { name: "body", type: "string" },
  ];

  it("should init the wasm", () => {
    createInstance();
  });

  it("should add the introduction", () => {
    content.push(`
---
sidebar_position: 3
---

import BenchmarkTest from './ObjectBenchmark';

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
${yaml.dump(exampleFields)}
\`\`\`

And Emi compiler would generate such shopisticated details (content is a bit long but worth to take a look at it.)
Emi compiler generates ts type and, class, with full getter, setters and validators for each field.

        `);
  });

  it("compile the Emi fields array and show it as json", async () => {
    const classResult = globalThis.jsGenObject(JSON.stringify(exampleFields), {
      Flags: "Anonymouse",
    });

    const formatted = await prettier.format(classResult, {
      parser: "typescript",
    });

    content.push("```ts\r\n" + formatted + "\r\n```");
  });

  it("should add the caveat section", () => {
    content.push(
      `As you see, there is a lot of details handled in such object generation, to make sure
      
      the data integrity is kept in perfect level both in compile time and run time.


## Performance and usage concerns

Generating such classes is very common on most programming languages, but in javascript it might be a bit of concern,
mainly for following reasons:

* Using getField and setField functions, is slower than directly accessing .field property of an object.
* Setting value using js path, such as set('field.user.address[0].stree', 'New street') might be compromised

## Comparing the performance

Let's make an interactive comparison, how slower it would be and you'll be suprised getters and setters are
faster in some environments.

It looks odd at first, but the reason getter/setter sometimes beats direct property access comes down to how modern 
JavaScript engines optimize. A plain object can change shape at any time, which makes the engine cautious and may 
cause de-optimizations during tight loops. By contrast, a class with fixed methods has a stable shape, so the JIT 
compiler can inline and optimize those calls very aggressively. On top of that, microbenchmarks can be misleading: 
if values aren’t used in a meaningful way, the engine may partially skip or simplify the operations. 
In practice, the difference is just a few milliseconds across millions of iterations, so for real-world
code, both approaches are effectively free.
I prefer this response


<BenchmarkTest />

As you might see, the numbers are very close, and in real life, you would not call millions of times
one after each other a javascript setter. Hence, Emi compiler extra checks make sense to bring the safety
into the code base.
      `
    );
  });

  /// Last step is to write the document down
  it("should write the final doc", () => {
    writeFileSync(
      "../../emi-web/docs/js/emi-object.mdx",
      content.join("\r\n").trim()
    );
  });
});
